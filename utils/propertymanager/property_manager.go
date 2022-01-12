package propertymanager

import (
	"flag"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var configDir string

func InitializeConfig() error {
	flag.StringVar(&configDir, "config.dir", "./conf", "Configuration directory")
	flag.Parse()

	viper.SetConfigName("application")
	viper.AddConfigPath(configDir)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	if err := viper.ReadInConfig(); err != nil{
		return err
	}
	log.Println("Configurations loaded successfully...")
	return nil
}

func GetStringProperty(key string, defaultValue ...string) string {
	keyValue := viper.GetString(key)
	if keyValue == "" && len(defaultValue) != 0 {
		return defaultValue[0]
	}
	return keyValue
}

func GetIntProperty(key string, defaultValue ...int) int {
	keyValue := viper.GetInt(key)
	if keyValue == 0 && len(defaultValue) != 0 {
		return defaultValue[0]
	}
	return keyValue
}

func GetBoolProperty(key string) bool {
	return viper.GetBool(key)
}
