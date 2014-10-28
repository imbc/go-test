package main

import (
	"encoding/json"
	"fmt"
)

// see http://stackoverflow.com/questions/24377907/golang-issue-with-accessing-nested-json-array-after-unmarshalling

type Base struct {
	Query   string `json:"query"`
	Count   int    `json:"count"`
	Objects []struct {
		ItemID      string `json:"ITEM_ID"`
		ProdClassID string `json:"PROD_CLASS_ID"`
		Available   int    `json:"AVAILABLE"`
	}
}

func main() {
	payload := []byte(`{"query": "QEACOR139GID","count": 1,"objects": [{"ITEM_ID": "QEACOR139GID","PROD_CLASS_ID": "BMXCPGRIPS","AVAILABLE": 19}]}`)
	var result map[string]interface{}
	if err := json.Unmarshal(payload, &result); err != nil {
		panic(err)
	}

	fmt.Println(result["objects"].([]interface{})[0].(map[string]interface{})["ITEM_ID"].(string))
	b := &Base{}
	e := json.Unmarshal(payload, &b)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("%+v\n", b)
	fmt.Println(b.Objects[0].ItemID)
}
