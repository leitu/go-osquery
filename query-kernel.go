package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {
	var obj []interface{}

	//exec osqueryi to get the json data
	bs, err := exec.Command("osqueryi", "SELECT version  FROM kernel_info", "--json").Output()
	if err != nil {
		panic(err)
	}

	//Unmarshal to obj
	err = json.Unmarshal(bs, &obj)

	//get the data with interface again
	//from http://stackoverflow.com/questions/25214036/getting-invalid-operation-mymaptitle-type-interface-does-not-support-in
	//http://blog.golang.org/json-and-go
	for i := range obj {
		m := obj[i].(map[string]interface{})
		fmt.Println(m["version"])
	}

}
