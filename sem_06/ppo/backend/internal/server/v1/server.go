package v1

import (
	"context"
	"net/http"
	"sort"
	"time"

	"github.com/google/uuid"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	openapi "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/openapi/v1"
)

type Server struct {
	openapi.DefaultApiService
	profileLogic interfaces.IProfileLogic
	feedLogic    interfaces.IFeedLogic
	postLogic    interfaces.IPostLogic
	subLogic     interfaces.ISubscriptionLogic
}

func NewServer(prl interfaces.IProfileLogic, fl interfaces.IFeedLogic, pl interfaces.IPostLogic, sl interfaces.ISubscriptionLogic) *Server {
	return &Server{
		profileLogic: prl,
		feedLogic:    fl,
		postLogic:    pl,
		subLogic:     sl,
	}
}

func (s *Server) Login(ctx context.Context, req openapi.LoginRequest) (openapi.ImplResponse, error) {
	token, err := s.profileLogic.Login(ctx, req.Login, req.Password)
	if err != nil {
		return toErrorResponse(err, "Cannot login.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AuthResponse{Token: token},
	}, nil
}

func (s *Server) Register(ctx context.Context, req openapi.RegisterRequest) (openapi.ImplResponse, error) {
	token, err := s.profileLogic.Register(ctx, &models.User{
		Login:       req.Login,
		Mail:        req.Mail,
		Picture:     req.Picture,
		Description: req.Description,
	}, req.Password)
	if err != nil {
		return toErrorResponse(err, "Cannot register.")
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.AuthResponse{Token: token},
	}, nil
}

func makePost(p *models.Post, types []*models.ReactionType, u *models.User) openapi.Post {
	r := make(map[string]openapi.Reaction, len(types))
	for _, t := range types {
		r[t.UUID.String()] = openapi.Reaction{
			Icon:   t.Icon,
			TypeID: t.UUID.String(),
		}
	}

	for _, re := range p.Reactions {
		r[re.TypeID.String()] = openapi.Reaction{
			Icon:   re.Icon,
			Num:    r[re.TypeID.String()].Num + 1,
			TypeID: re.TypeID.String(),
			Yours:  r[re.TypeID.String()].Yours || (re.ReactorID == u.UUID),
		}
	}

	reacts := make([]openapi.Reaction, 0, len(r))
	for _, v := range r {
		reacts = append(reacts, v)
	}

	sort.Slice(reacts, func(i, j int) bool {
		return reacts[i].Num > reacts[j].Num
	})

	return openapi.Post{
		Id:      p.UUID.String(),
		Content: p.Content,
		PubTime: p.PublicTime.Format(time.RFC3339),
		Author: openapi.User{
			Id:          p.Author.UUID.String(),
			Login:       p.Author.Login,
			Picture:     p.Author.Picture,
			Description: p.Author.Description,
			Balance:     int32(p.Author.Balance),
			Mail:        p.Author.Mail,
			IsAdmin:     p.Author.IsAdmin,
		},
		CommentsNum: int32(p.CommentNum),
		Reactions:   reacts,
		Perms:       p.Perms == models.Paid,
	}
}

func (s *Server) ViewFeed(ctx context.Context) (openapi.ImplResponse, error) {
	posts, err := s.feedLogic.View(ctx, nil)
	if err != nil {
		return toErrorResponse(err, "Cannot view feed.")
	}

	sub := make([]openapi.Post, len(posts))
	noSub := make([]openapi.Post, 0)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	types, err := s.postLogic.GetReactionTypes(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot react to posts.")
	}

	for i, p := range posts {
		sub[i] = makePost(p, types, user)
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ViewFeedResponse{
			SubPosts:   sub,
			NoSubPosts: noSub,
		},
	}, nil
}

func (s *Server) ViewUsers(ctx context.Context) (openapi.ImplResponse, error) {
	users, err := s.profileLogic.GetAll(ctx, nil)
	if err != nil {
		return toErrorResponse(err, "Cannot view users.")
	}

	usersOA := make([]openapi.User, len(users))

	for i, u := range users {
		usersOA[i] = openapi.User{
			Id:          u.UUID.String(),
			Login:       u.Login,
			Picture:     u.Picture,
			Description: u.Description,
			Balance:     int32(u.Balance),
			Mail:        u.Mail,
			IsAdmin:     u.IsAdmin,
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ViewUsersResponse{
			Users: usersOA,
		},
	}, nil
}

func (s *Server) GetPost(ctx context.Context, id string) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	post, err := s.postLogic.Get(ctx, postID)
	if err != nil {
		return toErrorResponse(err, "Cannot get post.")
	}

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	types, err := s.postLogic.GetReactionTypes(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot react to posts.")
	}

	postOA := makePost(post, types, user)

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetPostResponse{
			Post: postOA,
		},
	}, nil
}

func (s *Server) ChangePostPerms(ctx context.Context, id string) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.postLogic.ChangePerms(ctx, postID)
	if err != nil {
		return toErrorResponse(err, "Cannot change perms.")
	}

	post, err := s.postLogic.Get(ctx, postID)
	if err != nil {
		return toErrorResponse(err, "Cannot get post.")
	}

	p := ""
	if post.Perms == models.Free {
		p = "free"
	} else {
		p = "paid"
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ChangePermsResponse{
			Changed: p,
		},
	}, nil
}

func (s *Server) DeletePost(ctx context.Context, id string) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.postLogic.Delete(ctx, postID)
	if err != nil {
		return toErrorResponse(err, "Cannot delete post.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeletePostResponse{
			Deleted: true,
		},
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, login string) (openapi.ImplResponse, error) {
	err := s.profileLogic.Delete(ctx, login)
	if err != nil {
		return toErrorResponse(err, "Cannot delete user.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.DeleteUserResponse{
			Deleted: true,
		},
	}, nil
}

func (s *Server) ViewProfileUser(ctx context.Context, login string) (openapi.ImplResponse, error) {
	user, err := s.profileLogic.Get(ctx, login)
	if err != nil {
		return toErrorResponse(err, "Cannot view profile.")
	}

	userReader, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	self, subscribed := false, false
	if userReader.UUID == user.UUID {
		self = true
	} else {
		subscribed, err = s.subLogic.IsSubscribed(ctx, userReader.UUID, user.UUID)
		if err != nil {
			return toErrorResponse(err, "Cannot check subscription.")
		}
	}

	u := openapi.User{
		Id:          user.UUID.String(),
		Login:       user.Login,
		Picture:     user.Picture,
		Description: user.Description,
		Balance:     int32(user.Balance),
		Mail:        user.Mail,
		IsAdmin:     user.IsAdmin,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ViewProfileUserResponse{
			User:       u,
			Subscribed: subscribed,
			Self:       self,
		},
	}, nil
}

func (s *Server) ViewProfilePosts(ctx context.Context, login string) (openapi.ImplResponse, error) {
	user, err := s.profileLogic.Get(ctx, login)
	if err != nil {
		return toErrorResponse(err, "Cannot view profile.")
	}

	userReader, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	posts, err := s.postLogic.GetAll(ctx, user.UUID, nil)
	if err != nil {
		return toErrorResponse(err, "Cannot view profile posts.")
	}

	types, err := s.postLogic.GetReactionTypes(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot react to posts.")
	}

	p := make([]openapi.Post, len(posts))
	for i, post := range posts {
		p[i] = makePost(post, types, userReader)
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ViewProfilePostsResponse{
			Posts: p,
		},
	}, nil
}

func (s *Server) Subscribe(ctx context.Context, id string) (openapi.ImplResponse, error) {
	writerID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.subLogic.Subscribe(ctx, writerID)
	if err != nil {
		return toErrorResponse(err, "Cannot subscribe.")
	}

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	subscribed, err := s.subLogic.IsSubscribed(ctx, user.UUID, writerID)
	if err != nil {
		return toErrorResponse(err, "Cannot check subscription.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.SubscribeResponse{Subscribed: subscribed},
	}, nil
}

func (s *Server) Publish(ctx context.Context, req openapi.PublishRequest) (openapi.ImplResponse, error) {
	post := &models.Post{
		Content: req.Content,
	}

	if req.Perms {
		post.Perms = models.Paid
	} else {
		post.Perms = models.Free
	}

	p, err := s.postLogic.Publish(ctx, post)
	if err != nil {
		return toErrorResponse(err, "Cannot publish post.")
	}

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	types, err := s.postLogic.GetReactionTypes(ctx)
	if err != nil {
		return toErrorResponse(err, "Cannot react to posts.")
	}

	pOA := makePost(p, types, user)

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.PublishResponse{Post: pOA, Published: true},
	}, nil
}

func (s *Server) React(ctx context.Context, id string, req openapi.ReactRequest) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	reacted, err := s.postLogic.React(ctx, postID, typeID)
	if err != nil {
		return toErrorResponse(err, "Cannot react to post.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ReactResponse{Reacted: reacted},
	}, nil
}

func makeComm(comm *models.Comment) openapi.Comment {
	commOA := openapi.Comment{
		Id:      comm.UUID.String(),
		Content: comm.Content,
		PubTime: comm.PublicTime.Format(time.RFC3339),
		Commentator: openapi.User{
			Id:          comm.Commentator.UUID.String(),
			Login:       comm.Commentator.Login,
			Picture:     comm.Commentator.Picture,
			Description: comm.Commentator.Description,
			Balance:     int32(comm.Commentator.Balance),
			Mail:        comm.Commentator.Mail,
			IsAdmin:     comm.Commentator.IsAdmin,
		},
		PostID: comm.PostID.String(),
	}

	return commOA
}

func (s *Server) Comment(ctx context.Context, id string, req openapi.CommentRequest) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	comm := &models.Comment{
		Content: req.Content,
		PostID:  postID,
	}

	comm, err = s.postLogic.Comment(ctx, comm)
	if err != nil {
		return toErrorResponse(err, "Cannot comment.")
	}

	commOA := makeComm(comm)

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.CommentResponse{Comment: commOA},
	}, nil
}

func (s *Server) Uncomment(ctx context.Context, _ string, req openapi.UncommentRequest) (openapi.ImplResponse, error) {
	commID, err := uuid.Parse(req.CommID)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	err = s.postLogic.Uncomment(ctx, commID)
	if err != nil {
		return toErrorResponse(err, "Cannot uncomment.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UncommentResponse{Uncommented: true},
	}, nil
}

func (s *Server) ViewComments(ctx context.Context, id string) (openapi.ImplResponse, error) {
	postID, err := uuid.Parse(id)
	if err != nil {
		return toErrorResponse(err, "Invalid data.")
	}

	comms, err := s.postLogic.GetAllComments(ctx, postID, nil)
	if err != nil {
		return toErrorResponse(err, "Cannot view post comments.")
	}

	commsOA := make([]openapi.Comment, 0, len(comms))
	for _, v := range comms {
		commsOA = append(commsOA, makeComm(v))
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.ViewCommentsResponse{Comments: commsOA},
	}, nil
}

func (s *Server) UserInfo(ctx context.Context) (openapi.ImplResponse, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		return toErrorResponse(err, "User is not authorized.")
	}

	u := openapi.User{
		Id:          user.UUID.String(),
		Login:       user.Login,
		Picture:     user.Picture,
		Description: user.Description,
		Balance:     int32(user.Balance),
		Mail:        user.Mail,
		IsAdmin:     user.IsAdmin,
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.UserInfoResponse{User: u},
	}, nil
}
