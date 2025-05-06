package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hlmelton/dimbus/internal/database"
	"github.com/hlmelton/dimbus/templates/components"
	"github.com/hlmelton/dimbus/templates/pages"
)

const (
	codeLength = 6
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// Serve the home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handleRedirect(w, r, db)
			return
		}
		
		err := pages.Home().Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
			return
		}
	})

	// Handle URL shortening
	mux.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing URL", http.StatusBadRequest)
			return
		}

		code := generateCode()
		for {
			exists, err := db.CodeExists(code)
			if err != nil {
				http.Error(w, "DB error", http.StatusInternalServerError)
				return
			}
			if !exists {
				break
			}
			code = generateCode()
		}

		if err := db.CreateLink(code, url); err != nil {
			http.Error(w, "Failed to save", http.StatusInternalServerError)
			return
		}

		shortURL := fmt.Sprintf("https://dimbus.link/%s", code)
		err = components.ShortenedLink(shortURL).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Failed to render component", http.StatusInternalServerError)
			return
		}
	})

	// Wrap with security headers
	handler := securityHeaders(mux)

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func handleRedirect(w http.ResponseWriter, r *http.Request, db *database.DB) {
	code := r.URL.Path[1:]
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	url, err := db.GetLink(code)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if url == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' https://unpkg.com https://cdn.tailwindcss.com;")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		next.ServeHTTP(w, r)
	})
} 