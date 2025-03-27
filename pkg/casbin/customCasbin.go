package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormaAdapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ordering-platform/pkg/common/conf"
	"sync"
	"time"
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
	enforcer *casbin.Enforcer
	once     sync.Once
)

func Init(datasource string, redisConf redis.RedisConf) *casbin.Enforcer {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 settingLogConfig(),
		})
		if err != nil {
			panic(err)
		}
		adapter, err := gormaAdapter.NewAdapterByDBUseTableName(db, "sys", "casbin_rule")
		if err != nil {
			panic(err)
		}
		m, err := model.NewModelFromString(modelText)
		if err != nil {
			panic(err)
		}
		enforcer, err = casbin.NewEnforcer(m, adapter)
		if err != nil {
			panic(err)
		}
		err = enforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}

		w, err := rediswatcher.NewWatcher(redisConf.Host, rediswatcher.WatcherOptions{
			Options: redis2.Options{
				Network:  "tcp",
				Password: redisConf.Pass,
			},
			Channel:    fmt.Sprintf("%s", conf.RedisCasbinChannel),
			IgnoreSelf: false,
		})
		logx.Must(err)

		err = w.SetUpdateCallback(func(data string) {
			rediswatcher.DefaultUpdateCallback(enforcer)(data)
		})
		err = enforcer.SetWatcher(w)
		logx.Must(err)
		err = enforcer.SavePolicy()
		logx.Must(err)
	})
	return enforcer
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
