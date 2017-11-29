package main

import (
	"fmt"
	// "net/http"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"

	"tencentyun_im/timRestAPI"
)

type User struct {
	RegisterID  string
	UserID      string
	JoinGroupID string
	QuitGroupID string
}

type GroupIdList []struct {
	GroupId string
	Name    string
}

func main() {
	app := iris.New()
	app.StaticWeb("/static", "./static")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// set the view html template engine
	// app.RegisterView(iris.HTML("./templates", ".html").Reload(true))
	app.RegisterView(iris.HTML("./static", ".html").Reload(true))

	sess := sessions.New(sessions.Config{
		Cookie:  "mysessionid",
		Expires: time.Hour * 1,
		// if you want to invalid cookies on different subdomains
		// of the same host, then enable it
		DisableSubdomainPersistence: false,
	})

	app.Handle("GET", "/", func(ctx iris.Context) {
		sess.Start(ctx).Set("UserID", "")
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})

	app.Handle("POST", "/", func(ctx iris.Context) {
		if err := ctx.View("form.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
		user := User{}
		err := ctx.ReadForm(&user)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		} else {
			timRestAPI := timRestAPI.CreateRestAPI()

			if user.RegisterID != "" {
				timRestAPI.AccountImport(user.RegisterID, "test_user", "www.baidu.com")
				ctx.HTML("<h3>Register Complete. Please Log In.</h3>")
				return
			}

			// When a user log in firstly
			s := sess.Start(ctx)
			if user.UserID != "" {
				s.Set("UserID", user.UserID)
			}

			// Join a group
			if user.JoinGroupID != "" {
				timRestAPI.GroupAddGroupMember(user.JoinGroupID, s.GetString("UserID"), 1)
			}

			// Quit a group
			if user.QuitGroupID != "" {
				timRestAPI.GroupDeleteGroupMember(user.QuitGroupID, s.GetString("UserID"), 1)
			}

			// Show user's groups
			//if existGroups := timRestAPI.GroupGetJoinedGroupList(s.GetString("UserID")); existGroups != nil {
			//if existGroups, ok := timRestAPI.GroupGetJoinedGroupList(s.GetString("UserID")).(GroupIdList); ok && existGroups != nil {
			//if existGroups, ok := timRestAPI.GroupGetJoinedGroupList(s.GetString("UserID")).(GroupIdList); ok {
			//if existGroups, ok := timRestAPI.GroupGetJoinedGroupList("sukepp").(GroupIdList); ok {
			//if existGroups, ok := timRestAPI.GroupGetJoinedGroupList(s.GetString("UserID")).(GroupIdList); ok {
			if existGroups := timRestAPI.GroupGetJoinedGroupList(s.GetString("UserID")); existGroups != nil {
				ctx.HTML(fmt.Sprintf("<h3>Hi, %s, your groups:</h3>", s.GetString("UserID")))
				ctx.HTML(formatGroups(existGroups))
				// Show user's recommanded groups
				recommandedGroups := []string{}
				allGroups := timRestAPI.GroupGetAppidGroupList(50)
				existGroupsMap := make(map[string]bool)
				for _, group := range existGroups {
					existGroupsMap[group.GroupId] = true
				}
				for _, group := range allGroups {
					if _, found := existGroupsMap[group]; !found {
						recommandedGroups = append(recommandedGroups, group)
					}
				}
				ctx.HTML("<h3>The groups you may like:</h3>")
				ctx.HTML(formatGroupsOnlyID(recommandedGroups))
			} else {
				ctx.HTML("<h3>Get Groups Error</h3>")
			}
		}
	})
	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}

// Question: Why groups is a timRestAPI.GroupIdList?
func formatGroups(groups timRestAPI.GroupIdList) (str string) {
	if len(groups) == 0 {
		str = "&nbsp;&nbsp;&nbsp;&nbsp;Nothing"
	} else {
		str = "<ol>"
		for _, group := range groups {
			str += fmt.Sprintf("<li>%s: %s</li>", group.GroupId, group.Name)
		}
		str += "</ol>"
	}
	return str
}

func formatGroupsOnlyID(groups []string) (str string) {
	if len(groups) == 0 {
		str = "&nbsp;&nbsp;&nbsp;&nbsp;Nothing"
	} else {
		str = "<ol>"
		for _, group := range groups {
			str += fmt.Sprintf("<li>%s</li>", group)
		}
		str += "</ol>"
	}
	return str
}
