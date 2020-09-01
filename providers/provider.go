package providers

type provider interface {
	Start()
}

// Start ...
func Start(p provider) {
	p.Start()
}
