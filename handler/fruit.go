package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"shop-integral/model"
)

const FruitDocType = "fruit"

// 新增果实
func (c *FruitCC) UpdateFruit(stub contractapi.TransactionContextInterface, args string) (*TxIdResult, error) {
	var req = model.Fruit{}
	err := json.Unmarshal([]byte(args), &req)
	if err != nil {
		return nil, err
	}
	req.DocType = FruitDocType
	if req.Uuid == "" {
		return nil, errors.New("必传参数为空")
	}
	//查找以前的数据
	data, err := Read(stub, FruitDocType, []string{req.Uuid})
	if err != nil {
		return nil, fmt.Errorf("Read  err: " + err.Error())
	}
	//读取之前的
	readFruit := model.Fruit{}
	if string(data) != "" {
		err := json.Unmarshal(data, &readFruit)
		if err != nil {
			return nil, fmt.Errorf("Unmarshal  err: " + err.Error())
		}
	}
	//产地证明
	if req.CertOfOrigin != "" {
		readFruit.CertOfOrigin = req.CertOfOrigin
	}
	//仓储
	if req.Storehouse != "" {
		readFruit.Storehouse = req.Storehouse
	}
	//光照
	if len(req.PlantDevelopment.BeamList) != 0 {
		readFruit.PlantDevelopment.BeamList = append(readFruit.PlantDevelopment.BeamList, req.PlantDevelopment.BeamList...)
	}
	//湿度
	if len(req.PlantDevelopment.HumidityList) != 0 {
		readFruit.PlantDevelopment.HumidityList = append(readFruit.PlantDevelopment.HumidityList, req.PlantDevelopment.HumidityList...)
	}
	//温度
	if len(req.PlantDevelopment.TemperatureList) != 0 {
		readFruit.PlantDevelopment.TemperatureList = append(readFruit.PlantDevelopment.TemperatureList, req.PlantDevelopment.TemperatureList...)
	}
	//大棚图片
	if len(req.PlantDevelopment.Images) != 0 {
		readFruit.PlantDevelopment.Images = req.PlantDevelopment.Images
	}
	//存储交易ID
	readFruit.TxId = stub.GetStub().GetTxID()

	// 保存 req
	_, err = Save(stub, FruitDocType, []string{req.Uuid}, readFruit)
	if err != nil {
		return nil, fmt.Errorf("Save readFruit err: " + err.Error())
	}
	return &TxIdResult{TxId: req.TxId}, nil
}

// 查询果实信息
func (*FruitCC) QueryFruit(stub contractapi.TransactionContextInterface, args string) (*model.QueryFruitResp, error) {
	var req model.QueryFruitReq
	var result model.QueryFruitResp
	err := json.Unmarshal([]byte(args), &req)
	if err != nil {
		return nil, fmt.Errorf("query fruit failed: Not valid QueryFruitlistReq structure: " + err.Error())
	}
	fmt.Printf("req in QueryFruit is : %v \n", req)
	if req.Uuid == "" {
		return nil, errors.New("必填信息为空")
	}

	data, err := Read(stub, FruitDocType, []string{req.Uuid})
	if err != nil {
		return nil, fmt.Errorf("query fruit err: " + err.Error())
	}
	//如果没查到，直接返回空的
	if string(data) == "" {
		return &result, nil
	}
	t := model.Fruit{}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("query fruit err: " + err.Error())
	}

	result.Fruit = t
	result.TxID = t.TxId
	fmt.Printf("QueryFruit Result is %v \n", result)
	return &result, nil
}
