package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App           AppAccount
	Routes        RoutesAccount
	Connection    ConnectionAccount
	Logger        LoggerAccount
	Authorization AuthorizationAccount
	Redis         RedisClient
	Cookie        CookieAccount
}
type CookieAccount struct {
	PICDomain     string `json:",omitempty"`
	PICAT         string `json:",omitempty"`
	PICRT         string `json:",omitempty"`
	PICSigned     string `json:",omitempty"`
	CMSDomain     string `json:",omitempty"`
	CMSAT         string `json:",omitempty"`
	CMSRT         string `json:",omitempty"`
	CMSSigned     string `json:",omitempty"`
	PATIENTDomain string `json:",omitempty"`
	PATIENTAT     string `json:",omitempty"`
	PATIENTRT     string `json:",omitempty"`
	PATIENTSigned string `json:",omitempty"`
}

type AppAccount struct {
	Name            string
	ApiGwName       string
	ApplicationName string
	Endpoint        string
	Port            string
	Env             string
	SSL             string
	BodyLimit       int
	HexaSecretKey   string
	BaseUrl         string
	BaseUrlLocal    string
	IsProduction    bool
}
type ConnectionAccount struct {
	DatabaseApp DatabaseAccount
}

type DatabaseAccount struct {
	DriverName      string
	DriverSource    string
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type LoggerAccount struct {
	Logrus    LogrusAccount
	ZapLogger ZapLoggerAccount
}

type LogrusAccount struct {
	Level string
}

type ZapLoggerAccount struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type AuthorizationAccount struct {
	JWT   JWTAccount
	Basic BasicAccount
}

type JWTAccount struct {
	AccessTokenSecretKey  string
	AccessTokenDuration   time.Duration
	RefreshTokenSecretKey string
	RefreshTokenDuration  time.Duration
}

type BasicAccount struct {
	ApiKey    string
	ApiSecret string
}

type RedisClient struct {
	Host     string
	Password string
	Db       int
	RedisKey string
}

type RoutesAccount struct {
	Methods string
	Headers string
	Origins OriginsDetail
}

type OriginsDetail struct {
	IsDefault  bool
	FeLocal    string
	FeDevCMS   string
	FeProdCMS  string
	FeDevCore  string
	FeProdCore string
}

// =========================================================================================================
func InitConf(env string) *Config {
	configPath := GetConfigPath(env)

	confFile, err := LoadConfig(configPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	conf, err := ParseConfig(confFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return conf
}

func GetConfigPath(configPath string) string {
	return "./config/config-dev"
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
