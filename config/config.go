/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package config

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	ApiPort            string `env:"API_PORT"`
	DbUri              string `env:"DB_URI"`
	JwtTokenSecret     string `env:"JWT_TOKEN_SECRET"`
	JwtTokenExpMinutes int    `env:"JWT_TOKEN_EXP_MINUTES"`
}

// ENV - output variable
var ENV Environment

func init() {
	gotenv.Load() // load .env file (if exists)
	if _, err := env.UnmarshalFromEnviron(&ENV); err != nil {
		log.Fatal("Fatal error unmarshalling environment config: ", err)
	}
}
