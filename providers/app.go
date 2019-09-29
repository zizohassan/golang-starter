package providers

import (
	"os"
)

func Run()  {
	/**
	* init gin frame work
	* run default middleware like CROS
	*/
	r := Gin()
	/**
	* setup routes with three default groups
	* admin / auth / visitors
	*/
	Routing(r)
	/**
	* start app on port to change port just edit APP_PORT in
	* env file like :9090
	*/
	if os.Getenv("APP_ENV") == "local"{
		r.Run(":" + os.Getenv("APP_PORT_LOCAL"))
	}else {
		r.Run(":" + os.Getenv("APP_PORT_PRODUCTION"))
	}
}
