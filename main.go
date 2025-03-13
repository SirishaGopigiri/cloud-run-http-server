package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/SirishaGopigiri/sample-grpc-server/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is a simple HTTP server running on Cloud Run.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/getUsers", getUsers)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.NewClient("golang-grpc-server-184008090433.us-central1.run.app:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUsersClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := client.GetUsers(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Printf("error %v", err)
		http.Error(w, "Invalid response from upstream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Users)
}
