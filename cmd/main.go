package main

import (
	server "diploma/internal/app/http"
	"diploma/internal/app/service"
	"log"
)

func main() {
	service := service.New("sms.data", "http://localhost:8383/mms", "voice.data", "email.data", "billing.data", "http://localhost:8383/support", "http://localhost:8383/accendent")
	server := server.NewServer(service, 20)
	log.Println("server started")
	server.Serve()
	log.Println("goodbye")
}
