package _11

import "time"

// Anti-pattern: Using a Config Struct (Page 41)
// Problem: Config structs make it unclear which fields are required vs optional

// BadDBConfig demonstrates the anti-pattern
type BadDBConfig struct {
	Host      string
	Port      int
	MaxConns  int
	Timeout   time.Duration
	EnableSSL bool
	ReadOnly  bool
}

type BadDB struct {
	config BadDBConfig
}

// Problems:
// 1. Which fields are required? Host? Port?
// 2. Is 0 timeout intentional or forgotten?
// 3. Have to pass entire struct even for one option
func NewBadDB(config BadDBConfig) *BadDB {
	// Must check and set defaults everywhere
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}
	if config.MaxConns == 0 {
		config.MaxConns = 10
	}
	return &BadDB{config: config}
}

// Usage is verbose and unclear:
// db := NewBadDB(BadDBConfig{
//     Host: "localhost",
//     Port: 5432,
//     // Forgot MaxConns - is that ok?
// })

// GoodDB demonstrates functional options pattern
type GoodDB struct {
	host      string
	port      int
	maxConns  int
	timeout   time.Duration
	enableSSL bool
	readOnly  bool
}

type DBOption func(*GoodDB)

func WithHost(host string) DBOption {
	return func(db *GoodDB) { db.host = host }
}

func WithPort(port int) DBOption {
	return func(db *GoodDB) { db.port = port }
}

func WithMaxConns(n int) DBOption {
	return func(db *GoodDB) { db.maxConns = n }
}

func WithTimeout(d time.Duration) DBOption {
	return func(db *GoodDB) { db.timeout = d }
}

func WithSSL() DBOption {
	return func(db *GoodDB) { db.enableSSL = true }
}

func WithReadOnly() DBOption {
	return func(db *GoodDB) { db.readOnly = true }
}

// NewGoodDB has sensible defaults and clear optional configuration
func NewGoodDB(options ...DBOption) *GoodDB {
	// This initializes defaults
	db := &GoodDB{
		host:     "localhost",
		port:     5432,
		maxConns: 10,
		timeout:  30 * time.Second,
	}

	// This will result the defaults if a user chooses to do so
	for _, opt := range options {
		opt(db)
	}
	return db
}

func do() {
	// Usage is clear and flexible:
	db := NewGoodDB(
		WithHost("prod.db"),
		WithSSL(),
		WithTimeout(5))

	//or with all defaults:
	dbDefault := NewGoodDB()
	_, _ = db, dbDefault
}
