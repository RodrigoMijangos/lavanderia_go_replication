package lavanderia

func definir_msg(status int) Payload {

	switch status {
	case 0:
		return Payload{"status": "El servicio ha terminado de lavar"}
	case 1:
		return Payload{"status": "El servicio ha terminado de Centrifugar"}
	case 2:
		return Payload{"status": "El servico ha terminado de enjuagar"}
	case 3:
		return Payload{"status": "El servicio ha terminado de secar"}
	default:
		return Payload{"message": "El servicio ha terminado con éxito"}
	}
}
