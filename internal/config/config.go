package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"
)

var cfg *viper.Viper
var once sync.Once

func New() *viper.Viper {
	once.Do(func() {
		cfg = viper.New()

		etcd := cfg.Sub("etcd")
		etcd.SetDefault("addr", "http://localhost:2379")
		etcd.SetDefault("prefix", "tiktok-demo-micro")

		redis := cfg.Sub("redis")
		redis.SetDefault("addr", "localhost:6379")
		redis.SetDefault("password", "")

		mysql := cfg.Sub("mysql")
		mysql.SetDefault("addr", "localhost:3306")
		mysql.SetDefault("username", "root")
		mysql.SetDefault("password", "root")
		mysql.SetDefault("dbname", "tiktok-demo-micro")

		kafka := cfg.Sub("kafka")
		kafka.SetDefault("addr", "localhost:9092")

		s3 := cfg.Sub("s3")
		s3.SetDefault("endpoint", "http://localhost:9000")
		s3.SetDefault("accessKey", "")
		s3.SetDefault("secretKey", "")
		s3.SetDefault("bucket", "")
		s3.SetDefault("region", "")

		cfgpath := os.Getenv("TIKTOK_DEMO_MICRO_CONFIG_FILE")
		cfg.AddConfigPath(cfgpath)
		cfg.SetConfigName("config")
		cfg.AddConfigPath(".")
		cfg.AddConfigPath("/etc/tiktok-demo-micro")
		cfg.AddConfigPath("$HOME/.tiktok-demo-micro")
		cfg.SetConfigType("yaml")
		err := cfg.ReadInConfig()
		if err != nil {
			panic(err)
		}
		viper.AddRemoteProvider("etcd", etcd.GetString("addr"), etcd.GetString("prefix"))
		err = viper.ReadRemoteConfig()
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
