package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/ipqs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ipqs_apikey := os.Getenv("IPQS_APIKEY")
	fmt.Println(ipqs_apikey)

	conf := config.Config{
		// you can set auth values to config directly, otherwise used from environment variables.
		APIKey: ipqs_apikey,
		Debug:  false,
	}

	svc, err := ipqs.New(conf)
	if err != nil {
		panic(err)
	}

	// execute API
	resp, err := svc.IPReputation("8.8.8.8")
	if err != nil {
		panic(err)
	}
	if resp.HasError() {
		panic(fmt.Errorf("code=[%d] message=[%s]", resp.ErrData.StatusCode, resp.ErrData.Message))
	}

	// just print response in json format
	b, _ := json.Marshal(resp)
	fmt.Printf("%s", string(b))
}
