package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"anon-wa/internal/config"
	"anon-wa/internal/handler"
	"anon-wa/internal/sender"
	"anon-wa/internal/service"
)

func main() {
	// Load .env for local only (safe on Railway)
	_ = godotenv.Load()

	cfg := config.Load()

	// Init sender
	var msgSender sender.MessageSender
	switch cfg.Provider {
	case "meta":
		msgSender = sender.NewMeta(cfg.MetaToken, cfg.MetaPhoneID)
	default:
		msgSender = sender.NewFonte(cfg.FonteAPIKey)
	}

	msgService := service.New(msgSender, cfg.Watermark)

	mux := http.NewServeMux()

	// API endpoint
	mux.HandleFunc("/send", handler.SendHandler(msgService))

	// Serve frontend
	frontendPath := getFrontendPath()
	log.Println("Serving frontend from:", frontendPath)
	mux.Handle("/", http.FileServer(http.Dir(frontendPath)))

	// ✅ PORT handling (Railway + local)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port // fallback for local
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// Helper function to find frontend folder
func getFrontendPath() string {
	paths := []string{
		"./frontend",
		"../../frontend",
		"../../../frontend",
	}

	for _, p := range paths {
		absPath, _ := filepath.Abs(p)
		if _, err := os.Stat(absPath); err == nil {
			return absPath
		}
	}

	wd, _ := os.Getwd()

	if _, err := os.Stat(filepath.Join(wd, "frontend")); err == nil {
		return filepath.Join(wd, "frontend")
	}

	if filepath.Base(wd) == "frontend" {
		return wd
	}

	log.Println("Warning: frontend folder not found, using ../../frontend")
	return "../../frontend"
}
