package main

import "fmt"

// Person struct defines...
type Person struct {
	Name    string
	Age     int
	Address string
	Phone   string
}

// SetName sets name to person
func (p *Person) SetName(name string) {
	p.Name = name
	fmt.Println("set")
}

func main() {
	fabsPerson := Person{}
	fmt.Println(fabsPerson.Name)

	fabsPerson.SetName("Fabs")
	fmt.Println(fabsPerson.Name)
}
