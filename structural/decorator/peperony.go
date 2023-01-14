package decorator


type peperony struct {
}

func (p *peperony) getPrice() int {
	return 15
}
