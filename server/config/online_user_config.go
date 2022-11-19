package config

import "time"

type OnlineUser struct {
	Key       string        `mapstructure:"key" json:"key" yaml:"key"`                       // redis的哪个数据库
	OlineTime time.Duration `mapstructure:"online-time" json:"oline-time" yaml:"oline-time"` // redis的哪个数据库
}
