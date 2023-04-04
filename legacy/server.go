package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestContext struct {
	RequestHeaders map[string]string `json:"requestHeaders"`
	RequestTrailers map[string]string `json:"requestTrailers"`
	RequestBody string `json:"requestBody"`
	InvocationContext map[string]string `json:"invocationContext"`
}

type IRequestContextResponse struct {
	HeadersToAdd map[string]string `json:"headersToAdd"`
	HeadersToReplace map[string]string `json:"headersToReplace"`
	Body string `json:"body"`
	Properties map[string]string `json:"properties"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Printf("Headers: %v \n", r.Header)
		fmt.Printf("Method: %v \n", r.Method)
		fmt.Printf("URL: %v \n", r.URL)
		fmt.Printf("Body: %v \n", r.Body)
	
		if r.Method != "POST" {
			http.Error(w, "Only POST method is supported", http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var jsonObj interface{}
		err = json.Unmarshal(body, &jsonObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonReq, err := json.Marshal(jsonObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestContext RequestContext
		err1 := json.Unmarshal(body, &requestContext)
		if err1 != nil {
			fmt.Printf("ERROR unmarshalling requestContext: %v \n", err1)
		}

		fmt.Printf("XXXX requestContext: %v \n", requestContext)

		var requestContextResponse IRequestContextResponse
		if requestContext.RequestHeaders[":path"] != "/http-bin-api-basic/1.0.8/get" {
			fmt.Println("/http-bin-api-basic/1.0.8/get")
			requestContextResponse.Properties = map[string]string{"custom-key-1": "custom-value-1"}
			jsonReq, err := json.Marshal(requestContextResponse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonReq)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonReq)
	})

	fmt.Println("XXXX Listening on port :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

