package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	"shop-integral/handler"
)

func main() {
	cc, err := contractapi.NewChaincode(&handler.FruitCC{})
	if err != nil {
		log.Panicf("error create cc chaincode: %s", err)
	}

	if err := cc.Start(); err != nil { //
		log.Panicf("error starting bidcc chaincode: %s", err)
	}
}
