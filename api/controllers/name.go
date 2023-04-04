package controllers

import (
	"api/libs"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type NameController struct {
	beego.Controller
}

func (c *NameController) Get() {
	firstname := c.GetString(":firstname")
	sex := c.GetString(":sex")
	data, err := libs.GetChatGpt(map[string]string{"firstname": firstname, "sex": sex})
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}
