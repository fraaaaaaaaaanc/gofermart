package config

import (
	"flag"
	"os"
)

func ParseConfFlags() *Flags {
	flags := ParseFlags()
	ParseEnv(flags)
	return flags
}

func ParseEnv(flags *Flags) {
	if host := os.Getenv("RUN_ADDRESS"); host != "" {
		flags.Set(host)
	}

	if secretKeyJwtToken := os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN"); secretKeyJwtToken != "" {
		flags.SecretKeyJWTToken = secretKeyJwtToken
	}

	if dataBaseURI := os.Getenv("DATABASE_URI"); dataBaseURI != "" {
		flags.DataBaseURI = dataBaseURI
	}

	if accrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); accrualSystemAddress != "" {
		flags.AccrualSystemAddress = accrualSystemAddress
	}

	if LogFilePath := os.Getenv("LOG_FILE"); LogFilePath != "" {
		flags.LogFilePath = LogFilePath
	}

	if ProjLvl := os.Getenv("LOG_LVL"); ProjLvl != "" {
		flags.ProjLvl = ProjLvl
	}
}

func ParseFlags() *Flags {
	flags := newFlags()

	flag.Var(&flags.HTTPServerHost, "a", "address and port to run server")
	flag.StringVar(&flags.DataBaseURI, "d", flags.DataBaseURI, "database connection address:host user "+
		"password dbname sslmode")
	flag.StringVar(&flags.AccrualSystemAddress, "r", flags.AccrualSystemAddress,
		"address of the accrual calculation system")
	flag.StringVar(&flags.LogFilePath, "lf", flags.LogFilePath, "the path for the file to which the "+
		"logs will be written")
	flag.StringVar(&flags.ProjLvl, "ll", flags.ProjLvl, "project development stage")

	flag.Parse()
	return &flags
}
