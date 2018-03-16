package controllers

import (
	"short_url/models"
	"short_url/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type RedirectController struct {
	beego.Controller
}

//把任意请求跳转到"api/url"路由（从一个网址跳到另一个网址）也就是短链接还原成原来的链接
func (this *RedirectController) Get() {
	//根路径时的效果
	if this.Ctx.Request.URL.Path == "/" {
		//beego加载模板文件
		this.Data["title_name"] = "短连接服务_beego"
		this.TplName = "index.html"
		return
	}
	//不是根路径就跳转
	urlPath := this.Ctx.Request.URL.Path //获取源路径
	urlPath = strings.Trim(urlPath, "/") //处理掉"/"

	//通过我们的算法转换成id
	id := utils.StringToId(urlPath)
	u := &models.Url{
		Id: id,
	}

	//转换成的id在到MongoDB内查找(找到会植入SourceUrl)
	err := u.FindById()
	if err != nil {
		fmt.Errorf("没有找到相应的短链接:", err.Error())
		return
	}
	this.Redirect(u.SourceUrl, 302) //beego的url重定向功能（把路由到此功能方法的都跳转到/api/url路由上去）
}
