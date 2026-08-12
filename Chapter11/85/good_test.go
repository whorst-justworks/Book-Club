package main

import "testing"

// GOOD: Table-driven tests. Each case is a single struct entry.
// Easy to add new cases, descriptive names in output, zero duplication.
//
// Run with: go test -v ./...

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{
			name:    "empty name",
			user:    User{Name: "", Email: "alice@example.com", Age: 25},
			wantErr: ErrEmptyName,
		},
		{
			name:    "whitespace-only name",
			user:    User{Name: "   ", Email: "alice@example.com", Age: 25},
			wantErr: ErrEmptyName,
		},
		{
			name:    "negative age",
			user:    User{Name: "Alice", Email: "alice@example.com", Age: -1},
			wantErr: ErrInvalidAge,
		},
		{
			name:    "age exceeds maximum",
			user:    User{Name: "Alice", Email: "alice@example.com", Age: 200},
			wantErr: ErrInvalidAge,
		},
		{
			name:    "valid user",
			user:    User{Name: "Alice", Email: "alice@example.com", Age: 25},
			wantErr: nil,
		},
		{
			name:    "edge case age 0",
			user:    User{Name: "Baby", Email: "baby@example.com", Age: 0},
			wantErr: nil,
		},
		{
			name:    "edge case age 150",
			user:    User{Name: "Elder", Email: "elder@example.com", Age: 150},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUser(tt.user)
			if err != tt.wantErr {
				t.Errorf("ValidateUser() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
