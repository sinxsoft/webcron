package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/sinxsoft/webcron/app/controllers"
	"github.com/sinxsoft/webcron/app/jobs"
	_ "github.com/sinxsoft/webcron/app/mail"
	"github.com/sinxsoft/webcron/app/models"
)

const VERSION = "1.0.0"

func main() {
	models.Init()
	jobs.InitJobs()

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 生产环境不输出debug日志
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})

	//add a test page by henrik
	beego.AutoRouter(&controllers.CommandController{})

	beego.BConfig.WebConfig.Session.SessionOn = false
	beego.Run()
}

//---------------------------------以下是用reflect实现一些类型无关的泛型编程示例
//new object same the type as sample
func New(sample interface{}) interface{} {
	t := reflect.ValueOf(sample).Type()
	v := reflect.New(t).Interface()
	return v
}

//---------------------------------check type of aninterface
func CheckType(val interface{}, kind reflect.Kind) bool {
	v := reflect.ValueOf(val)
	return kind == v.Kind()
}

//---------------------------------if _func is not a functionor para num and type not match,it will cause panic
func Call(_func interface{}, params ...interface{}) (result []interface{}, err error) {
	f := reflect.ValueOf(_func)
	if len(params) != f.Type().NumIn() {
		ss := fmt.Sprintf("The number of params is not adapted.%s", f.String())
		panic(ss)
		return
	}
	var in []reflect.Value
	if len(params) > 0 { //prepare in paras
		in = make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
	}
	out := f.Call(in)
	if len(out) > 0 { //prepare out paras
		result = make([]interface{}, len(out), len(out))
		for i, v := range out {
			result[i] = v.Interface()
		}
	}
	return
}
