package config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Auth     Auth     `mapstructure:"auth"`
	DB       DB       `mapstructure:"db"`
	Redis    Redis    `mapstructure:"redis"`
	CronJobs CronJobs `mapstructure:"cronjobs"`
}

type Server struct {
	HTTPPort int    `mapstructure:"http_port"`
	Host     string `mapstructure:"host"`
}
type Auth struct {
	SecretToken            string `mapstructure:"token_secret"`
	TokenExpMinutes        uint   `mapstructure:"token_exp_minutes"`
	RefreshTokenExpMinutes uint   `mapstructure:"refresh_token_exp_minutes"`
}

type DB struct {
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DBName string `mapstructure:"db_name"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"pass"`
}

type CronJobs struct {
	UpdateOrdersStatusJob UpdateOrdersStatusJob `mapstructure:"update_orders_status"`
}

type UpdateOrdersStatusJob struct {
	Hour int `mapstructure:"hour"`
}
