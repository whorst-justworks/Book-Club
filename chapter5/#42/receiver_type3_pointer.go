type customer struct {
	data *data
}

type data struct {
	balance float64
}

func (c customer) deposit(amount float64) {
	c.data.balance += amount
}

func main() {
	c := customer{data: &data{
		balance: 100.
	}}
	c.deposit(50.)
	fmt.Printf("balance %.2f\n", c.balance)
	// Result: balance ?
}
