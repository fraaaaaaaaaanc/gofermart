package config

import (
	"errors"
	"strconv"
	"strings"
)

const (
	hostAddr     = "localhost"
	hostPort     = 8080
	addrDB       = "host=localhost password=1234 dbname=gofermart user=postgres sslmode=disable"
	accrualSA    = "test"
	logFileParth = "C:\\Users\\frant\\go\\go1.21.0\\bin\\pkg\\mod\\github.com\\fraaaaaaaaaanc\\gofermart\\internal\\tmp\\log_project.json"
	logLvlLocal  = "local"
)

type Flags struct {
	AccrualSystemAddress string
	LogFilePath          string
	ProjLvl              string
	DataBaseURI          string
	HTTPServerHost
}

type HTTPServerHost struct {
	hostAddress string
	hostPort    int
}

func newFlags() Flags {
	return Flags{
		HTTPServerHost: HTTPServerHost{
			hostAddress: hostAddr,
			hostPort:    hostPort,
		},
		AccrualSystemAddress: accrualSA,
		DataBaseURI:          addrDB,
		LogFilePath:          logFileParth,
		ProjLvl:              logLvlLocal,
	}
}

func (hs *HTTPServerHost) String() string {
	return hs.hostAddress + ":" + strconv.Itoa(hs.hostPort)
}

func (hs *HTTPServerHost) Set(address string) error {
	if address == "" {
		hs.hostAddress = hostAddr
		hs.hostPort = hostPort
		return nil
	}
	listAddress := strings.Split(address, ":")
	if len(listAddress) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(listAddress[1])
	if err != nil {
		return err
	}
	hs.hostAddress = listAddress[0]
	hs.hostPort = port
	return nil
}
