package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Update          time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册模型
	orm.RegisterModel(new(Category), new(Topic))
	//注册默认数据库（beego中sqlite和mysql都是默认注册过了的  而且地址和端口号都已经配置好了 所以下面的连接中不需要在写localhost和3306等  其实此处的驱动注册也可不写）
	orm.RegisterDataBase("default", "mysql", "root:1234@/test?charset=utf8")
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate) // 此处穿进去的必须是指针
	if err == nil {                           //表示找到此名称的category  返回nil
		return false
	}
	_, err = o.Insert(cate)
	if err != nil {
		return false
	}
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(id string) error {
	o := orm.NewOrm()
	cid, err := strconv.ParseInt(id, 10, 64)
	cate := Category{Id: id}
	_, err = o.Delete(cate)
	return err
}
