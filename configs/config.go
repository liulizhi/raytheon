package configs

// TomlConfig base config
type TomlConfig struct {
	MainConfig      MainInfo           `mapstructure:"main"`
	DBConfig        DBInfo             `mapstructure:"db"`
	RedisConfig     RedisInfo          `mapstructure:"redis"`
	LDAPConfig      LDAPInfo           `mapstructure:"ldap"`
	ThirdApisConfig map[string]APIInfo `mapstructure:"apis"`
}

// MainInfo main config
type MainInfo struct {
	LogLevel string `mapstructure:"log_level"`
	Port     int    `mapstructure:"port"`
	LogDir   string `mapstructure:"log_dir"`
}

// DBInfo db conn information
type DBInfo struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Charset      string `mapstructure:"charset"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisInfo redis conn information
type RedisInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// APIInfo apis conn information
type APIInfo struct {
	URL     string `mapstructure:"url"`
	APIUser string `mapstructure:"api_user"`
}

// LDAPInfo ldap information
type LDAPInfo struct {
	Enable   bool   `mapstructure:"enable"`
	Addr     string `mapstructure:"addr"`
	Port     int    `mapstructure:"port"`
	BaseDN   string `mapstructure:"base_dn"`
	BindDN   string `mapstructure:"bind_dn"`
	BindPass string `mapstructure:"bind_pass"`
}
