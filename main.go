package main

import "fmt"

func PrintAthfan() string {
	return "Athfan"
}
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(PrintAthfan())
	fmt.Println(Hello("Athfan"))
}