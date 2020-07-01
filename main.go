package main

import (
	//"colly_test/models"

	//uuid "github.com/iris-contrib/go.uuid"

	. "github.com/caoygx/golib/models"
	"github.com/caoygx/golib/services"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/qiniu/api.v7/v7/auth/qbox"
	// "github.com/qiniu/api.v7/v7/storage"
)

func main() {
	//utils.UploadImg("http://upload.sq1996.com/2020/0615/1592182789329.jpg")

	// uploadImg()
	// os.Exit(1)

	//db, _ := gorm.Open("mysql", "root:123456@tcp(192.168.140.118:3306)/toutiao?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	//a := models.Article{}
	// article := models.Article{}
	// article.Title = "泗阳县领导干部任职前公示"
	// fmt.Print(article)
	// os.Exit(1)
	// r := db.Where("title = ? and 1", article.Title).Find(&a)
	// if r.Error != nil {
	// 	//save
	// }
	// fmt.Print(r.Error)
	// //fmt.Print(r)
	// return

	var categorys = []Category{
		Category{1, "minsheng"},
		Category{2, "sqyw"},
		Category{3, "bmdt"},
		Category{4, "gnxw"},
		Category{5, "bmdt"},
	}
	for _, v := range categorys {
		//fmt.Println(v)
		services.Gather1996(v)
	}

	//fmt.Print(categorys)
	//
	//http://news.sq1996.com/minsheng/
	/*
			<a href="http://news.sq1996.com/" title="新闻中心" class="home">新闻</a>
		<a href="http://news.sq1996.com/sqyw/" title="宿迁要闻">要闻</a>
		<a href="http://news.sq1996.com/minsheng/" title="宿迁民生新闻">民生</a>
		<a href="http://news.sq1996.com/bmdt/" title="宿迁部门动态">部门</a>
		<a href="http://news.sq1996.com/gnxw/" title="国内">国内</a>
		<a href="http://news.sq1996.com/gjdt/" title="国际">国际</a>
		<a href="http://news.sq1996.com/ylbg/" title="娱乐">娱乐</a>
		<a href="http://news.sq1996.com/tyxw/" title="体育">体育</a>
		<a href="http://www.sqsc.gov.cn/" title="宿城">宿城</a>
		<a href="http://www.suyu.gov.cn/" title="宿豫">宿豫</a>
		<a href="http://www.shuyang.gov.cn/zgsy/" title="沭阳">沭阳</a>
	*/

}
