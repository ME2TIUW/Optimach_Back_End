package database

import (
	"database/sql"
	"fmt"
	"os"
	"time" // Added for Connection Pooling settings

	_ "github.com/lib/pq"
)

// 1. GLOBAL VARIABLE (Exported so you can use 'database.DB' anywhere)
var DB *sql.DB

// 2. Renamed to 'Init' because we only run this ONCE.
func Init() {
	var err error
	var connectionString string

	// --- ENVIRONMENT SWITCHING LOGIC ---
	// If DATABASE_URL exists, we are in Production (Neon).
	// If not, we fall back to Local variables.
	prodURL := os.Getenv("DATABASE_URL")

	if prodURL != "" {
		// PRODUCTION (NEON)
		connectionString = prodURL
	} else {
		// LOCAL
		connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable application_name='%s'",
			os.Getenv("DB_HOST_POSTGRES"),
			os.Getenv("DB_PORT_POSTGRES"),
			os.Getenv("DB_USER_POSTGRES"),
			os.Getenv("DB_PW_POSTGRES"),
			os.Getenv("DB_NAME_POSTGRES"),
			os.Getenv("DB_APP_NAME"),
		)
	}

	// 3. OPEN THE CONNECTION
	// Note: We assign to the global 'DB' variable, not a local one.
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Errorf("error opening database connection: %w", err))
	}

	// 4. PING (Fail fast if credentials are wrong)
	err = DB.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging database: %w", err))
	}

	// 5. CRITICAL PERFORMANCE SETTINGS (The Fix for your Lag)
	// This keeps connections "alive" in the background so you don't handshake every time.
	DB.SetMaxOpenConns(25)                  // Allow up to 25 simultaneous connections
	DB.SetMaxIdleConns(25)                  // Keep 25 connections "hot" and ready to use
	DB.SetConnMaxLifetime(30 * time.Minute) // Recycle connections every 30 mins to prevent staleness

	fmt.Println("🚀 Database connection pool initialized successfully.")
}
