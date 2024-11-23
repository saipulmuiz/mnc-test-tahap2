package config

import (
	"os"
	"time"
)

var JWTConfig = struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}{
	AccessSecret:  os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
	RefreshSecret: os.Getenv("REFRESH_TOKEN_SECRET_KEY"),
	AccessExpiry:  time.Minute * 15,
	RefreshExpiry: time.Hour * 24,
}
