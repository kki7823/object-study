package ticket

type Audience struct {
	Bag *Bag
}

func (a *Audience) GetBag() *Bag {
	return a.Bag
}
