package lavanderia

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

var lavadoras = make(chan struct{}, 3)

func attend(quantity int, carga int, ctx *gin.Context, status int, old_status int) {

	lavadoras <- struct{}{}

	if status == 4 {

		jsonData, _ := json.Marshal(definir_msg(status))

		ctx.Writer.Write(jsonData)
		ctx.Writer.Flush()

		defer func() {
			<-lavadoras
		}()

		return
	}

	if status == 0 || status == 2 {
		for {
			if is_water_available(carga) {
				tanque_agua -= carga * 2
				break
			}
		}
	}

	do_petition_for_energy(
		strconv.Itoa(quantity),
		ctx,
	)

	jsonData, _ := json.Marshal(definir_msg(status))

	ctx.Writer.Write(jsonData)
	ctx.Writer.Flush()
	new_code, old_code := definir_ciclo(status, old_status)
	<-lavadoras
	attend(quantity, carga, ctx, new_code, old_code)
}

func is_water_available(carga int) bool {
	return tanque_agua >= (carga * 2)
}

func definir_ciclo(status int, old_status int) (int, int) {
	switch status {
	case 0:
		return 1, 0
	case 1:
		if old_status == 0 {
			return 2, 0
		}
		if old_status == 2 {
			return 3, 2
		}
	case 2:
		return 1, 2
	case 3:
		return 4, 3
	}
	return -1, -1
}
