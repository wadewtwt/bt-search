module github.com/putyy/ai-share

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.7.2
	github.com/gin-gonic/gin v1.9.1
	github.com/go-redis/redis/v8 v8.11.6-0.20220405070650-99c79f7041fc
	github.com/putyy/gokv v0.0.0-20220303073116-12b2d275431c
	github.com/qiniu/go-sdk/v7 v7.13.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/silenceper/wechat/v2 v2.1.3
	github.com/ugorji/go v1.2.7 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
	mvdan.cc/xurls/v2 v2.4.0
)

replace (
	github.com/putyy/ai-share/app => ./app
	github.com/putyy/ai-share/config => ./config
	github.com/putyy/ai-share/router => ./router
)
