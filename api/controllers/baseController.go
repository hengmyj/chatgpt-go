package controllers

import (
	"api/libs"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	MSG_OK  = 1
	MSG_ERR = 0
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userId         string
	UserName       string
	loginName      string
	Language       string
	UserAgent      string
	pageSize       int
	allowUrl       string
}

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}

func (c *BaseController) Prepare() {
	if c.Ctx.Request.Method == http.MethodOptions {
		c.ajaxMsg("", MSG_OK, nil)
	}
	token := c.Ctx.Input.Header("Authorization")
	c.Language = c.Ctx.Input.Header("Accept-Language")
	c.UserAgent = c.Ctx.Input.Header("User-Agent")

	tmpURL, err := url.Parse(c.Ctx.Request.RequestURI)
	if err != nil {
		return
	}
	if "/v1/user/login" == tmpURL.Path {
		return
	}
	libs.Logger.Info("base-body CheckToken", token)
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := self.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	panic(nil)
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	if data != nil {
		out["data"] = data
	} else {
		out["data"] = make([]string, 0)
	}
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	self.Data["json"] = out
	self.ServeJSON()
	panic(nil)
}

//异常
func (self *BaseController) Catch() {
	if r := recover(); r != nil {
		message := fmt.Sprintf("%s", r)
		libs.Logger.Info("login-body", string(debug.Stack()), message)
		out := make(map[string]interface{})
		out["code"] = MSG_ERR
		out["msg"] = message
		out["data"] = ""
		self.Data["json"] = out
		self.ServeJSON()
		self.StopRun()
	}
}

func (self *BaseController) Options() {
	out := make(map[string]interface{})
	out["code"] = MSG_ERR
	out["msg"] = "ok"
	out["data"] = ""
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

func (self *BaseController) Valid(key []string, msg []string) map[string]interface{} {
	data := make(map[string]interface{})
	if err := json.Unmarshal(self.Ctx.Input.RequestBody, &data); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR, nil)
	}
	libs.Logger.Info("base Valid", data)
	valid := validation.Validation{}
	for i, v := range key {
		valid.Required(data[v], v).Message(msg[i] + " is empty")
	}
	// valid.Required(data["t_key"], "t_key").Message("t_key empty")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			self.ajaxMsg(err.Message, MSG_ERR, nil)
		}
	}
	return data
}

func (self *BaseController) ValidGet(key, msg string) string {
	valid := validation.Validation{}
	valid.Required(self.GetString(key), key).Message(msg + " is empty")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			self.ajaxMsg(err.Message, MSG_ERR, nil)
		}
	}
	return strings.TrimSpace(self.GetString(key))
}
