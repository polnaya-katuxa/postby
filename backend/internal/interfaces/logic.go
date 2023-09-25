package interfaces

import (
	"context"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type IBalanceTransactionLogic interface {
	Increase(context.Context, int, string) error
	Decrease(context.Context, int, string) error
	IncreaseSub(ctx context.Context, value int, userID uuid.UUID) error
}

type ISubscriptionLogic interface {
	Subscribe(context.Context, uuid.UUID) error
	GetSubscriptions(context.Context) ([]*models.User, error)
	IsSubscribed(context.Context, uuid.UUID, uuid.UUID) (bool, error)
}

type IPostLogic interface {
	Publish(context.Context, *models.Post) (*models.Post, error)
	GetAll(context.Context, uuid.UUID) ([]*models.Post, error)
	Get(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	Delete(context.Context, uuid.UUID) error
	View(context.Context, *models.Post, uuid.UUID) (bool, error)
	React(context.Context, uuid.UUID, uuid.UUID) (bool, error)
	Comment(ctx context.Context, comm *models.Comment) (*models.Comment, error)
	Uncomment(ctx context.Context, id uuid.UUID) error
	GetAllComments(ctx context.Context, postID uuid.UUID) ([]*models.Comment, error)
	ChangePerms(context.Context, uuid.UUID) error
	PopularityCheck(context.Context, uuid.UUID) error
	GetReactionTypes(context.Context) ([]*models.ReactionType, error)
}

type IFeedLogic interface {
	View(context.Context) ([]*models.Post, []*models.Post, error)
}

type IProfileLogic interface {
	AuthByToken(context.Context, string) (*models.User, error)
	Register(context.Context, *models.User, string) (string, error)
	Login(context.Context, string, string) (string, error)
	Get(context.Context, string) (*models.User, error)
	Delete(context.Context, string) error
	GetAll(context.Context) ([]*models.User, error)
}
