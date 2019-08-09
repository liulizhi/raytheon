package utils

import (
	"fmt"
	"raytheon/configs"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var cfgFile string

// APIConfig global api config
var APIConfig configs.TomlConfig

var DBConnUrl string

// DB global db object
var DBConn *gorm.DB

func init() {
	InitConfig()
	err := viper.Unmarshal(&APIConfig)
	if err != nil {
		panic(err)
	}
	DBConnUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		APIConfig.DBConfig.User,
		APIConfig.DBConfig.Password,
		APIConfig.DBConfig.Host,
		APIConfig.DBConfig.Port,
		APIConfig.DBConfig.DB,
		APIConfig.DBConfig.Charset,
	)
	DBConn, err = InitDB(DBConnUrl)
	if err != nil {
		panic(err)
	}
}

// InitConfig init config file
func InitConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".mysql" (without extension).
		viper.AddConfigPath("./configs/")
		viper.AddConfigPath("../configs/")
		viper.SetConfigName("config")
	}

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&APIConfig)
		if err != nil {
			panic(err)
		}
	})
}
