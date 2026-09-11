package main

type Publisher interface {
	Publish(string)
}

type Handler struct {
	publisher Publisher
}

func (h Handler) Handle(book string) {
	go h.publisher.Publish(book)
}
