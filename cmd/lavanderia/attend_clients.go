package lavanderia

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var tanque_agua int = 0

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

func get_water_and_store(reader *bufio.Reader) {

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error leyendo chunk:", err)
			return
		}

		var payload map[string]int
		if err := json.Unmarshal(line, &payload); err != nil {
			fmt.Println("Error al obtener json:", err)
			return
		}

		tanque_agua += payload["water"]

	}

}
