package menu

import (
	"context"
	"errors"
	"fmt"

	my_context "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"github.com/dixonwille/wmenu/v5"
	"github.com/fatih/color"
)

var errExit = errors.New("exit")

type Menu struct {
	mainMenu          *wmenu.Menu
	token             string
	profileLogic      interfaces.IProfileLogic
	postLogic         interfaces.IPostLogic
	feedLogic         interfaces.IFeedLogic
	subscriptionLogic interfaces.ISubscriptionLogic
}

func (m *Menu) AddOptions(ctx context.Context) {
	m.mainMenu.Option("Register", ctx, false, m.registerHandler)
	m.mainMenu.Option("Login", ctx, false, m.loginHandler)
	m.mainMenu.Option("Logout", ctx, false, m.logoutHandler)
	m.mainMenu.Option("Publish post", ctx, false, m.publicPostHandler)
	m.mainMenu.Option("Delete post", ctx, false, m.deletePostHandler)
	m.mainMenu.Option("Change post's perms", ctx, false, m.changePostPermsHandler)
	m.mainMenu.Option("View profile", ctx, false, m.viewProfileHandler)
	m.mainMenu.Option("View feed", ctx, false, m.viewFeedHandler)
	m.mainMenu.Option("React/Unreact", ctx, false, m.reactHandler)
	m.mainMenu.Option("Comment", ctx, false, m.commentHandler)
	m.mainMenu.Option("Unomment", ctx, false, m.uncommentHandler)
	m.mainMenu.Option("View post comments", ctx, false, m.viewCommentsHandler)
	m.mainMenu.Option("Subscribe/unsubscribe", ctx, false, m.subscriptionHandler)
	m.mainMenu.Option("Get reaction types", ctx, false, m.getReactionTypesHandler)
	m.mainMenu.Option("View users", ctx, false, m.viewUsersHandler)
	m.mainMenu.Option("Delete user", ctx, false, m.deleteUserHandler)
	m.mainMenu.Option("Exit", ctx, false, func(_ wmenu.Opt) error {
		return errExit
	})
}

func NewMenu(prl interfaces.IProfileLogic, pl interfaces.IPostLogic, fl interfaces.IFeedLogic,
	sl interfaces.ISubscriptionLogic,
) *Menu {
	return &Menu{
		profileLogic:      prl,
		postLogic:         pl,
		feedLogic:         fl,
		subscriptionLogic: sl,
	}
}

func (m *Menu) RunMenu(ctx context.Context) {
	m.mainMenu = wmenu.NewMenu("Choose an option.")
	m.AddOptions(ctx)

	for {
		err := m.mainMenu.Run()
		fmt.Println()
		if err != nil {
			if errors.Is(err, errExit) {
				break
			}

			c := color.New(color.FgRed)
			c.Printf("ERROR: %s\n\n", err)
		}
	}

	c := color.New(color.FgMagenta)
	c.Printf("Exited menu.\n")
	my_context.LoggerFromContext(ctx).Infow("exited menu")
}
