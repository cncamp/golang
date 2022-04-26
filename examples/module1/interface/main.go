package main

import "fmt"

type IF interface {
	getName() string
}

type Fruit interface {
	Describe() string
}

type Apple struct {
	Color string
	Shape string
	Taste string
}

func (a Apple) Describe() string {
	return fmt.Sprintf("Apple: the color is %s, the shape is %s, the taste is %s", a.Color, a.Shape, a.Taste)
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	var interfaces []IF
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())

	apple := new(Apple)
	apple.Color = "Red"
	apple.Shape = "Round"
	apple.Taste = "Sweet"
	description := apple.Describe()
	fmt.Print(description)
}
