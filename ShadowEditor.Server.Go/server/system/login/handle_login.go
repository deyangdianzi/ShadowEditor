package login

import (
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/server"
)

func init() {
	login := Login{}
	server.Mux.UsingContext().Handle(http.MethodPost, "/api/Login/Login", login.Login)
	server.Mux.UsingContext().Handle(http.MethodPost, "/api/Login/Logout", login.Logout)
}

// Login 登录控制器
type Login struct {
}

// Login 登录
func (Login) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := strings.TrimSpace(r.FormValue("Username"))
	password := strings.TrimSpace(r.FormValue("Password"))

	if username == "" {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "Username is not allowed to be empty.",
		})
		return
	}

	if password == "" {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "Password is not allowed to be empty.",
		})
		return
	}

	// 获取Salt
	db, err := server.Mongo()
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	filter := bson.M{
		"Username": username,
	}

	user := bson.M{}
	find, _ := db.FindOne(server.UserCollectionName, filter, &user)
	if !find {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "The username or password is wrong.",
		})
		return
	}

	salt := user["Salt"].(string)

	// 验证账号密码
	filter1 := bson.M{
		"Password": helper.MD5(password + salt),
	}
	filter = bson.M{
		"$and": bson.A{filter, filter1},
	}

	find, _ = db.FindOne(server.UserCollectionName, filter, &user)
	if !find {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  "The username or password is wrong.",
		})
		return
	}

	id := user["ID"].(primitive.ObjectID).Hex()
	_ = id

	// // 票据数据
	// var ticketData = new LoginTicketDataModel
	// {
	// 	UserID = id,
	// };

	// // 将用户信息写入cookie
	// var cookie = FormsAuthentication.GetAuthCookie(model.Username, false);
	// var ticket = FormsAuthentication.Decrypt(cookie.Value);

	// var newTicket = new FormsAuthenticationTicket(ticket.Version, ticket.Name, ticket.IssueDate, ticket.Expiration, ticket.IsPersistent, JsonConvert.SerializeObject(ticketData)); // 将用户ID写入ticket
	// cookie.Value = FormsAuthentication.Encrypt(newTicket);
	// cookie.Expires = DateTime.Now.AddMinutes(ConfigHelper.Expires);
	// HttpContext.Current.Response.Cookies.Add(cookie);

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Login successfully!",
		Data: map[string]string{
			"Username": user["Username"].(string),
			"Name":     user["Name"].(string),
		},
	})
}

// Logout 注销
func (Login) Logout(w http.ResponseWriter, r *http.Request) {
	// var cookie = HttpContext.Current.Request.Cookies.Get(FormsAuthentication.FormsCookieName);

	// if (cookie != null)
	// {
	// 	cookie.SameSite = SameSiteMode.Lax;
	// 	cookie.Expires = DateTime.Now.AddDays(-1);
	// 	HttpContext.Current.Response.Cookies.Add(cookie);
	// }

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Logout Successfully!",
	})
}
