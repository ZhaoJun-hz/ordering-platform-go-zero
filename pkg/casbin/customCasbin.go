package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormaAdapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"sync"
)

// Initialize the model from a string.
var modelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

var (
	enforcer *casbin.SyncedEnforcer
	once     sync.Once
)

func Init(db *gorm.DB) *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, err := gormaAdapter.NewAdapterByDBUseTableName(db, "sys", "casbin_rule")
		if err != nil {
			panic(err)
		}
		m, err := model.NewModelFromString(modelText)
		if err != nil {
			panic(err)
		}
		enforcer, err = casbin.NewSyncedEnforcer(m, adapter)
		if err != nil {
			panic(err)
		}
		err = enforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}
	})
	return enforcer
}
