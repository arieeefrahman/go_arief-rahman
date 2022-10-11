package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/google/uuid"
)

type Book struct {
	Id uuid.UUID
	Title string
	Price int
	Category string
}

type NameSorter []Book

var Books = []Book{}

func (n NameSorter) Len() int {
	return len(n)
}

func (n NameSorter) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NameSorter) Less(i, j int) bool {
	return n[i].Title < n[j].Title
}

func Menu() {
	choice := pick()
	HandlingInput(choice)
	fmt.Println("Bye...")
}

func MenuList() {
	fmt.Println("\n===BOOK MANAGEMENT===")
	fmt.Println("1. Get all books")
	fmt.Println("2. Create a book")
	fmt.Println("3. Update a book")
	fmt.Println("4. Delete a book")
	fmt.Println("5. Exit")
}

var (
	pick = func () int {
		MenuList()
		fmt.Print("Choose your menu: ")
		var c int
		if _, err := fmt.Scanln(&c); err != nil {
			log.Fatal(err)
		}
		
		return c
	}
)

func HandlingInput(choice int) {

	// error handling: prevent input number outside 1 to 5
	var bound = func () {
		for choice < 1 || choice > 5 {
			fmt.Println()
			fmt.Println("Wrong Number! Please enter between 1 to 5")
			fmt.Println("Enter 5 to exit")
			choice = pick()
		}
	}

	bound()

	// stop until user exit by enter input 5
	for choice >= 1 && choice <= 4 {
		HandlingChoice(choice)

		// user can choose from menu again
		choice = pick()
		bound()
	}
}

var GenerateUuid = func() uuid.UUID{
	return uuid.New()
}

func HandlingChoice(choice int) []Book{
	switch choice {
	case 1:
		result := Books
		sort.Sort(NameSorter(result))
		fmt.Println("\nAll books")
			
		for i := 0; i < len(result); i++ {
			fmt.Println("===")
			fmt.Printf("ID:   %s\n", result[i].Id)
			fmt.Printf("Title:   %s\n", result[i].Title)
			fmt.Printf("Price:   %d\n", result[i].Price)
			fmt.Printf("Category:   %s\n", result[i].Category)
			fmt.Printf("===\n")
		}
		
	case 2:
		input := Book{}
		input.Id = GenerateUuid()
		
		fmt.Print("enter title: ")
		if _, err := fmt.Scanln(&input.Title); err != nil {
			log.Fatal(err)
		}

		fmt.Print("enter price: ")
		if _, err := fmt.Scanln(&input.Price); err != nil {
			log.Fatal(err)
		}

		fmt.Print("enter category: ")
		if _, err := fmt.Scanln(&input.Category); err != nil {
			log.Fatal(err)
		}	

		Books = append(Books, input)
		fmt.Println("Book added!")
	case 3:
		input := Book{}
		
		// read input from user
		var inputStringId string
		fmt.Print("enter id: ")
		if _, err := fmt.Scanln(&inputStringId); err != nil {
			log.Fatal(err)
		}

		input.Id, _ = uuid.Parse(inputStringId)
		
		fmt.Print("enter title: ")
		if _, err := fmt.Scanln(&input.Title); err != nil {
			log.Fatal(err)
		}

		fmt.Print("enter price: ")
		if _, err := fmt.Scanln(&input.Price); err != nil {
			log.Fatal(err)
		}

		fmt.Print("enter category: ")
		if _, err := fmt.Scanln(&input.Category); err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(Books); i++ {
			if input.Id == Books[i].Id{
				Books[i].Title = input.Title
				Books[i].Price = input.Price
				Books[i].Category = input.Category
				break
			}
		}

		fmt.Println("Book updated!")
	case 4:
		input := Book{}

		// read input from user
		var inputStringId string
		fmt.Print("enter id: ")
		if _, err := fmt.Scanln(&inputStringId); err != nil {
			log.Fatal(err)
		}

		input.Id, _ = uuid.Parse(inputStringId)

		for i := 0; i < len(Books); i++ {
			if input.Id == Books[i].Id {
				Books = append(Books[:i], Books[i+1:]...)
			}
		}

		fmt.Println("Book deleted!")
	}

	return Books
}

func main() {
	Menu()
}