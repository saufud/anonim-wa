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
	godotenv.Load()
	cfg := config.Load()

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

	// ✅ Detect frontend path smartly
	frontendPath := getFrontendPath()
	log.Println("Serving frontend from:", frontendPath)
	mux.Handle("/", http.FileServer(http.Dir(frontendPath)))

	log.Println("Server running on http://localhost:" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}

// Helper function untuk cari frontend folder
func getFrontendPath() string {
	// Try relative path dari binary location
	paths := []string{
		"./frontend",           // Kalau run dari root project
		"../../frontend",       // Kalau run dari cmd/server
		"../../../frontend",    // Kalau build & run dari cmd/server
	}

	for _, p := range paths {
		absPath, _ := filepath.Abs(p)
		if _, err := os.Stat(absPath); err == nil {
			return absPath
		}
	}

	// Fallback: cari dari working directory
	wd, _ := os.Getwd()
	
	// Cek apakah ada frontend/ di working dir
	if _, err := os.Stat(filepath.Join(wd, "frontend")); err == nil {
		return filepath.Join(wd, "frontend")
	}

	// Cek apakah working dir sudah di dalam frontend
	if filepath.Base(wd) == "frontend" {
		return wd
	}

	// Default fallback
	log.Println("Warning: frontend folder not found, using ../../frontend")
	return "../../frontend"
}