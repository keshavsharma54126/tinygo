package main

import (
	"fmt"
	"sort"
)

type sortedBooks []Book

func (b sortedBooks) Len() int {
	return len(b)
}

func (b sortedBooks) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b sortedBooks) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

func sortBooks(books []Book) []Book {
	sort.Sort(sortedBooks(books))
	// sort.Slice(books, func(i, j int) bool {
	// 	if books[i].Author != books[j].Author {
	// 		return books[i].Author < books[j].Author
	// 	}
	// 	return books[i].Title < books[j].Title
	// })
	return books
}

func commonBookCounts(bookworms []Bookworms) map[Book]uint {
	commonBookCount := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			commonBookCount[book]++
		}
	}
	return commonBookCount
}

func findCommonBookworms(bookwroms []Bookworms) []Book {
	commonBooksCounts := commonBookCounts(bookwroms)
	var commonBooks []Book
	for book := range commonBooksCounts {
		if commonBooksCounts[book] > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	sortBooks(commonBooks)
	return commonBooks
}

func displayBooks(books []Book) {
	for _, Book := range books {
		fmt.Printf("-%s by %s\n", Book.Title, Book.Author)
	}
}

func main() {
	bookworms, err := loadBookworms("./testdata/bookworm.json")
	if err != nil {
		fmt.Println("an error occured while loading file")
	}
	commonBooks := findCommonBookworms(bookworms)
	fmt.Println("Here are the common books:")
	displayBooks(commonBooks)
}
