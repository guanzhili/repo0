package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

var people []Person
var products []Product

// Simulating a global lock for logging, which might become a bottleneck
var logMutex sync.Mutex

func main() {
	// Initialize people and products data
	initializeData()

	// Simulate concurrent processing of people and products
	var wg sync.WaitGroup
	wg.Add(2)

	go processPeople(&wg)
	go processProducts(&wg)

	// Wait for both goroutines to complete
	wg.Wait()

	// Log final activity
	logActivity("Processing completed")

	// Simulate file reading with potential errors
	fileContent := readFile("example.txt")
	fmt.Println("File content:", fileContent)
}

// Initialize data with some people and products
func initializeData() {
	people = []Person{
		{"Alice", 25},
		{"Bob", 17},
		{"Charlie", 30},
		{"Dave", 42},
		{"Eve", 28},
	}

	products = []Product{
		{1, "Laptop", 999.99},
		{2, "Smartphone", 599.99},
		{3, "Tablet", 399.99},
		{4, "Headphones", 199.99},
		{5, "Monitor", 249.99},
	}
}

// Process people - example of potential logic issue (handling data incorrectly)
func processPeople(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, person := range people {
		// Potential issue: Incorrect age categorization logic
		if person.Age < 18 {
			logError(fmt.Sprintf("Underage person: %s", person.Name))
		} else if person.Age > 40 {
			logError(fmt.Sprintf("Person over 40: %s", person.Name))
		}
		printPersonDetails(person)
	}
}

// Process products - introduces potential performance bottleneck in logging
func processProducts(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, product := range products {
		// Potential performance issue: logging for every product processed
		logActivity(fmt.Sprintf("Processing product: %s", product.Name))
		// Some basic processing logic
		time.Sleep(time.Millisecond * 100) // Simulate delay
	}
}

// Print person details
func printPersonDetails(person Person) {
	// Potential issue: printing each person detail every time could be inefficient
	logActivity(fmt.Sprintf("Person details: %s, Age: %d", person.Name, person.Age))
}

// Log error with potential issue of repetitive logging
func logError(message string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	log.Println("ERROR: " + message)
}

// Log activity with potential performance bottleneck (frequent file operations)
func logActivity(activity string) {
	// Open the log file with every log entry
	logFile, err := os.OpenFile("activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file", err)
	}
	defer logFile.Close()

	// Create a logger to log the activity with timestamps
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(activity)
}

// Read file content with a potential issue in error handling
func readFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err) // Fatal log - terminates program
	}
	defer file.Close()

	var content string
	_, err = fmt.Fscanf(file, "%s", &content)
	if err != nil {
		log.Fatal("Error reading file", err) // Fatal log - terminates program
	}
	return content
}

// Function with memory leak due to not closing resources
func processLargeData() {
	var data []byte
	for i := 0; i < 1000000; i++ {
		data = append(data, byte(i))
	}

	// Memory is being allocated, but not properly released (potential memory leak)
	// Data isn't being processed or freed after usage
	fmt.Println("Processed large data")
}

// Example of incorrect data type usage
func processInvalidData() {
	var data interface{} = "This should be an integer" // Incorrect data type

	// Potential panic: incorrect type assertion
	if intData, ok := data.(int); ok {
		fmt.Println("Integer data:", intData)
	} else {
		logError("Failed to assert data to int")
	}
}

// Example of an unoptimized algorithm
func inefficientSort() {
	// Bubble Sort - highly inefficient for large datasets
	data := []int{5, 2, 9, 1, 5, 6}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				// Swap elements
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println("Sorted data:", data)
}

// Add a long-running operation with poor error handling
func longRunningOperation() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			// Simulating an error condition but not handling it properly
			logError(fmt.Sprintf("Error during operation at iteration %d", i))
		}
		time.Sleep(time.Second)
	}
}

