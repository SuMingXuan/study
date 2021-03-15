package main

import (
	"log"
	"net/http"
	"encoding/json"
	"sync"
)

func main() {
	getProductsHandler := http.HandlerFunc(getProducts)	
	http.Handle("/", getProductsHandler)
	e := http.ListenAndServeTLS(":443", "../server.crt", "../server.key", nil)
	if e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}
var inputMap map[string]bool = make(map[string]bool)
var wg sync.WaitGroup
func getProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	outputSlice := make([]bool, 0)

	for _, values := range query {
		// 获取到url的参数，并遍历
		for _, value := range values {
			wg.Add(1)
			go func (value string)  {
				defer wg.Done()
				exist := inputMap[value]
				if exist {
					outputSlice = append(outputSlice, false)
				} else {
					inputMap[value] = true
					outputSlice = append(outputSlice, true)
				}
				inputMap[value] = true
			}(value)
			wg.Wait()
		}
	}
	b, err := json.Marshal(reverse(outputSlice))
	if err != nil {
		log.Fatal("json parse error: ", err)
	}
	w.Write(b)
}

func reverse(s []bool) []bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}