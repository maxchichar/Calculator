package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/Knetic/govaluate"
)

type Request struct {
	Expr string `json:"expr"`
}
type Response struct {
	Result interface{} `json:"result"`
}

// the brain of the calculator
func calculate(w http.ResponseWriter, r *http.Request) {
	var req Request // were we store the data the user sent us

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Result: "mperi"})
		return
	}

	expression, err := govaluate.NewEvaluableExpression(req.Expr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Result: "mperi"})
		return
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Result: "mperi"})
		return
	}

	// Handles dividing by zero
	switch result.(type) {
	case float64:
		value := result.(float64)
		if math.IsInf(value, 0) || math.IsNaN(value) {
			json.NewEncoder(w).Encode(Response{Result: "mperi"})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{Result: result})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/calc", calculate)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
