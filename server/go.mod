module github.com/nuonuo534/doing

go 1.14

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools/gopls v0.6.6 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
	server/algorithm v1.0.0 //
	server/api v0.0.0
	server/e v0.0.0
	server/middleware v0.0.0
	server/models v0.0.0
	server/routers v0.0.0
	server/setting v0.0.0
	server/util v0.0.0
)

replace (
	server/algorithm => ./algorithm
	server/api => ./routers/api
	server/e => ./pkg/e
	server/middleware => ./middleware
	server/models => ./models
	server/routers => ./routers
	server/setting => ./pkg/setting
	server/util => ./pkg/util
)
