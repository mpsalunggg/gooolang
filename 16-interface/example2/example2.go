package example2

import "fmt"

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}

func Example() {
	p := &Person{Name: "Putra", Age: 23}
	PrintDescription(p)

	p.Name = "Test"
	p.Age = 24

	PrintDescription(p)
}
