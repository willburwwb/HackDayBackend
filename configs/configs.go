package configs

import (
	"HackDayBackend/utils"
	"time"
)

type ServerConfigs struct {
	RunMode      string        `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	Addr         string        `mapstructure:"addr" json:"addr" yaml:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
}

type PgsqlConfigs struct {
	Host           string `mapstructure:"host" json:"host" yaml:"host"`                                      // 服务器地址
	Port           string `mapstructure:"port" json:"port" yaml:"port"`                                      // 端口
	Dbname         string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                                // 数据库名
	User           string `mapstructure:"user" json:"user" yaml:"user"`                                      // 用户名
	Password       string `mapstructure:"password" json:"password" yaml:"password"`                          // 密码
	MaxIdleConns   int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`        // 空闲中的最大连接数
	MaxOpenConns   int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`        // 打开到数据库的最大连接数
	MaxLifeSeconds int64  `mapstructure:"max_life_seconds" json:"max_life_seconds" yaml:"max_life_seconds" ` // 数据库连接最长生命周期
}

type RedisConfigs struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

var (
	Pgsql_config  *PgsqlConfigs
	Redis_config  *RedisConfigs
	Server_config *ServerConfigs
)

func init() {
	settings, err := NewSettings()
	if err != nil {
		utils.ErrorF("new settings error: %s", err)
		panic("error from new settings ")
	}

	err = settings.ReadToStruct("Pgsql", &Pgsql_config)
	utils.DebugF("pgsql config: %+v", Pgsql_config)
	if err != nil {
		utils.ErrorF("set ogsql config error: %s", err)
		panic("set pgsql error")
	}

	err = settings.ReadToStruct("Redis", &Redis_config)
	utils.DebugF("redis config: %+v", Redis_config)
	if err != nil {
		utils.ErrorF("set redis config error: %s", err)
		panic("set redis error")
	}

	err = settings.ReadToStruct("Server", &Server_config)
	utils.DebugF("server config: %+v", Server_config)
	if err != nil {
		utils.ErrorF("set server config error: %s", err)
		panic("set server error")
	}
}