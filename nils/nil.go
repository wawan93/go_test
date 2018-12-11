package nils

type First struct {
	Name string
}

func (first *First) GetName() string {
	if first == nil {
		return "none"
	}
	return first.Name
}

type Second struct {
	first *First
}

func (b *Second) GetFirst() *First {
	if b == nil {
		return nil
	}
	return b.first
}

type Third struct {
	second *Second
}

func (third *Third) GetSecond() *Second {
	if third == nil {
		return nil
	}
	return third.second
}

func (third *Third) GetFirstName() string {
	return third.GetSecond().GetFirst().GetName()
}
