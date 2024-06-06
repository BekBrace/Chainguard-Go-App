# Books CLI

Books CLI is a command-line application built in Go that allows users to search for books by author names using the Open Library API. This tool is ideal for quick searches to find books from specific authors directly from your terminal.

## Features

- **Search by Author**: Enter an author's name to fetch a list of books written by them.
- **Simple CLI**: Easy-to-use command line interface.
- **Immediate Results**: Quickly obtain a list of books without the need for a GUI.

## Prerequisites

Before you begin, ensure you have the Go programming language installed on your system. You can download it from [here](https://golang.org/dl/).

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/books-cli.git
cd books-cli
```
Build the application:
```bash
go build
```
Usage
To search for books by a specific author, run the following command in your terminal:
```bash
./books search "Author Name"
```
Replace "Author Name" with the name of the author whose books you want to search.
# Examples
## Here's how you can search for books by J.K. Rowling:
```bash
./books search "J.K. Rowling"
```
Output:
```Books found:
 - Harry Potter and the Sorcerer's Stone
 - Harry Potter and the Chamber of Secrets
 - ...
```
