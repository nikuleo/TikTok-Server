package config

import (
	"TikTokServer/pkg/tlog"

	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

func GetConfig(cfgName string) Config {
	v := viper.New()
	cfg := Config{
		Viper: v,
	}

	viper := cfg.Viper

	viper.SetConfigName(cfgName)
	viper.AddConfigPath(".config")

	if err := viper.ReadInConfig(); err != nil {
		tlog.Fatalf("reading config: %s, failed: %s", cfgName, err.Error())
	}

	tlog.Infof("reading config: %s, success.", viper.ConfigFileUsed())

	return cfg
}
