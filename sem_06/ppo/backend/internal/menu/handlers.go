package menu

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/dixonwille/wmenu/v5"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

var ErrLeft = errors.New("not authorized")

func (m *Menu) registerHandler(opt wmenu.Opt) error {
	user := models.User{}

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&user.Login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&user.Password); err != nil {
		return fmt.Errorf("input password: %w", err)
	}

	fmt.Print("Enter email: ")
	if _, err := fmt.Scan(&user.Mail); err != nil {
		return fmt.Errorf("input email: %w", err)
	}

	fmt.Print("Enter picture link: ")
	if _, err := fmt.Scan(&user.Picture); err != nil {
		return fmt.Errorf("input picture link: %w", err)
	}

	var err error
	fmt.Print("Enter profile description: ")
	user.Description, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return fmt.Errorf("input profile description: %w", err)
	}
	user.Description = strings.TrimSpace(user.Description)

	token, err := m.profileLogic.Register(opt.Value.(context.Context), &user, user.Password)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	m.token = token

	c := color.New(color.FgGreen)
	c.Printf("Successful registration.\nToken: %s\n", m.token)

	return nil
}

func (m *Menu) loginHandler(opt wmenu.Opt) error {
	var login, password string

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Enter password: ")
	if _, err := fmt.Scan(&password); err != nil {
		return fmt.Errorf("input password: %w", err)
	}

	token, err := m.profileLogic.Login(opt.Value.(context.Context), login, password)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}

	m.token = token

	c := color.New(color.FgGreen)
	c.Printf("Successful login.\nToken: %s\n", m.token)

	return nil
}

func (m *Menu) logoutHandler(opt wmenu.Opt) error {
	if m.token == "" {
		return fmt.Errorf("already left: %w", ErrLeft)
	}

	m.token = ""

	c := color.New(color.FgGreen)
	c.Printf("Successfully left.\n")

	return nil
}

func (m *Menu) auth(opt wmenu.Opt) (context.Context, error) {
	if m.token == "" {
		return nil, fmt.Errorf("already left: %w", ErrLeft)
	}

	user, err := m.profileLogic.AuthByToken(opt.Value.(context.Context), m.token)
	if err != nil {
		return nil, fmt.Errorf("auth by token: %w", err)
	}

	ctx := userContext.UserToContext(opt.Value.(context.Context), user)

	return ctx, nil
}

func (m *Menu) publicPostHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	post := models.Post{}

	fmt.Print("Enter perms (0 for 'free', or anything else for 'paid'): ")
	if _, err := fmt.Scan(&post.Perms); err != nil {
		return fmt.Errorf("input perms: %w", err)
	}

	fmt.Print("Enter post content in .md format: ")
	post.Content, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return fmt.Errorf("input content: %w", err)
	}
	post.Content = strings.TrimSpace(post.Content)

	_, err = m.postLogic.Publish(ctx, &post)
	if err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully published a new post.\n")

	return nil
}

func (m *Menu) deletePostHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var postID string
	fmt.Print("Enter post uuid: ")
	if _, err := fmt.Scan(&postID); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	id, err := uuid.Parse(postID)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.postLogic.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully deleted post.\n")

	return nil
}

func (m *Menu) changePostPermsHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var postID string
	fmt.Print("Enter post uuid: ")
	if _, err := fmt.Scan(&postID); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	id, err := uuid.Parse(postID)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.postLogic.ChangePerms(ctx, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully changed post perms.\n")

	return nil
}

func printUser(user *models.User, n int) {
	c := color.New(color.FgCyan)
	c.Printf("USER №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", user.UUID.String())
	fmt.Printf("\tlogin: %s\n", user.Login)
	fmt.Printf("\temail: %s\n", user.Mail)
	fmt.Printf("\tdescription: %s\n", user.Description)
	fmt.Printf("\tpicture: %s\n", user.Picture)
	fmt.Printf("\tbalance: %d\n", user.Balance)
	fmt.Printf("\tenter time: %s\n", user.EnterTime.Format(time.RFC822))
	fmt.Printf("\tis admin: %v\n", user.IsAdmin)
}

func (m *Menu) viewProfileHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var login string

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	user, err := m.profileLogic.Get(ctx, login)
	if err != nil {
		return fmt.Errorf("get: %w", err)
	}

	printUser(user, 0)

	posts, err := m.postLogic.GetAll(ctx, user.UUID, nil)
	if err != nil {
		return fmt.Errorf("get all: %w", err)
	}

	printPosts(posts, "AUTHORED")

	return nil
}

func printReaction(reaction *models.Reaction, n int) {
	c := color.New(color.FgCyan)
	c.Printf("REACTION №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", reaction.UUID.String())
	fmt.Printf("\ticon: %s\n", reaction.Icon)
	fmt.Printf("\treactor: %s\n", reaction.ReactorID.String())
	fmt.Printf("\treact type: %s\n", reaction.TypeID.String())
}

func printReactions(reactions []*models.Reaction) {
	for i, r := range reactions {
		printReaction(r, i)
	}
}

func printLimit(limit models.Limit, n int) {
	c := color.New(color.FgCyan)
	c.Printf("LIMIT №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", limit.UUID.String())
	fmt.Printf("\tvalue: %d\n", limit.Value)
	fmt.Printf("\tbonus: %d\n", limit.Bonus)
}

func printPost(post *models.Post, n int) {
	c := color.New(color.FgCyan)
	c.Printf("POST №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", post.UUID.String())
	fmt.Printf("\tcontent: %s\n", post.Content)
	fmt.Printf("\tpublished: %s\n", post.PublicTime.Format(time.RFC822))
	fmt.Printf("\tperms: %s\n", post.Perms.String())
	printUser(post.Author, 0)
	printReactions(post.Reactions)
	printLimit(post.NextLimit, 0)
	fmt.Println()
}

func printPosts(posts []*models.Post, header string) {
	c := color.New(color.FgBlue)
	c.Printf("%s:\n", header)
	for i, post := range posts {
		printPost(post, i)
	}
}

func (m *Menu) viewFeedHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	posts, err := m.feedLogic.View(ctx, nil)
	if err != nil {
		return fmt.Errorf("view: %w", err)
	}

	printPosts(posts, "FEED")

	return nil
}

func (m *Menu) reactHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var postIDStr, typeIDStr string

	fmt.Print("Enter post uuid: ")
	if _, err := fmt.Scan(&postIDStr); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	fmt.Print("Enter reaction type uuid: ")
	if _, err := fmt.Scan(&typeIDStr); err != nil {
		return fmt.Errorf("input reaction type uuid: %w", err)
	}

	typeID, err := uuid.Parse(typeIDStr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	_, err = m.postLogic.React(ctx, postID, typeID)
	if err != nil {
		return fmt.Errorf("react: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully changed reaction.\n")

	return nil
}

func (m *Menu) commentHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var postIDStr, content string

	fmt.Print("Enter post uuid: ")
	if _, err := fmt.Scan(&postIDStr); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	fmt.Print("Enter comment content: ")
	content, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return fmt.Errorf("input content: %w", err)
	}
	content = strings.TrimSpace(content)

	user, err := userContext.UserFromContext(ctx)
	if err != nil {
		return fmt.Errorf("user from context: %w", err)
	}

	comm := &models.Comment{
		Content:     content,
		Commentator: user,
		PostID:      postID,
	}

	_, err = m.postLogic.Comment(ctx, comm)
	if err != nil {
		return fmt.Errorf("comment: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully commented.\n")

	return nil
}

func printComm(comm *models.Comment, n int) {
	c := color.New(color.FgCyan)
	c.Printf("COMMENT №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", comm.UUID.String())
	fmt.Printf("\tcontent: %s\n", comm.Content)
	fmt.Printf("\tpublished: %s\n", comm.PublicTime.Format(time.RFC822))
	fmt.Printf("\tpost id: %s\n", comm.PostID.String())
	printUser(comm.Commentator, 0)
	fmt.Println()
}

func printComments(comms []*models.Comment, header string) {
	c := color.New(color.FgBlue)
	c.Printf("%s:\n", header)
	for i, comm := range comms {
		printComm(comm, i)
	}
}

func (m *Menu) viewCommentsHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var postIDStr string

	fmt.Print("Enter post uuid: ")
	if _, err := fmt.Scan(&postIDStr); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	comms, err := m.postLogic.GetAllComments(ctx, postID, nil)
	if err != nil {
		return fmt.Errorf("get all comments: %w", err)
	}

	printComments(comms, "COMMENTS")

	c := color.New(color.FgGreen)
	c.Printf("Successfully got all comments.\n")

	return nil
}

func (m *Menu) uncommentHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var commIDStr string

	fmt.Print("Enter comment uuid: ")
	if _, err := fmt.Scan(&commIDStr); err != nil {
		return fmt.Errorf("input comment uuid: %w", err)
	}

	commID, err := uuid.Parse(commIDStr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	err = m.postLogic.Uncomment(ctx, commID)
	if err != nil {
		return fmt.Errorf("uncomment: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully uncommented.\n")

	return nil
}

func (m *Menu) subscriptionHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var userIDStr string

	fmt.Print("Enter user-writer uuid: ")
	if _, err := fmt.Scan(&userIDStr); err != nil {
		return fmt.Errorf("input post uuid: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("parse uuid: %w", err)
	}

	err = m.subscriptionLogic.Subscribe(ctx, userID)
	if err != nil {
		return fmt.Errorf("react: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully changed subscription.\n")

	return nil
}

func printUsers(users []*models.User) {
	for i, user := range users {
		printUser(user, i)
	}
}

func (m *Menu) viewUsersHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	users, err := m.profileLogic.GetAll(ctx, nil)
	if err != nil {
		return fmt.Errorf("get all: %w", err)
	}

	printUsers(users)

	return nil
}

func (m *Menu) deleteUserHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	var login string

	fmt.Print("Enter login: ")
	if _, err := fmt.Scan(&login); err != nil {
		return fmt.Errorf("input login: %w", err)
	}

	err = m.profileLogic.Delete(ctx, login)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	c := color.New(color.FgGreen)
	c.Printf("Successfully deleted user.\n")

	return nil
}

func printRT(rt *models.ReactionType, n int) {
	c := color.New(color.FgCyan)
	c.Printf("REACTION TYPE №%d:\n", n)
	fmt.Printf("\tuuid: %s\n", rt.UUID.String())
	fmt.Printf("\ticon: %s\n", rt.Icon)
}

func printRTs(rt []*models.ReactionType) {
	for i, r := range rt {
		printRT(r, i)
	}
}

func (m *Menu) getReactionTypesHandler(opt wmenu.Opt) error {
	ctx, err := m.auth(opt)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	rt, err := m.postLogic.GetReactionTypes(ctx)
	if err != nil {
		return fmt.Errorf("get reaction types: %w", err)
	}

	printRTs(rt)

	return nil
}
