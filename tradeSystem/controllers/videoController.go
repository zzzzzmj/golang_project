package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"tradeSystem/models"
)

type ToVideoIndex struct {
	beego.Controller
}

type SubmitMyVideo struct {
	beego.Controller
}

type WatchSubmitedVideo struct {
	beego.Controller
}

//跳转到video模块首页
func (vp *ToVideoIndex) Get() {
	vp.TplName = "video/videoIndex.html"
}

//跳转到上传视频页面
func (sv *SubmitMyVideo) Get() {
	sv.TplName = "video/submitPage.html"
}

//处理上传视频表单的数据
func (sv *SubmitMyVideo) Post() {
	o := orm.NewOrm()

	ua := sv.GetSession("userAccount")
	title := sv.GetString("title")
	f, h, err := sv.GetFile("video")
	createTime := time.Now().Format("2006-01-02-15-04-05")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		filepath := "./static/video/" + createTime + h.Filename
		videoUrl := "/static/video/" + createTime + h.Filename
		sv.SaveToFile("video", filepath)
		video := models.Video{SubmitUser: ua.(string), SubmitTime: createTime, VideoTitle: title, VideoUrl: videoUrl}
		o.Insert(&video)
		sv.Ctx.WriteString("视频上传成功！")
	}
}

//跳转到已经发布的视频页面
func (wsv *WatchSubmitedVideo) Get() {
	o := orm.NewOrm()
	var videos []orm.Params
	o.QueryTable("video").Values(&videos)
	wsv.Data["videos"] = videos
	wsv.TplName = "video/watchSubmitedVideo.html"
}
