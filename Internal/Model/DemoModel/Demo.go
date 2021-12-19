package DemoModel

import (
	"github.com/bingxindan/bxd_go_lib/db"
	"xorm.io/core"
)

type Article struct {
	Id        int    `xorm:"not null pk autoincr INT(10)"`
	ArticleSn string `xorm:"not null default '' comment('编码') VARCHAR(200)"`
	Title     string `xorm:"not null default '' comment('标题') VARCHAR(255)"`
	Author    string `xorm:"not null default '' comment('作者') VARCHAR(200)"`
	CreateAt  string `xorm:"not null default '' comment('创建时间') VARCHAR(50)"`
	UpdatedAt string `xorm:"not null default '' comment('更新时间') TINYINT(2)"`
	UserId    int    `xorm:"not null default 0 comment('用户ID') INT(10)"`
}

var (
	TableArticle = "bxd_article"
)

type ArticleDao struct {
	db.DbBaseDao
}

func NewArticleDao(v ...interface{}) *ArticleDao {
	this := new(ArticleDao)
	if ins := db.GetDbInstance("blog", "writer"); ins != nil {
		this.UpdateEngine(ins.Engine)
	} else {
		return nil
	}
	if len(v) != 0 {
		this.UpdateEngine(v...)
	}
	this.Engine.ShowSQL(true)
	this.Engine.Logger().SetLevel(core.LOG_DEBUG)
	return this
}

func (this *ArticleDao) GetById(id db.Param) (response Article, err error) {
	_, err = this.Engine.Table(TableArticle).
		Where("id=?", id).Get(&response)

	if err != nil {
		return response, err
	}

	return response, nil
}
