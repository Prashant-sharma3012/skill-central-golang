package model

import "time"

type SqlSettings struct {
	DriverName   string `json:"DriverName"`
	DatabaseURL  string `json:"DataSource"`
	DatabaseName string `json:"DatabaseName"`
	MaxOpenConns int    `json:"MaxOpenConns"`
	MaxIdleConns int    `json:"MaxIdleConns"`
}

type LogSettings struct {
	LogLevel  string `json:"LgLevel"`
	EnableLog bool   `json:"EnableLog"`
}
type ServerSettings struct {
	HostDomain       string        `json:"HostDomain"`
	UseHTTP          bool          `json:"UseHTTP"`
	UseHTTPS         bool          `json:"UseHTTPS"`
	HTTPPort         int           `json:"HTTPPort"`
	HTTPSPort        int           `json:"HTTPSPort"`
	CertFile         string        `json:"CertFile"`
	KeyFile          string        `json:"KeyFile"`
	ConfigRefresh    time.Duration `json:"ConfigRefresh" id:"config_refresh"`
	EnableAccessLogs bool          `json:"EnableAccessLogs" id:"enable_access_logs"`
	ReadTimeout      time.Duration `json:"ReadTimeout" id:"read_timeout"`
	WriteTimeout     time.Duration `json:"WriteTimeout" id:"write_timeout"`
	IdleTimeout      time.Duration `json:"IdleTimeout" id:"idle_timeout"`
}

type Config struct {
	SqlSettings    *SqlSettings
	LogSettings    *LogSettings
	ServerSettings *ServerSettings
}
