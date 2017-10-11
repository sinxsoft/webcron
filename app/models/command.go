package models

import (
	"github.com/astaxie/beego/orm"
)

type Command struct{
	// <!--`id`, `command_id`, `command_name`, `command_type`, `description`, 
	// `command_text`, `status`, `timeout`, `create_time` -->

	Id int
	CommandId int
	CommandName string
	CommandType int
	Description string
	CommandText string
	Status int
	Timeout int
	CreateTime int64
}


func (c *Command) TableName() string{
	return TableName("command")
}

func (c *Command) Update(fields ...string) error{
	if _,err := orm.NewOrm().Update(c,fields...); err != nil{
		return err
	}
	return nil
}

func CommandAdd(command *Command) (int64,error){
	return orm.NewOrm().Insert(command)
}

func CommandGetAllList() ([]*Command,error){

	list :=make( []*Command,0)

	_,err:=orm.NewOrm().QueryTable(TableName("command")).All(&list)

	if err != nil{
		return nil,err
	}
	return list,nil

}

func CommandGetById(id int) (*Command ,error){
	c := new(Command)
	err:=orm.NewOrm().QueryTable(TableName("command")).Filter("id",id).One(c)
	if err != nil{
		return nil,err
	}
	return c,nil
}

func CommandGetByName(commandName string)(*Command,error){
	c := new (Command)

	err := orm.NewOrm().QueryTable(TableName("command")).Filter("command_name",commandName).One(c)

	if err != nil{
		return nil ,err
	}
	return c,nil
}

func CommandUpdate(command *Command,fields ...string) error{

	_,err := orm.NewOrm().Update(command,fields...)

	return err
}

