package controllers

import (
	"context"
	"fmt"
	"time"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type Feed struct {
	subscriptionRepo interfaces.ISubscriptionRepository
	postRepo         interfaces.IPostRepository

	span time.Duration
}

func NewFL(s interfaces.ISubscriptionRepository, p interfaces.IPostRepository, span time.Duration) *Feed {
	return &Feed{
		subscriptionRepo: s,
		postRepo:         p,
		span:             span,
	}
}

func (f *Feed) View(ctx context.Context) ([]*models.Post, []*models.Post, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to view feed", "error", err)
		return nil, nil, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started view feed", "user", user.Login)

	subs, err := f.subscriptionRepo.GetAll(ctx, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subs for feed view", "user", user.Login, "error", err)
		return nil, nil, fmt.Errorf("user from context: %w", err)
	}

	subsID := make([]uuid.UUID, 0, len(subs))
	for _, s := range subs {
		subsID = append(subsID, s.WriterID)
	}

	subsPosts, err := f.postRepo.GetByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), true)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subs posts for feed view", "user", user.Login, "error", err)
		return nil, nil, fmt.Errorf("get subs posts by span: %w", err)
	}

	notSubsPosts, err := f.postRepo.GetByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), false)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get not subs free posts for feed view", "user", user.Login, "error", err)
		return nil, nil, fmt.Errorf("get subs posts by span: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully view feed", "user", user.Login)

	return subsPosts, notSubsPosts, nil
}
