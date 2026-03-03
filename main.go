package main

import(
	"encoding/json"
	"log"
	"net/http"
)

type CalcRequest struct{
	A float64 `json: "a"`
	B float64 `json: "b"`
	Op string `json: "op"`
}

type CalcResponse struct{
	Result float64 `json:"result"`
	Error string `json:"error,omitemty"`
}

// the brain of the calculator
func calculate(w http.ResponseWriter, r *http.Request) {
	var req CalcRequest // were we store the data the user sent us

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
	}

	var res CalcResponse // were we put the answers

	switch req.Op {
	case "+":
		res.Result = req.A + req.B
	case "-":
		res.Result = req.A - req.B
	case "*":
		res.Result = req.A * req.B
	case "/":
		if req.B == 0 {
			res.Error = "division by zero"
		}else {
			res.Result = req.A / req.B
		}
	default:
		res.Error = "unknown operator"
	}

	w.Header().Set("Content-Type", "application/json") //here's my response in json format
	json.NewEncoder(w).Encode(res)

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/calc", calculate)
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe("8080:", nil)
}