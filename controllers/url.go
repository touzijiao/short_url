package controllers

import (
	"short_url/models"
	"short_url/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/url"
)

type UrlController struct {
	beego.Controller
}

//Get生成一个短链接
func (this *UrlController) Get() {
	//u.Ctx.WriteString("this is urlController:Get(获取一个链接)") //写内容到到页面
	//sourceUrl := this.Ctx.Input.Query("source") //获取输入的url（长连接网址）
	sourceUrl := this.GetString("source")
	//校验
	logs.Debug("源网址为：", sourceUrl)

	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		fmt.Errorf("Url is not a valid url:" + err.Error())
		return
	}

	//正常的url保存
	u := &models.Url{
		SourceUrl: sourceUrl,
	}
	//获取或者分配此url的id
	u.GenId()

	//转换成短链接
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()

	this.Data["json"] = u
	//this.ServeJSON() //以json格式输出（一般api应用用此形式返回数据）

	this.TplName = "index.html" //在index.html文件内获取json形式的u的值显示出来
}
