package main

import (
	"lab-test.com/module/cmd/cfe"
	"lab-test.com/module/cmd/lavanderia"
	"lab-test.com/module/cmd/sapam"
)

func main() {

	go sapam.Run()

	go lavanderia.Run()

	go cfe.Run()

	//time.Sleep(10 * time.Second)
	go lavanderia.Do_after_init()

	select {}
}
