package main

import (
	server "diploma/internal/app/http"
	"diploma/internal/app/service"
	"diploma/internal/config"
	"log"
)

func main() {
	cfg := config.Read()

	service := service.New(cfg.Service.PathSms,
		cfg.Service.MmsAddress,
		cfg.Service.PathVoiceCall,
		cfg.Service.PathEmail,
		cfg.Service.PathBilling,
		cfg.Service.SupportAddress,
		cfg.Service.IncidentAddress,
	)
	server := server.NewServer(service, 20)
	log.Println("server started")
	server.Serve()
	log.Println("goodbye")
}
