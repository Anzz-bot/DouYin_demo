/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 18:12:11
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 18:12:11
 */
package bootstrap

import (
	"fmt"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitializeConfig() *viper.Viper {
	//read config path
	config := "config.yaml"

	//Modify the configuration file through environment variables
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	//initialize viper server
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	//listen conf file
	v.WatchConfig()

	//hot reload
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)

		//unmarshal to original yaml
		if err := v.Unmarshal(&global.App.Config); err != nil {
			panic(err)
		}

	})

	//if environment variable is available
	//unmarshal to global
	if err := v.Unmarshal(&global.App.Config); err != nil {
		panic(err)
	}

	return v
}
