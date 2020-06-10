package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const port = ":3001"

type request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type response struct {
	Result  float64 `json:"result"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
}

func main() {
	http.HandleFunc("/calculator.sum", calculatorHandler)
	http.HandleFunc("/calculator.mul", calculatorHandler)
	http.HandleFunc("/calculator.div", calculatorHandler)
	http.HandleFunc("/calculator.sub", calculatorHandler)

	log.Println("Starting calculator server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func jsonResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("[summationHandler] decode json error: ", err)
		jsonResponse(w, http.StatusOK, response{Code: "99", Message: "decode json error"})
		return
	}
	defer r.Body.Close()

	res, err := calculate(req.A, req.B, r.RequestURI)
	if err != nil {
		jsonResponse(w, http.StatusOK, response{Code: "99", Message: err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, response{Code: "00", Message: "success", Result: res})
	return
}

func calculate(a, b float64, opt string) (float64, error) {

	if opt == "/calculator.sum" {
		return a + b, nil
	}

	if opt == "/calculator.sub" {
		return a - b, nil
	}

	if opt == "/calculator.div" {
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}

		return a / b, nil
	}

	if opt == "/calculator.mul" {
		return a * b, nil
	}

	return 0, errors.New("invalid operation")
}
