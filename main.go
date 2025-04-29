package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

const (
	codeLength = 6
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)


func main() {
	var err error
	db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", shortenHandler)
	mux.HandleFunc("/", redirectHandler)

	// Wrap your mux with security headers
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

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	var code string
	for {
		code = generateCode()
		var exists bool
		err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM short_links WHERE code = $1)`, code).Scan(&exists)
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		if !exists {
			break
		}
	}

	_, err := db.Exec(`INSERT INTO short_links (code, url) VALUES ($1, $2)`, code, url)
	if err != nil {
		http.Error(w, "Failed to save", http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("https://yourdomain.com/%s", code)
	fmt.Fprintf(w, "Shortened URL: %s\n", shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Path[1:]
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	var url string
	err := db.QueryRow(`SELECT url FROM short_links WHERE code = $1`).Scan(&url)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Security best practices
		w.Header().Set("Content-Security-Policy", "default-src 'none';")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		next.ServeHTTP(w, r)
	})
}