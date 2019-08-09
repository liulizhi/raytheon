package utils

import (
	"fmt"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

// InitCasbin init casbin information
func InitCasbin() error {
	m := casbin.NewModel()
	m.AddDef("r", "r", "sub, dom, obj, act, suf")
	m.AddDef("p", "p", "sub, dom, obj, act, suf")
	m.AddDef("g", "g", "_, _, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub, r.dom) &&"+
		"r.dom == p.dom && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act")

	a := gormadapter.NewAdapter(DBTYPE, DBConnUrl, true)
	CasbinEnforcer = casbin.NewEnforcer(m, a)
	if CasbinEnforcer == nil {
		return fmt.Errorf("casbin init err")
	}
	err := CasbinEnforcer.LoadPolicy()
	return err
}
