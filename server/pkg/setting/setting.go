package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	Config       map[string]interface{}
	RunMode      string
	HTTPPort     int64
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int64
	JwtSecret string
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	Config = viper.AllSettings()
	// map[
	// 	app:map[
	// 		jwt_secret:nuonuo534 page_size:10
	// 	]
	// 	database:map[
	// 		host:127.0.0.1:3306 name:blog password:nuonuo534 type:mysql user:root
	// 	]
	// 	run_mode:debug
	// 	server:map[
	// 		http_port:8000
	// 		read_timeout:60
	// 		write_timeout:60
	// 	]
	// ]
	fmt.Println(Config)
	LoadBase()
	LoadSever()
	LoadApp()
}

func LoadBase() {
	RunMode = Config["run_mode"].(string)
}

func LoadSever() {
	server := Config["server"].(map[string]interface{})
	HTTPPort = server["http_port"].(int64)
	ReadTimeout = time.Duration(server["read_timeout"].(int64)) * time.Second

	WriteTimeout = time.Duration(server["write_timeout"].(int64)) * time.Second
}

func LoadApp() {
	app := Config["app"].(map[string]interface{})

	JwtSecret = app["jwt_secret"].(string)
	PageSize = app["page_size"].(int64)
}
