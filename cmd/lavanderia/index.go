package lavanderia

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Do_after_init() {
	do_petition_for_water(200)
}

type Payload map[string]string

/*
	status: 0 -> lavar
	status: 1 -> centrifugar
	status: 2 -> enjuagar
	status: 3 -> secar
	status: 4 -> terminado

*/

func Run() {
	s := gin.Default()

	s.GET("/servicio/:quantity", handleService)

	srv2 := &http.Server{
		Addr:         ":4001",
		Handler:      s,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv2.ListenAndServe(); err != nil {
		fmt.Println("Error: Server LAVANDERIA hasn't begin")
	}
}
