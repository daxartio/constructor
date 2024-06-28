package mystruct

//go:generate go run ../.. -p . -s MyStruct -w -f %s_gen.go -n
type MyStruct struct {
	Name      string
	Age       int
	Interests []string

	flag bool

	// Some info for this field.
	MyParent *MyStruct
	Children [5]*MyStruct

	Friends []*MyStruct

	Data map[string]interface{}
}

//go:generate go run ../.. -s MyStruct2 -w
type MyStruct2 struct {
	Field1 string
}
