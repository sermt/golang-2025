package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	name := "John Doe"
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", name, age)

	editAgeToAdultYears(agePointer)

	fmt.Println("My adult years are:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
