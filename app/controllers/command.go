package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/sinxsoft/webcron/app/libs"
	"github.com/sinxsoft/webcron/app/models"
	// "strconv"
	"strings"
)

type CommandController struct{
	BaseController
}


func (this *CommandController) List(){

	list,err := models.CommandGetAllList()
	if err != nil{
		return
	}

	this.Data["pageTitle"] = "命令列表"
	this.Data["list"] = list
	//this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("TaskController.List", "groupid", groupId), true).ToString()
	this.display()

}

func (this *CommandController) Edit(){

	id,_ := this.GetInt("id")

	command, err:= models.CommandGetById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		//command.Id = id
		command.CommandId,_ = this.GetInt("commandId")
		command.CommandName = strings.TrimSpace(this.GetString("commandName"))
		command.CommandText = strings.TrimSpace(this.GetString("commandName"))
		command.CommandType,_ = this.GetInt("commandType")
		command.Description = strings.TrimSpace(this.GetString("description"))
		command.Status,_ = this.GetInt("status")
		err := command.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}else{
		this.Data["command"] = command
		this.Data["pageTitle"] = "添加指令"
		this.display()
	}
}

func (this *CommandController) Add(){
	
	if this.isPost() {
		command := new(models.Command)
		command.CommandId,_ = this.GetInt("commandId")
		command.CommandName = this.GetString("commandName")
		command.CommandText = this.GetString("commandName")
		command.CommandType = 1
		command.Description = this.GetString("description")

		if command.CommandName == "" || command.CommandText == "" || command.Description== "" {
			this.ajaxMsg("请填写完整信息", MSG_ERR)
		}
	
		_, err := models.CommandAdd(command)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}else{
			this.ajaxMsg("", MSG_OK)
		}
	}

	// 分组列表
	commands, _ := models.CommandGetAllList()
	this.Data["commands"] = commands
	this.Data["pageTitle"] = "添加指令"
	this.display()
	
}




