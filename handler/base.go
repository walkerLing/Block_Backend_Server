package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

//商城积分
type FruitCC struct {
	contractapi.Contract
}

type TxIdResult struct {
	TxId string `json:"txId"`
}

//保存
func Save(stub contractapi.TransactionContextInterface, objectType string, attributes []string, payload interface{}) ([]byte, error) {
	now := time.Now()
	fmt.Println("Saving", objectType)
	fmt.Printf("%#v\n", payload)

	key, err := stub.GetStub().CreateCompositeKey(objectType, attributes)
	if err != nil {
		return nil, err
	}

	bytes, ok := payload.([]byte)
	if !ok {
		bytes, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	err = stub.GetStub().PutState(key, bytes)
	if err != nil {
		return nil, err
	}
	since := time.Since(now)
	fmt.Println("Save时长：", since)

	return bytes, nil
}

func Read(stub contractapi.TransactionContextInterface, objectType string, attributes []string) ([]byte, error) {
	now := time.Now()
	key, err := stub.GetStub().CreateCompositeKey(objectType, attributes)
	if err != nil {
		return nil, err
	}

	res, err := stub.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	since := time.Since(now)
	fmt.Println("Read时长：", since)
	return res, nil
}
