package main

import (
	"github.com/Xe/pawd/cmd/pawd/database"
	"github.com/asdine/storm"
)

type Config struct {
	DBPath          string   `env:"DB_PATH" envDefault:"./var/pawd.db"`
	HyperAccessKey  string   `env:"HYPER_ACCESS_KEY,required"`
	HyperSecretKey  string   `env:"HYPER_SECRET_KEY,required"`
	HTTPSPort       string   `env:"TLS_PORT" envDefault:"443"`
	GRPCTLSBindhost string   `env:"GRPC_TLS_BINDHOST" envDefault:"0.0.0.0:4228"`
	AdminEmails     []string `env:"ADMIN_EMAILS" envSeparator:"."`
}

type Server struct {
	Config
	db *storm.DB

	us database.Users
	tk database.Tokens
}
