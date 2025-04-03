package main

import (
	"alpha/routers"
	"fmt"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8083")
	err := r.Run(fmt.Sprintf(""))
	if err != nil {
		fmt.Printf("Run server failed, err: %s\n", err)
		return
	}
}
