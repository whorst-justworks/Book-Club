package _interface

// Good Use Case: Getter returning interface enables mocking in tests

// DB interface can be mocked
type DB interface {
	Query(sql string) ([]string, error)
}

// PostgresDB is the concrete implementation
type PostgresDB struct{ connString string }

// Service stores concrete type internally
type Service struct {
	db *PostgresDB
}

// DB getter returns interface, not concrete type
// This allows tests to mock the database
func (s *Service) DB() DB {
	return s.db
}
