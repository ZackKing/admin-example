package conf

import (
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigVp struct {
	Path    string
	AppVp   *viper.Viper
	DbVp    *viper.Viper
	JwtVp   *viper.Viper
	RedisVp *viper.Viper
}

var Env string
var Conf *ConfigVp
var App *AppConf
var Db map[string]*DbConf
var Jwt *JwtConf
var Redis map[string]*RedisConf

var once sync.Once

func InitConf(path string, env string) {
	once.Do(func() {
		Conf = &ConfigVp{Path: path}
		Env = env

		// conf init (will watch)
		Conf.AppVp = NewConfVp(Conf.Path, "app", &App)
		Conf.AppVp.OnConfigChange(func(_ fsnotify.Event) {
			fmt.Printf("configs/%s/app.json chenged !", Env)
			Conf.AppVp.Unmarshal(&App)
		})
		Conf.AppVp.WatchConfig()

		// conf init (not watch)
		Conf.DbVp = NewConfVp(Conf.Path, "db", &Db)
		Conf.RedisVp = NewConfVp(Conf.Path, "redis", &Redis)
		Conf.JwtVp = NewConfVp(Conf.Path, "jwt", &Jwt)

		log.Default().Print("Conf initialized")
	})
}

func NewConfVp(path string, file string, vars any) *viper.Viper {
	vp := viper.New()
	vp.AddConfigPath(path + "/" + Env + "/")
	vp.SetConfigName(file)
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("%s conf vp init err: %v", file, err))
	}
	vp.Unmarshal(&vars)
	return vp
}
