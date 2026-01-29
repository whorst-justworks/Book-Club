package _interface

import "testing"

// MockDB implements DB interface for testing
type MockDB struct {
	queryResult []string
	queryError  error
}

func (m *MockDB) Query(sql string) ([]string, error) {
	return m.queryResult, m.queryError
}

// Test shows how getter returning interface enables mocking
func TestServiceWithMock(t *testing.T) {
	// Create a mock database
	mockDB := &MockDB{
		queryResult: []string{"user1", "user2"},
		queryError:  nil,
	}

	// Create service with real PostgresDB
	service := &Service{db: &PostgresDB{connString: "test"}}

	// In tests, we can work with the DB() interface
	// and verify behavior without hitting a real database
	db := service.DB()
	if db == nil {
		t.Fatal("DB() should return non-nil interface")
	}

	// We can also test with the mock directly
	results, err := mockDB.Query("SELECT * FROM users")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}
