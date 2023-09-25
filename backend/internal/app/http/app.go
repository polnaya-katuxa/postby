package http

import (
	"context"
	"errors"
	"fmt"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"path/filepath"

	openapi "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/server"
	"github.com/rs/cors"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/controllers"
	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	mg "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo"
	pg "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

const (
	PostgresDB = "postgres"
	MongoDB    = "mongo"
)

type App struct {
	postgresDB *gorm.DB
	mongoDB    *mongo.Database
	cfg        Config
	handler    http.Handler
	logger     *zap.SugaredLogger
}

func New() *App {
	return &App{}
}

func (a *App) readConfig(cfgFile string) error {
	viper.SetConfigName(filepath.Base(cfgFile))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(cfgFile))

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read in config: %w", err)
	}

	err = viper.Unmarshal(&a.cfg)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (a *App) initAdmin(u *controllers.Profile) error {
	_, err := u.Get(context.Background(), "admin")
	if !errors.Is(err, errors2.ErrNotFound) {
		if err != nil {
			return fmt.Errorf("get: %w", err)
		}
		return nil
	}

	admin := &models.User{
		Login:   a.cfg.AdminLogin,
		Mail:    a.cfg.AdminEmail,
		IsAdmin: true,
	}

	_, err = u.Register(context.Background(), admin, a.cfg.AdminPassword)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	return nil
}

func (a *App) initLogger() error {
	lvl, err := zap.ParseAtomicLevel(a.cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("parse level: %w", err)
	}

	logConfig := zap.Config{
		Level:    lvl,
		Encoding: a.cfg.Logger.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeTime:   zapcore.RFC3339TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{a.cfg.Logger.File},
		ErrorOutputPaths: []string{a.cfg.Logger.File},
	}

	logger, err := logConfig.Build()
	if err != nil {
		return fmt.Errorf("build logger: %w", err)
	}

	a.logger = logger.Sugar()

	return nil
}

func (a *App) Init(cfg string) error {
	err := a.readConfig(cfg)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	err = a.initLogger()
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	var postRepo interfaces.IPostRepository
	var subRepo interfaces.ISubscriptionRepository
	var reactRepo interfaces.IReactionRepository
	var userRepo interfaces.IUserRepository
	var btRepo interfaces.IBalanceTransactionRepository
	var commRepo interfaces.ICommentRepository

	if a.cfg.DB == PostgresDB {
		a.postgresDB, err = gorm.Open(postgres.Open(a.cfg.PG.toDSN()), &gorm.Config{
			Logger: zapgorm2.New(a.logger.Desugar()),
		})
		if err != nil {
			a.logger.Fatalw("cannot open gorm connection", "error", err)
			return fmt.Errorf("gorm open: %w", err)
		}

		postRepo = pg.NewPR(a.postgresDB)
		subRepo = pg.NewSR(a.postgresDB)
		reactRepo = pg.NewRR(a.postgresDB)
		userRepo = pg.NewUR(a.postgresDB)
		btRepo = pg.NewBTR(a.postgresDB)
		commRepo = pg.NewCR(a.postgresDB)
	} else if a.cfg.DB == MongoDB {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(a.cfg.Mongo.toDSN()))
		if err != nil {
			a.logger.Fatalw("cannot open mongo connection", "error", err)
			return fmt.Errorf("mongo connect: %w", err)
		}
		a.mongoDB = client.Database("postby")

		postRepo = mg.NewPR(a.mongoDB)
		subRepo = mg.NewSR(a.mongoDB)
		reactRepo = mg.NewRR(a.mongoDB)
		userRepo = mg.NewUR(a.mongoDB)
		btRepo = mg.NewBTR(a.mongoDB)
		commRepo = mg.NewCR(a.mongoDB)
	}

	btl := controllers.NewBTL(userRepo, btRepo)
	fl := controllers.NewFL(subRepo, postRepo, a.cfg.Span)
	sl := controllers.NewSL(userRepo, btl, subRepo, a.cfg.Cost)
	pl := controllers.NewPL(reactRepo, btl, postRepo, sl, commRepo)
	prl := controllers.NewPRL(userRepo, btl, subRepo, pl, a.cfg.DailyBonus, a.cfg.TokenExp, a.cfg.SecretKey)

	err = a.initAdmin(prl)
	if err != nil {
		a.logger.Fatalw("cannot init admin", "error", err)
		return fmt.Errorf("init admin: %w", err)
	}

	service := NewServer(prl, fl, pl, sl)

	a.handler = middleware(prl, a.logger, openapi.NewRouter(openapi.NewDefaultApiController(service)))

	return nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Infow("started running application")
	ctx = mycontext.LoggerToContext(ctx, a.logger)

	port := fmt.Sprintf(":%d", a.cfg.HTTPPort)
	err := http.ListenAndServe(port, cors.AllowAll().Handler(a.handler))
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
