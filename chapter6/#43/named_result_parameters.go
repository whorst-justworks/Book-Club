
// Unnamed result parameters
type Locator interface {
	getCoordinates(address string) (float64, float64, error)
}

// Named result parameters
// Makes clear what the coordinates represent
// Could also consider a struct return
type Locator interface {
	getCoordinates(address string) (lat, lng float64, error)
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf := buff[nr:]
	}
	return //naked return
}
