package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

var (
	globalOrm orm.Ormer
	once      sync.Once
)

func init() {
	// init orm tables
	orm.RegisterModel(
		new(Rules),
		new(Plans),
		new(Proms),
		new(Alerts),
		new(Receivers),
		new(Groups),
		new(Manages),
		new(Configs),
		new(Users),
		new(InhibitLog),
		new(InhibitRule),
		new(SourceMatcher),
		new(TargetMatcher),
		new(Silence),
		new(SilenceMatcher),
	)
}

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use orm.NewOrm()
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
