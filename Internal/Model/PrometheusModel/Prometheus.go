package PrometheusModel

import (
	"github.com/bingxindan/bxd_go_lib/db"
	"xorm.io/core"
)

type Prometheus struct {
	Id       int    `xorm:"not null pk autoincr INT(10)" json:"id"`
	PrometheusName string `xorm:"not null default '' comment('用户名称') VARCHAR(100)" json:"userName"`
	NickName string `xorm:"not null default '' comment('用户弥称') VARCHAR(100)" json:"nickName"`
	Password string `xorm:"not null default '' comment('密码') VARCHAR(100)" json:"password"`
	Avatar   string `xorm:"not null default '' comment('头像') VARCHAR(500)" json:"avatar"`
	CreateAt string `xorm:"not null default '' comment('创建时间') VARCHAR(50)" json:"createAt"`
	UpdateAt string `xorm:"not null default '' comment('更新时间') VARCHAR(50)" json:"updateAt"`
}

var (
	TablePrometheus = "jz_user"
)

type PrometheusDao struct {
	db.DbBaseDao
}

func NewPrometheusDao(v ...interface{}) *PrometheusDao {
	this := new(PrometheusDao)
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

func (d *PrometheusDao) GetByIds(ids []string) (response []Prometheus, err error) {
	err = d.Engine.
		Table(TablePrometheus).
		Where("is_delete = ?", 0).
		In("id", ids).
		Find(&response)

	if err != nil {
		return response, err
	}

	return response, nil
}
