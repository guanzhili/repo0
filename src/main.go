package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 17},
	}

	for _, person := range people {
		person.Print()
		fmt.Println(person.Greet())
	}

	// Reading a file
	fileContent := readFile("example.txt")
	fmt.Println("File content:", fileContent)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {
	fmt.Println(p.Name, p.Age)
}

func (p Person) Greet() string {
	if p.Age > 18 {
		return fmt.Sprintf("Hello, %s!", p.Name)
	}
	return fmt.Sprintf("Hey, %s!", p.Name)
}

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var content string
	_, err = fmt.Fscanf(file, "%s", &content)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	return content
}


