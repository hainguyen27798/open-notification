package setting

type Config struct {
	Server ServerSettings `mapstructure:"server"`
	Logger LoggerSettings `mapstructure:"log"`
	SMTP   SMTPSettings   `mapstructure:"smtp"`
}

type ServerSettings struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type SMTPSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	From     string `mapstructure:"from"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type LoggerSettings struct {
	FileName   string `mapstructure:"file_name"`
	Level      string `mapstructure:"log_level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
