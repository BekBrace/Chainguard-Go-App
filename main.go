package main // Define the package name, standard for an executable program in Go

import ( // Import necessary packages
	"encoding/json" // Package for encoding and decoding JSON
	"fmt"           // Package for formatted I/O with functions like Printf
	"log"           // Package for logging messages and errors
	"net/http"      // Package for HTTP client and server implementations
	"net/url"       // Package for URL parsing
	"os"            // Package for operating system functionality, including command line arguments

	// CaskaydiaCove NF
	// importing package cobra
	// Package cobra is a commander providing a simple interface to create  CLI interfaces.
	// go get -u github.com/spf13/cobra@latest
	"github.com/spf13/cobra" // Third-party library for building command line applications
)

// Book struct to hold book data
type Book struct {
	// Title of the book, mapped from JSON field "title"
	Title string `json:"title"`
}

// SearchResults struct to capture the JSON response from API
type SearchResults struct {
	// Slice of Book structs, mapped from JSON field "docs"
	Docs []Book `json:"docs"`
}

// Declare the root command for the CLI application
var rootCmd = &cobra.Command{
	// The first word used in the CLI command
	Use: "books",
	// Short description shown in the help text
	Short: "Books CLI is an application to fetch books by authors",
}

// Declare the search command to search books by author
var cmdSearch = &cobra.Command{
	// Defines how the command is used, with an argument for the author
	Use: "search [author]",
	// Short description of the command
	Short: "Search for books by a specific author",
	// Specifies that exactly one argument is required
	Args: cobra.ExactArgs(1),
	// Function to execute when this command is called
	Run: func(cmd *cobra.Command, args []string) {
		// Call searchBooks function with the author argument
		searchBooks(args[0])
	},
}

// Function to search books by author via API call
func searchBooks(author string) {
	safeAuthor := url.QueryEscape(author)

	// URL-encode the author name to safely include it in the URL
	url := fmt.Sprintf("https://openlibrary.org/search.json?author=%s", safeAuthor) // Format the search URL with the encoded author name

	// Make an HTTP GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		// Log and exit if there is an error fetching data
		log.Fatalf("Error fetching data: %s", err)
	}
	// Ensure the response body is closed after the function returns
	defer resp.Body.Close()

	var results SearchResults // Declare a variable to store decoded data
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Fatalf("Error decoding data: %s", err) // Log and exit if there is an error decoding data
	}

	if len(results.Docs) == 0 { // Check if no books were found
		fmt.Println("No books found for this author.") // Notify user no books were found
		return
	}

	fmt.Println("Books found:") // Print message indicating books were found
	for _, book := range results.Docs {
		fmt.Printf(" - %s\n", book.Title) // Print each book title
	}
}

// Main function to set up and execute the CLI application
func main() {
	rootCmd.AddCommand(cmdSearch)             // Add the search command as a sub-command of the root
	if err := rootCmd.Execute(); err != nil { // Execute the root command, handling any errors
		fmt.Fprintln(os.Stderr, err) // Print errors to standard error
		os.Exit(1)                   // Exit with status 1 on error
	}
}
