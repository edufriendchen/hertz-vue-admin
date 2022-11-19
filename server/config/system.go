package config

type System struct {
	Env                   string `mapstructure:"env" json:"env" yaml:"env"`                                                                // 环境值
	Addr                  int    `mapstructure:"addr" json:"addr" yaml:"addr"`                                                             // 端口值
	DbType                string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                                                    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType               string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                                                 // Oss类型
	SinglePointLogin      bool   `mapstructure:"single-point-login" json:"single-point-login" yaml:"single-point-login"`                   // 是否单点登录
	MaxMultipointLoginNum int    `mapstructure:"max-multipoint-login-num" json:"max-multipoint-login-num" yaml:"max-multipoint-login-num"` // 最大多点登录数量
	UseRedis              bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                                              // 使用redis
	LimitCountIP          int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP           int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
