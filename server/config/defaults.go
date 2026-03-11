package config

import "github.com/google/uuid"

// SetDefaults 为所有未设置的配置项填充默认值
// 在 Viper Unmarshal 之后调用，确保即使配置文件缺少某些字段，系统也能正常运行
func (s *Server) SetDefaults() {
	s.System.setDefaults()
	s.JWT.setDefaults()
	s.Zap.setDefaults()
	s.Captcha.setDefaults()
	s.AutoCode.setDefaults()
	s.Local.setDefaults()
	s.Excel.setDefaults()

	// 根据 db-type 设置对应数据库的默认值
	switch s.System.DbType {
	case "sqlite":
		s.Sqlite.setDefaults()
	case "mysql":
		s.Mysql.setDefaults()
	case "pgsql":
		s.Pgsql.setDefaults()
	case "mssql":
		s.Mssql.setDefaults()
	case "oracle":
		s.Oracle.setDefaults()
	}

	s.Redis.setDefaults()
}

// ==================== System ====================

func (c *System) setDefaults() {
	if c.DbType == "" {
		c.DbType = "sqlite"
	}
	if c.OssType == "" {
		c.OssType = "local"
	}
	if c.Addr == 0 {
		c.Addr = 9688
	}
	if c.ApiAddr == 0 {
		c.ApiAddr = 9689
	}
	if c.LimitCountIP == 0 {
		c.LimitCountIP = 15000
	}
	if c.LimitTimeIP == 0 {
		c.LimitTimeIP = 3600
	}
	if c.TimeZone == "" {
		c.TimeZone = "UTC"
	}
	if c.TranslationDir == "" {
		c.TranslationDir = "./translation"
	}
	if c.Environment == "" {
		c.Environment = "dev"
	}
	if c.LoginMode == "" {
		c.LoginMode = "simple"
	}
}

// ==================== JWT ====================

func (c *JWT) setDefaults() {
	if c.SigningKey == "" {
		c.SigningKey = uuid.New().String()
	}
	if c.ExpiresTime == "" {
		c.ExpiresTime = "7d"
	}
	if c.BufferTime == "" {
		c.BufferTime = "1d"
	}
	if c.Issuer == "" {
		c.Issuer = "hab"
	}
}

// ==================== Zap ====================

func (c *Zap) setDefaults() {
	if c.Level == "" {
		c.Level = "info"
	}
	if c.Prefix == "" {
		c.Prefix = "[hab]"
	}
	if c.Format == "" {
		c.Format = "console"
	}
	if c.Director == "" {
		c.Director = "log"
	}
	if c.EncodeLevel == "" {
		c.EncodeLevel = "LowercaseColorLevelEncoder"
	}
	if c.StacktraceKey == "" {
		c.StacktraceKey = "stacktrace"
	}
}

// ==================== Captcha ====================

func (c *Captcha) setDefaults() {
	if c.KeyLong == 0 {
		c.KeyLong = 6
	}
	if c.ImgWidth == 0 {
		c.ImgWidth = 240
	}
	if c.ImgHeight == 0 {
		c.ImgHeight = 80
	}
	// OpenCaptcha 默认 0 表示总是需要验证码，0 是有效值所以不做默认
	if c.OpenCaptchaTimeOut == 0 {
		c.OpenCaptchaTimeOut = 3600
	}
}

// ==================== AutoCode ====================

func (c *Autocode) setDefaults() {
	if c.Web == "" {
		c.Web = "web/src"
	}
	if c.Server == "" {
		c.Server = "server"
	}
	// Module 由 initialize.OtherInit() 从 go.mod 读取，这里不做默认
}

// ==================== Local (文件存储) ====================

func (c *Local) setDefaults() {
	if c.Path == "" {
		c.Path = "uploads/file"
	}
	if c.StorePath == "" {
		c.StorePath = "uploads/file"
	}
}

// ==================== Excel ====================

func (c *Excel) setDefaults() {
	if c.Dir == "" {
		c.Dir = "./resource/excel/"
	}
}

// ==================== GeneralDB 通用数据库默认值 ====================

func (c *GeneralDB) setGeneralDefaults() {
	if c.LogMode == "" {
		c.LogMode = "info"
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 100
	}
	if c.Engine == "" {
		c.Engine = "InnoDB"
	}
}

// ==================== SQLite ====================

func (c *Sqlite) setDefaults() {
	if c.Dbname == "" {
		c.Dbname = "data"
	}
	c.setGeneralDefaults()
}

// ==================== MySQL ====================

func (c *Mysql) setDefaults() {
	if c.Path == "" {
		c.Path = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	if c.Dbname == "" {
		c.Dbname = "hab"
	}
	if c.Username == "" {
		c.Username = "root"
	}
	if c.Config == "" {
		c.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	}
	c.setGeneralDefaults()
}

// ==================== PostgreSQL ====================

func (c *Pgsql) setDefaults() {
	if c.Path == "" {
		c.Path = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "5432"
	}
	if c.Dbname == "" {
		c.Dbname = "hab"
	}
	if c.Username == "" {
		c.Username = "postgres"
	}
	if c.Config == "" {
		c.Config = "sslmode=disable TimeZone=UTC"
	}
	c.setGeneralDefaults()
}

// ==================== MSSQL ====================

func (c *Mssql) setDefaults() {
	if c.Path == "" {
		c.Path = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "1433"
	}
	if c.Dbname == "" {
		c.Dbname = "hab"
	}
	if c.Username == "" {
		c.Username = "sa"
	}
	c.setGeneralDefaults()
}

// ==================== Oracle ====================

func (c *Oracle) setDefaults() {
	if c.Path == "" {
		c.Path = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "1521"
	}
	if c.Dbname == "" {
		c.Dbname = "hab"
	}
	if c.Username == "" {
		c.Username = "system"
	}
	c.setGeneralDefaults()
}

// ==================== Redis ====================

func (c *Redis) setDefaults() {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:6379"
	}
}
