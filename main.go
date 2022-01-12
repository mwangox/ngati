package main

import (
	"log"
	"nitra/registry/cmd"
)

func main()  {
	if err := cmd.Start(); err != nil{
		log.Fatal("Failed to start ngati registry:", err)
	}
}
