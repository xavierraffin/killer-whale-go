package main

import (
	"encoding/json"
	"log"
)

func printObj(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
