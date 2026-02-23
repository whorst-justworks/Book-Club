func (l locator) getCoordinates(ctx context.Context, address string) (lat, lng float64, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if ctx.Err() != nil {
		return 0, 0, err // wrong error value
	}

	// Get coodinates

	return // generally shouldn't mix naked and explicit returns
}
