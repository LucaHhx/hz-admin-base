package config

type System struct {
	DbType         string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType        string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix   string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr           int    `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	ApiAddr        int    `mapstructure:"api-addr" json:"api-addr" yaml:"api-addr"` // Api端口值
	LimitCountIP   int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP    int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint  bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`    // 多点登录拦截
	UseRedis       bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	UseMongo       bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                   // 使用mongo
	UseStrictAuth  bool   `mapstructure:"use-strict-auth" json:"use-strict-auth" yaml:"use-strict-auth"` // 使用树形角色分配模式
	TimeZone       string `mapstructure:"time-zone" json:"time-zone" yaml:"time-zone"`                   // 时区
	Migration      bool   `mapstructure:"migration" json:"migration" yaml:"migration"`                   // 是否迁移数据库
	TranslationDir string `mapstructure:"translation-dir" json:"translation-dir" yaml:"translation-dir"` // 翻译文件目录
	Environment    string `mapstructure:"environment" json:"environment" yaml:"environment"`             // 环境
}
