package main


import (
	"encoding/json"
	"fmt"
	"os"
)
type response struct{
	User	string 				`json:"username"`
	Policy	map[string]string	`json:"policy"`
}
func main(){
	t := []string{
		"blue",
		"jack",
	}
	json.Marshal(t)
	rep := &response{
		"xwt",
		map[string]string{
			"jack":"friends",
			"tony":"fff",
		},
	}
	
	bolD, _ := json.Marshal(rep)
	fmt.Println(string(bolD))
	m := make(map[string]interface{})
	inner := make(map[string]interface{})
	inner["lalal"] = "123"
	m["username"] = "jack"
	m["age"] = 10
	m["m"] = inner
	test, _ := json.Marshal(m)
	fmt.Println(string(test))
	
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}