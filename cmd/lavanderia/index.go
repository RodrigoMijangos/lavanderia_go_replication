package lavanderia

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Do_after_init() {
	do_petition_for_water(200)
}

func Run() {
	s := gin.Default()

	s.GET("/main", func(ctx *gin.Context) {
		payload, _ := json.Marshal((map[string]int{"message": 0}))

		ctx.JSON(http.StatusOK, payload)
	})

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
