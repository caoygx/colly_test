package models

type Article struct {
	Id             int    `xorm:"not null pk autoincr comment('') INT(11)" json:"id"`
	Title          string `xorm:"comment('标题')" json:"title"`
	Style          string `xorm:"comment('标题颜色')" json:"style"`
	Flags          int    `xorm:"comment('属性')" json:"flags"`
	Des            int    `xorm:"comment('排序')" json:"des"`
	Tags           string `xorm:"comment('TAG标签')" json:"tags"`
	Image          string `xorm:"comment('缩略图')" json:"image"`
	Typeid         int    `xorm:"comment('栏目')" json:"typeid"`
	Click          int    `xorm:"comment('访问次数')" json:"click"`
	Keywords       string `xorm:"comment('关键词')" json:"keywords"`
	Description    string `xorm:"comment('内容描述')" json:"description"`
	Arcrank        int    `xorm:"comment('阅读权限')" json:"arcrank"`
	Template       string `xorm:"comment('模板文件')" json:"template"`
	CreateTime     string `xorm:"comment('添加时间')" json:"create_time"`
	UpdateTime     string `xorm:"comment('更新时间')" json:"update_time"`
	Content        string `xorm:"comment('')" json:"content"`
	Hide           int    `xorm:"comment('0隐藏 1显示文章 2审核 3拒绝')" json:"hide"`
	Url            string `xorm:"comment('跳转')" json:"url"`
	Source         string `xorm:"comment('来源')" json:"source"`
	Uid            int    `xorm:"comment('会员UID 识别注明哪个会员投稿')" json:"uid"`
	Writer         string `xorm:"comment('文章作者')" json:"writer"`
	Mychannel      int    `xorm:"comment('频道模型ID')" json:"mychannel"`
	Images         string `xorm:"comment('图片集，如果mytype值是2，则是图片文章')" json:"images"`
	Video          string `xorm:"comment('视频')" json:"video"`
	VideoType      string `xorm:"comment('第三方视频类型')" json:"videoType"`
	Videodate      string `xorm:"comment('视频时间')" json:"videodate"`
	Zan            int    `xorm:"comment('')" json:"zan"`
	Cai            int    `xorm:"comment('')" json:"cai"`
	PingNum        int    `xorm:"comment('评论数量')" json:"pingNum"`
	Weitoutiao     int    `xorm:"comment('1微头条 2 首页+微头条 3 贴吧 4 小视频广告')" json:"weitoutiao"`
	ToutiaoTiemId  string `xorm:"comment('今日头条原文章ID')" json:"toutiao_tiemId"`
	QiniuVideo     string `xorm:"comment('')" json:"qiniu_video"`
	QiniuVideoType int    `xorm:"comment('1转码中 0转码成功')" json:"qiniu_video_type"`
	Zhiding        int    `xorm:"comment('1置顶')" json:"zhiding"`
	Tuijian        int    `xorm:"comment('推荐')" json:"tuijian"`
	Zsort          int    `xorm:"comment('排序')" json:"zsort"`
	Stuijian       int    `xorm:"comment('搜索推荐')" json:"stuijian"`
	Fangshi        string `xorm:"comment('0采集，1用户')" json:"fangshi"`
	ReplyTime      int    `xorm:"comment('最新回复时间')" json:"reply_time"`
}
