package main

import (
	"fmt"
	Shaper "goOnebyOne/packages"
)

type Person struct {
	firstname string
	lastname  string
	age       uint16
}

func Display(p *Person) error {
	p.firstname = "Arad"
	fmt.Println(p.firstname)
	return nil
}

func main() {

	person := Person{"abas", "khalili", 40}

	var ptr = &person

	Display(&person)

	fmt.Println(person.firstname)

	fmt.Println(ptr)

	var rec = Shaper.Rectangle{Lendth: 12.1, Width: 2.2}
	fmt.Printf("Area of Rec: %f\n", rec.Area())

	cir := Shaper.Circle{Radius: 5}
	fmt.Printf("Area of Cir: %f\n", cir.Area())

	fmt.Printf("XXX of Cir: %f\n", cir.XXX())

	var itf interface{}
	itf = 159
	fmt.Printf("data type interface is : %d \n", itf)
	itf = "Hellp Golang"
	fmt.Printf("data type interface is : %s \n", itf)
	itf = 3.14
	fmt.Printf("data type interface is : %f \n", itf)

	var anonymousFunction = func(itf interface{}) {
		int_itf, itf_converted_without_err := itf.(int)
		fmt.Println("Type Assertion in Anonymouse func :", int_itf, itf_converted_without_err)
	}

	//type assertion
	anonymousFunction(12450.5)
	anonymousFunction(12450)

	//type assertion
	itf = 123022.2
	int_itf, itf_converted_without_err := itf.(int)
	fmt.Println(`Type Assertion :`, int_itf, itf_converted_without_err)

	//type assertion
	itf = 123022
	int_itf, itf_converted_without_err = itf.(int)
	fmt.Println("Type Assertion :", int_itf, itf_converted_without_err)

	//type assertion
	itf = "Golang"
	str_itf, itf_converted_without_err := itf.(string)
	fmt.Println("Type Assertion :", str_itf, itf_converted_without_err)

	//Closure
	odd := calculate(3)
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	// closure => This helps us to work with multiple data in isolation from one another.
	even := calculate(8)
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())
}

// Closure function
func calculate(start int) func() int {
	num := start
	return func() int {
		num = num + 2
		return num
	}
}
