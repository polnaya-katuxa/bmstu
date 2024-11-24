package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type Post struct {
	postRepo     interfaces.IPostRepository
	reactionRepo interfaces.IReactionRepository
	commentRepo  interfaces.ICommentRepository

	balanceTransactionLogic interfaces.IBalanceTransactionLogic
	subscriptionLogic       interfaces.ISubscriptionLogic
}

func NewPL(r interfaces.IReactionRepository, bt interfaces.IBalanceTransactionLogic, p interfaces.IPostRepository,
	s interfaces.ISubscriptionLogic, c interfaces.ICommentRepository,
) *Post {
	return &Post{
		reactionRepo:            r,
		balanceTransactionLogic: bt,
		postRepo:                p,
		subscriptionLogic:       s,
		commentRepo:             c,
	}
}

func (p *Post) Publish(ctx context.Context, post *models.Post) (*models.Post, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to publish post", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started post publication", "user", user.Login)

	limits, err := p.postRepo.GetSortedLimits(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get sorted limits to publish post", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get limits: %w", err)
	}

	post.PublicTime = time.Now().UTC()
	post.NextLimit = *limits[0]
	post.Author = user
	post.CommentNum = 0

	postBD, err := p.postRepo.Create(ctx, post)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot create published post", "user", user.Login, "post", post.PublicTime, "error", err)
		return nil, fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully published new post", "user", user.Login, "post", post.PublicTime)

	return postBD, nil
}

func (p *Post) getReactions(ctx context.Context, post *models.Post) ([]*models.Reaction, error) {
	reactions, err := p.reactionRepo.GetAll(ctx, post.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get reactions for a post", "post", post.UUID, "error", err)
		return nil, fmt.Errorf("get reactions: %w", err)
	}

	return reactions, nil
}

func (p *Post) GetAll(ctx context.Context, userID uuid.UUID, pg *models.Paginator) ([]*models.Post, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all posts", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting all user's posts", "user-reader", user.Login, "user-writer", userID)

	posts, err := p.postRepo.GetAll(ctx, userID, pg)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all user's posts", "user-reader", user.Login, "user-writer", userID, "error", err)
		return nil, fmt.Errorf("get all: %w", err)
	}

	if user.UUID == userID {
		mycontext.LoggerFromContext(ctx).Infow("viewed own profile", "user", user.Login)
		return posts, nil
	}

	accessed := make([]*models.Post, 0)
	for _, post := range posts {
		ok, err := p.View(ctx, post, user.UUID)
		if err != nil && !errors.Is(err, errors2.ErrPaid) {
			mycontext.LoggerFromContext(ctx).Warnw("cannot view accessed post", "user-reader", user.Login, "user-writer", userID, "post", post.UUID, "error", err)
			return nil, fmt.Errorf("view: %w", err)
		}

		if ok {
			accessed = append(accessed, post)
		}
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully gor all user's posts", "user-reader", user.Login, "user-writer", userID)

	return accessed, nil
}

func (p *Post) Get(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get post", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting post", "user-reader", user.Login, "post", postID)

	post, err := p.postRepo.Get(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post", "user-reader", user.Login, "post", postID, "error", err)
		return nil, fmt.Errorf("get all: %w", err)
	}

	if user.UUID == post.Author.UUID {
		mycontext.LoggerFromContext(ctx).Infow("viewed own post", "user", user.Login)
		return post, nil
	}

	ok, err := p.View(ctx, post, user.UUID)
	if err != nil && !errors.Is(err, errors2.ErrPaid) {
		mycontext.LoggerFromContext(ctx).Warnw("cannot view post", "user-reader", user.Login, "user-writer", post.Author.UUID, "post", post.UUID, "error", err)
		return nil, fmt.Errorf("view: %w", err)
	}

	if !ok {
		mycontext.LoggerFromContext(ctx).Warnw("cannot view post: paid", "user-reader", user.Login, "user-writer", post.Author.UUID, "post", post.UUID, "error", err)
		return nil, fmt.Errorf("view paid post: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got post", "user-reader", user.Login, "post", postID)

	return post, nil
}

func (p *Post) Delete(ctx context.Context, postID uuid.UUID) error {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to delete post", "post", postID, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started deleting post", "user", user.Login, "post", postID)

	post, err := p.postRepo.Get(ctx, postID)
	if err != nil {
		return fmt.Errorf("get post: %w", err)
	}

	if user.UUID != post.Author.UUID {
		mycontext.LoggerFromContext(ctx).Errorw("deleting not authored post", "user", user.Login, "post author", post.Author.Login, "post", postID, "error", err)
		return fmt.Errorf("delete: %w", errors2.ErrAuthor)
	}

	err = p.postRepo.Delete(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete post", "user", user.Login, "post", postID, "error", err)
		return fmt.Errorf("delete: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully deleted post", "user", user.Login, "post", postID)

	return nil
}

func (p *Post) View(ctx context.Context, post *models.Post, userID uuid.UUID) (bool, error) {
	mycontext.LoggerFromContext(ctx).Infow("started viewing post", "user", userID, "post", post.UUID)

	if post.Author.UUID == userID {
		mycontext.LoggerFromContext(ctx).Infow("viewing own post", "user", userID, "post", post.UUID)
		return false, nil
	}

	if post.Perms == models.Free {
		mycontext.LoggerFromContext(ctx).Infow("viewing free post", "user", userID, "post", post.UUID)
		return true, nil
	}

	ok, err := p.subscriptionLogic.IsSubscribed(ctx, userID, post.Author.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot check if subscribed for the paid post", "user", userID, "post", post.UUID, "error", err)
		return false, fmt.Errorf("is subscribed: %w", err)
	}

	if !ok {
		mycontext.LoggerFromContext(ctx).Infow("cannot access post", "user", userID, "post", post.UUID)
		return false, fmt.Errorf("post access: %w", errors2.ErrPaid)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully viewed post", "user", userID, "post", post.UUID)

	return true, nil
}

func (p *Post) React(ctx context.Context, postID uuid.UUID, typeID uuid.UUID) (bool, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to react", "post", postID, "error", err)
		return false, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started reacting", "user", user.Login, "post", postID)

	post, err := p.postRepo.Get(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post to react", "user", user.Login, "post", postID, "error", err)
		return false, fmt.Errorf("get post: %w", err)
	}

	ok, err := p.View(ctx, post, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot view post to react", "user", user.Login, "post", postID, "error", err)
		return false, fmt.Errorf("view: %w", err)
	}
	if !ok {
		mycontext.LoggerFromContext(ctx).Infow("cannot react to own post", "user", user.Login, "post", postID)
		return false, fmt.Errorf("own post: %w", errors2.ErrReact)
	}

	ok, err = p.reactionRepo.Reacted(ctx, user.UUID, post.UUID, typeID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot check if reacted", "user", user.Login, "post", postID, "error", err)
		return false, fmt.Errorf("reacted: %w", err)
	}

	if ok {
		err = p.reactionRepo.Delete(ctx, user.UUID, post.UUID, typeID)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot unreact", "user", user.Login, "post", postID, "error", err)
			return false, fmt.Errorf("delete: %w", err)
		}
	} else {
		r := &models.Reaction{
			TypeID:    typeID,
			ReactorID: user.UUID,
		}

		err = p.reactionRepo.Create(ctx, r, post.UUID)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot react", "user", user.Login, "post", postID, "error", err)
			return false, fmt.Errorf("create: %w", err)
		}
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully reacted to post", "user", user.Login, "post", postID)

	return !ok, nil
}

func (p *Post) Comment(ctx context.Context, comm *models.Comment) (*models.Comment, error) {
	mycontext.LoggerFromContext(ctx).Infow("started commenting", "post", comm.PostID)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to comment", "post", comm.PostID, "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	post, err := p.postRepo.Get(ctx, comm.PostID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post to comment", "user", user.Login, "post", comm.PostID, "error", err)
		return nil, fmt.Errorf("get post: %w", err)
	}

	if post.Author.UUID != user.UUID {
		_, err := p.View(ctx, post, user.UUID)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot view post to comment", "user", user.Login, "post", comm.PostID, "error", err)
			return nil, fmt.Errorf("view: %w", err)
		}
	}

	c := &models.Comment{
		Content:     comm.Content,
		PublicTime:  time.Now().UTC(),
		Commentator: user,
		PostID:      comm.PostID,
	}

	c, err = p.commentRepo.Create(ctx, c)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot comment", "user", user.Login, "post", comm.PostID, "error", err)
		return nil, fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully commented post", "user", user.Login, "post", comm.PostID)

	return c, nil
}

func (p *Post) Uncomment(ctx context.Context, commID uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Infow("started uncommenting", "comment", commID)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to uncomment", "comment", commID, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	comm, err := p.commentRepo.GetByID(ctx, commID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get by commID", "user", user.Login, "comment", commID, "error", err)
		return fmt.Errorf("get comment by commID: %w", err)
	}

	post, err := p.postRepo.Get(ctx, comm.PostID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post by commID", "user", user.Login, "comment", commID, "error", err)
		return fmt.Errorf("get post by commID: %w", err)
	}

	if user.UUID != post.Author.UUID && user.UUID != comm.Commentator.UUID {
		mycontext.LoggerFromContext(ctx).Warnw("cannot uncomment: rules", "user", user.Login, "comment", commID, "error", err)
		return fmt.Errorf("cannot uncomment: %w", errors2.ErrUncomment)
	}

	err = p.commentRepo.Delete(ctx, commID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot uncomment", "user", user.Login, "comment", commID, "error", err)
		return fmt.Errorf("get post: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully uncommented post", "user", user.Login, "comment", commID)

	return nil
}

func (p *Post) GetAllComments(ctx context.Context, postID uuid.UUID, pg *models.Paginator) ([]*models.Comment, error) {
	mycontext.LoggerFromContext(ctx).Infow("started getting all comments", "post", postID)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all comments", "post", postID, "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	post, err := p.postRepo.Get(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post to get all comments", "user", user.Login, "post", postID, "error", err)
		return nil, fmt.Errorf("get post: %w", err)
	}

	_, err = p.View(ctx, post, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot view post to get all comments", "user", user.Login, "post", postID, "error", err)
		return nil, fmt.Errorf("view: %w", err)
	}

	c, err := p.commentRepo.GetAll(ctx, postID, pg)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all comments", "user", user.Login, "post", postID, "error", err)
		return nil, fmt.Errorf("view: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got all comments", "user", user.Login, "post", postID)

	return c, nil
}

func (p *Post) GetReactionTypes(ctx context.Context) ([]*models.ReactionType, error) {
	mycontext.LoggerFromContext(ctx).Infow("started getting reaction types")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get reaction types", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	reactTypes, err := p.postRepo.GetReactionTypes(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get reaction types", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get reaction types: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got reaction types")

	return reactTypes, nil
}

func (p *Post) ChangePerms(ctx context.Context, postID uuid.UUID) error {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to change post perms", "post", postID, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started changing post perms", "user", user.Login, "post", postID)

	post, err := p.postRepo.Get(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get post to change post perms", "user", user.Login, "post", postID, "error", err)
		return fmt.Errorf("get post: %w", err)
	}

	if user.UUID != post.Author.UUID {
		mycontext.LoggerFromContext(ctx).Errorw("cannot change perms of non authored post", "user", user.Login, "post", postID, "error", err)
		return fmt.Errorf("perms: %w", errors2.ErrAuthor)
	}

	if post.Perms == models.Free {
		post.Perms = models.Paid
	} else {
		post.Perms = models.Free
	}

	err = p.postRepo.Update(ctx, post)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot update post perms", "user", user.Login, "post", postID, "new perm", post.Perms, "error", err)
		return fmt.Errorf("update: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully changed post perms", "user", user.Login, "post", postID, "new perm", post.Perms)

	return nil
}

func (p *Post) nextLimit(old models.Limit, limits []*models.Limit) models.Limit {
	nextIndex := -1

	for i, l := range limits {
		if l.Value > old.Value {
			nextIndex = i
		}
	}

	return *limits[nextIndex]
}

func (p *Post) GetTotalCommentsByPostID(ctx context.Context, postID uuid.UUID) (int, error) {
	_, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get total number of comments by post id", "post", postID, "error", err)
		return 0, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting total number of comments by post id", "post", postID)

	num, err := p.commentRepo.GetTotal(ctx, postID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get number comments by post id", "post", postID, "error", err)
		return 0, fmt.Errorf("get total by user id: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got total number of comments by post id", "post", postID)

	return num, nil
}

func (p *Post) GetTotalByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	_, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get total number of posts by id", "error", err)
		return 0, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting total number of posts by id", "user", userID)

	num, err := p.postRepo.GetTotalByUserID(ctx, userID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user's posts number by id", "user", userID, "error", err)
		return 0, fmt.Errorf("get total by user id: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got total number of posts by id", "user", userID)

	return num, nil
}

func (p *Post) PopularityCheck(ctx context.Context, userID uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Infow("started popularity check for posts", "user", userID)

	posts, err := p.postRepo.GetAll(ctx, userID, nil)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user's posts", "user", userID, "error", err)
		return fmt.Errorf("get posts: %w", err)
	}

	for _, post := range posts {
		if len(post.Reactions) > post.NextLimit.Value {
			s := fmt.Sprintf("post limit %d bonus", post.NextLimit.Value)

			err := p.balanceTransactionLogic.Increase(ctx, post.NextLimit.Bonus, s)
			if err != nil {
				mycontext.LoggerFromContext(ctx).Warnw("cannot increase balance for passing post limit", "user", post.Author.Login, "post", post.UUID, "passed limit", post.NextLimit.Value, "error", err)
				return fmt.Errorf("increase: %w", err)
			}

			limits, err := p.postRepo.GetSortedLimits(ctx)
			if err != nil {
				mycontext.LoggerFromContext(ctx).Warnw("cannot get sorted limits", "user", post.Author.Login, "post", post.UUID, "limit", post.NextLimit.Value, "error", err)
				return fmt.Errorf("get limits: %w", err)
			}

			post.NextLimit = p.nextLimit(post.NextLimit, limits)

			err = p.postRepo.Update(ctx, post)
			if err != nil {
				mycontext.LoggerFromContext(ctx).Warnw("cannot update post limit", "user", post.Author.Login, "post", post.UUID, "limit", post.NextLimit.Value, "error", err)
				return fmt.Errorf("update: %w", err)
			}

			mycontext.LoggerFromContext(ctx).Infow("passed post limit", "user", post.Author.Login, "post", post.UUID, "new limit", post.NextLimit.Value)
		}
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully finished popularity check for posts", "user", userID)

	return nil
}
