package config

type AppConf struct {
	App     *App     `mapstructure:"app"`
	Mysql   *Mysql   `mapstructure:"mysql"`
	Redis   *Redis   `mapstructure:"redis"`
	LogConf *LogConf `mapstructure:"log"`
	Auth    *Auth    `mapstructure:"auth"`
}

type App struct {
	Mode    string `mapstructure:"mode"`
	Name    string `mapstructure:"name"`
	Swagger bool   `mapstructure:"swagger"`
}

type Mysql struct {
	Url      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// LogConf app logger config
type LogConf struct {
	Level      string `mapstructure:"level"`
	InfoLog    string `mapstructure:"infoLog"`
	ErrorLog   string `mapstructure:"errorLog"`
	TimeFormat string
	Order      []string
}
type Auth struct {
	Jwt *Jwt `mapstructure:"jwt"`
}

type Jwt struct {
	ExpireTime uint64 `mapstructure:"expireTime"`
	Issuer     string `mapstructure:"issuer"` //签发人
}

type Redis struct {
	Addr        string `mapstructure:"addr"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"DB"`
	PoolSize    int    `mapstructure:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn"`
}
