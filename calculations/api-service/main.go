package main

import (
	pb "calculations/pb/fib/v1"
	pi "calculations/pb/pi/v1"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	fibClient := pb.NewFibServiceClient(conn)
	piClient := pi.NewPiServiceClient(conn)

	type FibJSON struct {
		K int64 `json:"k"`
	}

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(r.Body)
		var fibJSON FibJSON
		err := json.NewDecoder(r.Body).Decode(&fibJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(fibJSON)
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &pb.FibRequest{FibNum: uint64(fibJSON.K)}
		if resp, err := fibClient.Compute(ctx, req); err == nil {
			msg := fmt.Sprintf("Fibonacci number is %d", resp.Result)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("POST")
	routes.HandleFunc("/pi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()
		req := &pi.PiRequest{}
		if resp, err := piClient.Compute(ctx, req); err == nil {
			fmt.Println(resp.Result)
			msg := fmt.Sprintf("Pi number is %f", resp.Result)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}
	}).Methods("GET")

	fmt.Println("Application is running on :8080")
	http.ListenAndServe(":8080", routes)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}
