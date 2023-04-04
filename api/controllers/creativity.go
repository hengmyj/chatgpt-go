package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type CreativityController struct {
	beego.Controller
}

func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}
