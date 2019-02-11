package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pair struct {
	First  string `json:"first"`
	Second string `json:"sec"`
}

func main() {
	data := make(map[string]interface{})
	data["response"] = Pair{First: "Три", Second: "Three"}
	data["error"] = nil
	enc := json.NewEncoder(os.Stdout)
	fmt.Println(data)
	enc.Encode(data)
}
