package person

//go:generate go run ../.. -p . -s Person -w -f %s_gen.go -n
type Person struct {
	Name      string
	Age       int
	Interests []string

	flag bool

	// Some info for this field.
	Parent   *Person
	Children [5]*Person

	Friends []*Person

	Data map[string]interface{}
}

//go:generate go run ../.. -s Car -w
type Car struct {
	Name string
}
