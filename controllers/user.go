package controllers

import (
	"strings"

	"myproject/models"

	"github.com/astaxie/beego/validation"

	"github.com/lisijie/goblog/util"
)

type UserController struct {
	baseController
}

func (c *UserController) Register() {

	c.TplName = "register.html"
}

//添加用户
func (this *UserController) Save() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		confirmpass := strings.TrimSpace(this.GetString("confirmpass"))
		email := strings.TrimSpace(this.GetString("email"))

		input["username"] = username
		input["password"] = password
		input["confirmpass"] = confirmpass
		input["email"] = email

		valid := validation.Validation{}

		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名"
		} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能大于15个字符"
		}

		if v := valid.Required(password, "password"); !v.Ok {
			errmsg["password"] = "请输入密码"
		}

		if v := valid.Required(confirmpass, "confirmpass"); !v.Ok {
			errmsg["confirmpass"] = "请再次输入密码"
		} else if password != confirmpass {
			errmsg["confirmpass"] = "两次输入的密码不一致"
		}

		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		}

		if len(errmsg) == 0 {
			var user models.User
			user.UserName = username
			user.Password = util.Md5([]byte(password))
			user.Email = email
			if err := user.Insert(); err != nil {
				//				this.showmsg(err.Error())
			}
			//			this.Redirect("/admin/user/list", 302)
		}

	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.TplName = "dosave.html"
}
