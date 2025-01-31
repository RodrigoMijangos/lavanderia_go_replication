package cfe

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/electricidad/:quantity", func(ctx *gin.Context) {

		ctx.Header("Content-Type", "application/json")
		ctx.Header("Transfer-Encoding", "chunked")

		quantityParam := ctx.Param("quantity")

		quantity, _ := strconv.Atoi(quantityParam)

		writer := ctx.Writer

		for i := 1; i <= 10; i++ {

			payload := map[string]int{"energia": quantity}

			jsonData, _ := json.Marshal(payload)

			writer.Write(jsonData)
			writer.Write([]byte("\n"))
			writer.Flush()

			time.Sleep(1 * time.Second)
		}

		message, _ := json.Marshal(map[string]string{"message": "Servicio entregado"})

		writer.Write(message)
		writer.Flush()
	})

	srv := &http.Server{
		Addr:         ":4002",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	srv.ListenAndServe()
}
