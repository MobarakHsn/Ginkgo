package book

type Book struct {
	name  string
	Price int
}

func (b *Book) IsCheap() bool {
	if b.Price <= 500 {
		return true
	}
	return false
}
