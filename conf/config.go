package conf

//全局config全局对象，就是程序在内存中配置对象，程序内部通过读取该对象
//什么时候初始化  == 》 配置加载时候
//LoadConfigFromTEnv  LoadConfigFromToml
// 为了不被程序运行时恶意修改，设置成私有变量
var config *Config

//若想获取配置，单独提供函数
//全局config对象获取函数

func C() *Config {
	return config
}
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

// Config config 应用配置，通过封装一个对象，来与外部配置进行对接
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

//func (a *App) HttpAddr() string {
//	return fmt.Sprintf("%s:%s", a.Host, a.Port)
//}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "",
		Database:    "go_course",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的MySQL连接池, 需要池做一些规划配置
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	//控制mysql的复用 比如5 最多运行5个复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期 这个跟 mysql 的server 的配置有关系， 例子：一个连接12h，保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	//idle的最多允许 存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
}

func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}
