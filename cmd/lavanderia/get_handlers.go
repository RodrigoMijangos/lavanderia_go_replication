package lavanderia

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SimpleResponse map[string]string

func handleService(ctx *gin.Context) {
	quantityParam := ctx.Param("quantity")

	println(quantityParam)

	quantity, err := strconv.Atoi(quantityParam)

	if err != nil || quantity <= 0 {
		ctx.JSON(
			http.StatusBadRequest,
			SimpleResponse{"error": "El parametro debe ser un número entero positivo"},
		)
	}

	jsonData, _ := json.Marshal(SimpleResponse{"message": "Solicitud recibida"})
	ctx.Writer.Write(jsonData)
	ctx.Writer.Flush()

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Transfer-Encoding", "chunked")

	if quantity <= 9 && quantity > 6 {
		attend(30, quantity, ctx, 0, -1)
		return
	}

	if quantity <= 6 && quantity > 3 {
		attend(20, quantity, ctx, 0, -1)
		return
	}

	if quantity <= 3 {
		attend(10, quantity, ctx, 0, -1)
		return
	}

	ctx.JSON(http.StatusBadRequest, SimpleResponse{"error": "El parametro no puede ser mayor a 9"})

}
