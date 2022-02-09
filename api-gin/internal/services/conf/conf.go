package conf

import (
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigVp struct {
	CommonVp   *viper.Viper
	DbVp       *viper.Viper
	JwtVp      *viper.Viper
	SettingsVp *viper.Viper
}

var Env string
var Conf *ConfigVp
var Common *CommonConf
var Db map[string]*DbConf
var Jwt *JwtConf
var Settings *SettingsConf

var once sync.Once

func InitConf(path string, env string) {
	once.Do(func() {
		Conf = &ConfigVp{}
		Env = env
		var err error
		// common conf
		Conf.CommonVp, err = NewConfVp(Env, "common")
		if err != nil {
			log.Fatalf("common conf vp init err: %v", err)
		}
		Conf.CommonVp.Unmarshal(&Common)

		// db conf
		Conf.DbVp, err = NewConfVp(Env, "db")
		if err != nil {
			log.Fatalf("db conf vp init err: %v", err)
		}
		Db = make(map[string]*DbConf)
		Conf.DbVp.Unmarshal(&Db)

		// jwt conf
		Conf.JwtVp, err = NewConfVp(Env, "jwt")
		if err != nil {
			log.Fatalf("jwt conf vp init err: %v", err)
		}
		Conf.JwtVp.Unmarshal(&Jwt)

		// settings conf (will watch)
		Conf.SettingsVp, err = NewConfVp(Env, "settings")
		if err != nil {
			log.Fatalf("settings conf vp init err: %v", err)
		}
		Conf.SettingsVp.Unmarshal(&Settings)
		Conf.SettingsVp.OnConfigChange(func(_ fsnotify.Event) {
			fmt.Print("configs/" + Env + "/settings.json chenged !\n")
			Conf.SettingsVp.Unmarshal(&Settings)
		})
		Conf.SettingsVp.WatchConfig()

		fmt.Print("Conf init success! \n")
	})
}

func NewConfVp(env string, file string) (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName(file)
	vp.AddConfigPath("configs/" + env + "/")
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vp, nil
}
