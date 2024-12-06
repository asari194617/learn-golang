package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"learn-golang/utils"
)

type CalcResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	if expression == "" {
		http.Error(w, "Missing 'expression' parameter", http.StatusBadRequest)
		return
	}

	fmt.Println("expression: ", expression)

	result, err := utils.Calculate(expression)
	response := CalcResponse{}

	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Result = result
		fmt.Println("result: ", result)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/calc", calcHandler)

	port := ":8080"
	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
