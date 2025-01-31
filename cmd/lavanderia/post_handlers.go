package lavanderia

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SimpleRequest map[string]int
type NumberPayload struct {
	Numero int `json:"numero"`
}

var bd []int = []int{}

func transferDataHandler(ctx *gin.Context) {
	bodyReader := bufio.NewReader(ctx.Request.Body)

	defer func() {
		ctx.Request.Body.Close()
	}()

	for {
		line, err := bodyReader.ReadBytes('\n')

		if err != nil {
			if err.Error() == "EOF" {
				ctx.JSON(http.StatusOK, SimpleResponse{"message": "Se han recibido correctamente los datos"})
				println(bd)
				break
			}

			fmt.Println("Error leyendo chunks")
			ctx.JSON(http.StatusInternalServerError, SimpleResponse{"error": "Ocurrió un error leyendo el cuerpo"})
			return
		}

		var payload NumberPayload

		if err := json.Unmarshal(line, &payload); err != nil {
			fmt.Println("Error al obtener json", err.Error())
			ctx.JSON(http.StatusInternalServerError, SimpleResponse{"Error": "Error obteniendo los datos"})
			return
		}

		println("numero a guardar: ", payload.Numero)

		bd = append(bd, payload.Numero)

	}

}
