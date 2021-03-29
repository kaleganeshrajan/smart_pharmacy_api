package app

import (
	"fmt"
	"strconv"

	cr "github.com/brkelkar/common_utils/configreader"
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

//StartApplication Entry point of service
func StartApplication(cfg cr.Config) {
	fmt.Println("Calling mapsUrls ")
	mapUrls()
	//port := strconv.Itoa(cfg.Server.Port)
	// s := &http.Server{
	// 	Addr:           cfg.Server.Host + ":8088",
	// 	Handler:        r,
	// 	MaxHeaderBytes: 1 << 32,
	// }
	port := strconv.Itoa(cfg.Server.Port)
	Swaggerinit(cfg.Server.Host, port)
	r.Run(cfg.Server.Host + ":" + port)
}
