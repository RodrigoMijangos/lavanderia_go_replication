package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	connections = make(chan struct{}, 10)
)

func getWaterHandler(c *gin.Context) {

	println("Received a Request!!")

	quantity, exists := c.Params.Get("quantity")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The param quantity is missing"})
		return
	}

	count, err := strconv.Atoi(quantity)

	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity param must be a positive number"})
		return
	}

	connections <- struct{}{}
	defer func() {
		<-connections
	}()

	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Ecnoding", "chunked")

	writer := c.Writer

	sendWater(count, c, writer)

}

func sendWater(quantity int, c *gin.Context, writer gin.ResponseWriter) {
	done := c.Request.Context().Done()

	for i := 1; i <= quantity; i++ {
		select {
		case <-done:
			return
		default:
			deliver := map[string]int{"water": 10}

			println("Sending Water: 10u")

			jsonData, err := json.Marshal(deliver)
			if err != nil {
				println("Error generating JSON: %s", err.Error())
				return
			} else {
				writer.Write(jsonData)
				writer.Write([]byte("\n"))
				writer.Flush()
			}
		}

		time.Sleep(500 * time.Millisecond)

	}

	message := map[string]string{"message": "All data have been sent"}
	jsonData, _ := json.Marshal(message)
	writer.Write(jsonData)
	writer.Flush()
}

func main() {
	r := gin.Default()

	r.GET("/get_water/:quantity", getWaterHandler)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server hasn't begin")
	}
}
