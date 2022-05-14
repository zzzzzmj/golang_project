package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"tradeSystem/models"
)

type ImageUploadController struct {
	beego.Controller
}

type AddImageController struct {
	beego.Controller
}

func (ai *AddImageController) Get() {
	goodsId := ai.GetString("goodsId")
	ai.Data["goodsId"] = goodsId
	ai.TplName = "image/addImage.html"
}

func (ui *ImageUploadController) Post() {
	f, h, err := ui.GetFile("upload")
	createTime := time.Now().Format("2006-01-02-15-04-05")
	goodsId := ui.GetString("goodsId")
	IntGoodsId, _ := strconv.ParseInt(goodsId, 10, 64)
	o := orm.NewOrm()
	image := models.Image{GoodsId: goodsId, Createtime: createTime, FilePath: "/static/img/",
		FileName: createTime + h.Filename}
	o.Insert(&image)
	goods := models.Goods{GoodsId: IntGoodsId, ImageUrl: image.FilePath + image.FileName}
	o.Update(&goods, "imageUrl")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		ui.SaveToFile("upload", "./static/img/"+createTime+h.Filename)
		ui.Ctx.WriteString("添加成功！")
	}
}
