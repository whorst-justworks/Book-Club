type customer struct {
	balance float64
}

func (c *customer) deposit(amount float64) {
	c.balance += amount
}

func main() {
	c := &customer{balance: 100.}
	c.deposit(50.)
	fmt.Printf("balance %.2f\n", c.balance)
	// Result: balance 150.00
}
