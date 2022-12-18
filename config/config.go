package config

import "time"

type Config struct {
	Appname  string
	Version  string
	Port     int
	Database Database
	Logger   Logger
	Server   Server
}

type Database struct {
	Driver          string
	DSN             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime int
}

type Logger struct {
	Name         string
}

type Server struct {
	ReadTimeOut, WriteTimeOut, ShutdownDelay time.Duration
}
