package config

import (
	"fmt"
	"os"
)

var (
	DatabaseStringConnection string = ""
	DatabaseHost             string = ""
	DatabasePort             string = ""

	APIPort uint64 = 9091
)

func Load() {
	DatabaseHost = os.Getenv("DB_HOST")
	DatabasePort = os.Getenv("DB_PORT")
	DatabaseStringConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		DatabaseHost,
		DatabasePort,
		os.Getenv("DB_NAME"),
	)
}
