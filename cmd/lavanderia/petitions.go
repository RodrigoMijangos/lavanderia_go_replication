package lavanderia

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func do_petition_for_water(quantity int) {

	for {
		resp, err := http.Get(`http://localhost:4000/get_water/` + strconv.Itoa(quantity))

		if err != nil {
			fmt.Printf("Error con petición: %s\n", err.Error())
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Respuesta inesperada: %d\n", resp.StatusCode)
		}

		reader := bufio.NewReader(resp.Body)

		get_water_and_store(reader)

		time.Sleep(10 * time.Second)
	}

}

func do_petition_for_energy(quantity string, ctx *gin.Context) {
	resp, err := http.Get("http://localhost:4002/electricidad/" + quantity)

	if err != nil {
		fmt.Printf("Error con petición: %s\n", err.Error())
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Respuesta inesperada: %d\n", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)

	get_energy_and_use(reader)

}
