package config

import "os"

var (
	instance *Connection
	ENVVARS  = os.Getenv("VARS")
)

type (
	Connection struct {
		//some connection to another apps/DB
		C string
	}
)

func init() {
	if &instance != nil {
		return
	}

	//some initialitation

	instance = &Connection{
		C: "someConnection",
	}
}
func GetInstance() *Connection {
	return instance
}
