package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Article struct {
	Id    int    `xorm:"not null pk autoincr comment('') INT(11)" json:"id"`
	Title string `xorm:"comment('标题')" json:"title"`

	Keywords    string `xorm:"comment('关键词')" json:"keywords"`
	Description string `xorm:"comment('内容描述')" json:"description"`

	Content string `xorm:"comment('')" json:"content"`
	Typeid  int    `xorm:"comment('分类')" json:"typeid"`
}

func (keyword *Article) TableName() string {
	return "fx_article"
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func saveToDb(article Article) {
	db, _ := gorm.Open("mysql", "root:123456@tcp(192.168.140.118:3306)/toutiao?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	fmt.Print(article)
	a := Article{}
	rQuery := db.Where("title = ? and 1", article.Title).Find(&a)
	if rQuery.Error == nil || a.Id > 0 {
		return
	}
	fmt.Print(article)
	result := db.Create(&article)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println(article.Id)
	//os.Exit(1)

}

func getArticleForSq1996(body string) Article {
	//doc, err := htmlquery.LoadDoc(body)
	//fmt.Println(body)
	var article = Article{}
	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		return article
		//panic(err)
	}

	title := htmlquery.FindOne(doc, "//div[@id=\"title\"]")
	if title != nil {
		article.Title = htmlquery.OutputHTML(title, false)
	}

	content := htmlquery.FindOne(doc, "//div[@id=\"content\"]")
	if content != nil {
		article.Content = htmlquery.OutputHTML(content, false)
		//article.Content = "xxx"
	}
	return article

}

func allowVisit(reg string, url string) bool {

	match, _ := regexp.MatchString(reg, url)
	return match
}

func gather1996(category Category) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		//c.Visit(e.Request.URL(link))
		r := allowVisit(`http://news.sq1996.com/.+/\d+/\d+/\d+.shtml`, link) //"http://news.sq1996.com/sqyw/2020/0610/321870.shtml")
		if r == true {
			//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// c.OnHTML(".in_left_list_title", func(e *colly.HTMLElement) {
	// 	fmt.Println("First column of a table row:", e.Text)
	// })

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		//fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		//r2 := rand.New(rand.NewSource(9))

		//fmt.Println(r2.Int())
		//fmt.Println(r.Request.URL)
		//h := md5.New()
		//h.Write([]byte(string(*r.Request.URL)))
		//md5Str := hex.EncodeToString(h.Sum(nil))

		article := getArticleForSq1996(string(r.Body))

		if article.Title != "" && article.Content != "" {
			article.Typeid = category.Id
			saveToDb(article)
			fmt.Println(article.Id)
		}

		// filename := "html/" + strconv.Itoa(r2.Int()) + ".txt"
		// r.Save(filename)

		// fmt.Println()
		//fmt.Println(string(r.Body))
	})

	//c.Visit("http://go-colly.org/")
	c.Visit("http://news.sq1996.com/" + category.En)
	//c.Visit("http://news.sq1996.com/sqyw/2020/0610/321869.shtml")

}

type Category struct {
	Id int
	En string
}

func main() {

	db, _ := gorm.Open("mysql", "root:123456@tcp(192.168.140.118:3306)/toutiao?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	// a := Article{}
	// article := Article{}
	// article.Title = "泗阳县领导干部任职前公示"
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
		gather1996(v)
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
