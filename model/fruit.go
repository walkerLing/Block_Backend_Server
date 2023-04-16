package model

//果实结构体
type Fruit struct {
	// 对象类型
	DocType string `json:"docType"`
	// 溯源码
	Uuid string `json:"uuid"`
	// 交易哈希
	TxId string `json:"txId"`
	//原产地证明 图片地址
	CertOfOrigin string `json:"certOfOrigin"`
	//仓储
	Storehouse string `json:"storehouse"`
	//生产种植
	PlantDevelopment PlantDevelopment `json:"plantDevelopment"`
}

//生产种植
type PlantDevelopment struct {
	S string `json:"s"`
	//温度
	TemperatureList []EnvironmentalInfo `json:"temperatureList,omitempty" metadata:",optional"`
	//湿度
	HumidityList []EnvironmentalInfo `json:"humidityList,omitempty" metadata:",optional"`
	//光照
	BeamList []EnvironmentalInfo `json:"beamList,omitempty" metadata:",optional"`
	//图片
	Images []string `json:"images,omitempty" metadata:",optional"`
}

//环境信息
type EnvironmentalInfo struct {
	//数据
	Data string `json:"data"`
	//
	Time string `json:"time"`
}

//查询积分请求
type QueryFruitReq struct {
	//溯源码ID
	Uuid string `json:"uuid"`
	// 当前页码
	PAGE_INDEX int `json:"PAGE_INDEX"`
	// 每页数量
	PAGE_SIZE int `json:"PAGE_SIZE"`
}

//返回数据
type QueryFruitResp struct {
	TOTAL int `json:"TOTAL"`
	SIZE  int `json:"SIZE"`
	//PAGE_INDEX int    `json:"PAGE_INDEX"`
	//PAGE_SIZE  int    `json:"PAGE_SIZE"`
	TxID string `json:"txId"`
	//果实信息
	Fruit Fruit `json:"fruit,omitempty"`
}
