package mylib2

import (
	"log"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func Demo() {
	mux := gwruntime.NewServeMux()
	log.Printf("Mux is %v", mux)
}
