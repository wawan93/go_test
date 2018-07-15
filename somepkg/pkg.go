package somepkg

type SomeType int

func (SomeType) SomeFunc() string {
	return "hello"
}

func (SomeType) AnotherInternalFunction() string {
	return "another"
}

