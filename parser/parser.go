package main

import "encoding/json"

type Properties struct {
	EvcID string `json:"EVCId"`
	Name  string `json:"Name"`
}

func main() {
	input := []byte(`
		{
			"EVCId": "123445",
			"Extra": "xxx"
		}
	`)
	var test Properties
	err := json.Unmarshal(input, &test)
	if err != nil {
		panic(err)
	}
	//data, err := json.Marshal(test)
	println(test.EvcID)
	println(test.Name)
}
