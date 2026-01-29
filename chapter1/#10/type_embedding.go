package _10

import (
	"errors"
	"fmt"
)

// Anti-pattern: Problems with Type Embedding (Page 36)
// Problem: Type embedding can lead to unexpected behavior and confusion
// between "is-a" relationships and "has-a" relationships

// BadLogger demonstrates problematic type embedding
type Logger struct {
	prefix string
}

func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.prefix, message)
}

func (l *Logger) DoWeNeedThis(message string) {}

// BadService embeds Logger, which promotes all Logger methods
// This creates confusion - is BadService a Logger or does it have a Logger?
type BadService struct {
	Logger // Embedded type - promotes methods
	name   string
}

// This can lead to confusion and unintended method promotion
func BadEmbeddingExample() {
	service := BadService{
		Logger: Logger{prefix: "SERVICE"},
		name:   "UserService",
	}

	service.Log("Starting service")
	service.DoWeNeedThis("NoWeDont")

	// What if we want to change the logger implementation?
	// We're stuck with the embedded type
}

// GoodService demonstrates proper composition using an explicit field
type GoodService struct {
	logger *Logger // Explicit field - clear "has-a" relationship
	name   string
}

func NewGoodService(name string, logger *Logger) *GoodService {
	return &GoodService{
		logger: logger,
		name:   name,
	}
}

// Explicit method that delegates to the logger
// This makes it clear what's happening and gives us more control
func (s *GoodService) Log(message string) {
	s.logger.Log(fmt.Sprintf("[%s] %s", s.name, message))
}

func GoodEmbeddingExample() {
	logger := &Logger{prefix: "APP"}
	service := NewGoodService("UserService", logger)

	// Clear delegation - we know the service is using a logger
	service.Log("Starting service")

	// Easy to swap out the logger implementation
	// or add logging middleware
}

// Another problem: Embedding interfaces can lead to surprising behavior
type BadWriter struct {
	// This embeds the Write method, but if not implemented,
	// will panic at runtime
	Writer
}

type Writer interface {
	Write(data []byte) error
}

// Better: Explicit composition
type GoodWriter struct {
	writer Writer // Explicit field
}

func (w *GoodWriter) Write(data []byte) error {
	if w.writer == nil {
		return errors.New("writer not initialized")
	}
	return w.writer.Write(data)
}
