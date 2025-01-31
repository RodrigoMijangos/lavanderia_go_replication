package lavanderia

import (
	"bufio"
	"encoding/json"
	"fmt"
)

var tanque_agua int = 0

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

func get_energy_and_use(reader *bufio.Reader) {

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				//ctx.JSON(http.StatusOK, map[string]string{"message": "completado"})
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

		//fmt.Printf("Energia recibida: %d", payload["energia"])
	}
}
