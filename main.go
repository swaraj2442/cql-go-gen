package main

import (
	"cql-gen/generator"
	"cql-gen/parser"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Command line flags
	queryFile := flag.String("query", "", "Cassandra CQL query file to parse")
	outputFile := flag.String("output", "output.go", "Output file for the generated Go code")
	flag.Parse()

	if *queryFile == "" {
		fmt.Println("Error: query file must be provided")
		os.Exit(1)
	}

	// Read the query file
	cql, err := os.ReadFile(*queryFile)
	if err != nil {
		fmt.Printf("Error reading query file: %v\n", err)
		os.Exit(1)
	}

	// Parse the CQL
	parsedQueries, err := parser.ParseCQL(string(cql))
	if err != nil {
		fmt.Printf("Error parsing CQL: %v\n", err)
		os.Exit(1)
	}

	// Generate Go code
	err = generator.GenerateGoCode(parsedQueries, *outputFile)
	if err != nil {
		fmt.Printf("Error generating Go code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Go code generated successfully.")
}
