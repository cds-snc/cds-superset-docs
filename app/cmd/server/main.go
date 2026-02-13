package main

import (
	"fmt"
	"log"
	"net/http"

	"wordpress-go-proxy/internal/api"
	"wordpress-go-proxy/internal/config"
	"wordpress-go-proxy/internal/handlers"
	"wordpress-go-proxy/internal/middleware"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	_ "golang.org/x/crypto/x509roots/fallback"
)

func main() {

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Create WordPress client.  This will fetch menus asynchronously.
	wordPressClient := api.NewWordPressClient(
		cfg.WordPressBaseURL,
		cfg.WordPressUsername,
		cfg.WordPressPassword,
		cfg.WordPressMenuIdEn,
		cfg.WordPressMenuIdFr)

	siteNames := map[string]string{
		"en": cfg.SiteNameEn,
		"fr": cfg.SiteNameFr,
	}

	// Set up routes
	http.Handle("/static/", http.StripPrefix("/static/", handlers.NewStaticHandler("static")))
	http.Handle("/robots.txt", middleware.SecurityHeaders(handlers.NewRobotsHandler()))
	http.Handle("/.well-known/security.txt", middleware.SecurityHeaders(handlers.NewSecurityHandler(cfg.SecurityTxtContent)))
	http.Handle("/", middleware.SecurityHeaders(handlers.NewPageHandler(cfg.GoogleAnalyticsID, siteNames, wordPressClient)))

	// Determine if this is a Lambda or HTTP server startup
	if cfg.Port == "" {
		log.Println("Starting Lambda Handler")
		lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
	} else {
		fmt.Printf("Server starting on port %s...\n", cfg.Port)
		log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
	}
}
