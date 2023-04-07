package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// Item represents a product in the store
type Item struct {
	name  string
	price float64
	rating int
}

// User represents a user in the store
type User struct {
	username string
	password string
	loggedIn bool
}

// Store represents the store with a list of items and users
type Store struct {
	items []Item
	users []User
}

func (s *Store) register(username, password string) {
	s.users = append(s.users, User{username: username, password: password, loggedIn: false})
}

func (s *Store) authorize(username, password string) bool {
	for i := range s.users {
		if s.users[i].username == username && s.users[i].password == password {
			s.users[i].loggedIn = true
			return true
		}
	}
	return false
}

func (s *Store) rating(username string, itemName string, rating int) bool {
	loggedIn := false
	for i := range s.users {
		if s.users[i].username == username && s.users[i].loggedIn == true {
			loggedIn = true
			break
		}
	}

	if !loggedIn {
		return false
	}

	for i := range s.items {
		if strings.ToLower(s.items[i].name) == strings.ToLower(itemName) {
			s.items[i].rating = rating
			return true
		}
	}

	return false
}

func (s *Store) search(searchTerm string) []Item {
	results := []Item{}
	for i := range s.items {
		if strings.Contains(strings.ToLower(s.items[i].name), strings.ToLower(searchTerm)) {
			results = append(results, s.items[i])
		}
	}
	return results
}

func (s *Store) collectData() []float64 {
	data := []float64{}
	for i := range s.items {
		data = append(data, s.items[i].price)
	}
	return data
}

func main() {
	store := Store{
		items: []Item{
			{name: "item1", price: 10.5, rating: 0},
			{name: "item2", price: 20.0, rating: 0},
			{name: "item3", price: 15.0, rating: 0},
		},
		users: []User{},
	}

	fmt.Println("Store System")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Rate Item")
	fmt.Println("4. Search Item")
	fmt.Println("5. Collect Data")
	fmt.Println("0. Exit")
	fmt.Println()
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Option: ")
		scanner.Scan()
		input := scanner.Text()
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Option")
			continue
		}
	
		switch option {
		case 1:
			fmt.Print("Enter username: ")
			scanner.Scan()
			username := scanner.Text()
	
			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()
	
			store.register(username, password)
			fmt.Println("User registered successfully")
		case 2:
			fmt.Print("Enter username: ")
			scanner.Scan()
			username := scanner.Text()
	
			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()
	
			if store.authorize(username, password) {
				fmt.Println("Login successful")
			} else {
				fmt.Println("Login failed")
			}
		case 3:
			fmt.Print("Enter username: ")
			scanner.Scan()
			username := scanner.Text()
	
			fmt.Print("Enter item name: ")
			scanner.Scan()
			itemName := scanner.Text()
	
			fmt.Print("Enter rating (1-5): ")
			scanner.Scan()
			input = scanner.Text()
			rating, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid rating")
				continue
			}
	
			if store.rating(username, itemName, rating) {
				fmt.Println("Item rated successfully")
			} else {
				fmt.Println("Rating failed")
			}
		case 4:
			fmt.Print("Enter search term: ")
			scanner.Scan()
			searchTerm := scanner.Text()
	
			searchResults := store.search(searchTerm)
			fmt.Println("Search Results:")
			for i := range searchResults {
				fmt.Println("-", searchResults[i].name)
			}
		case 5:
			data := store.collectData()
			fmt.Println("Data:", data)
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Option")
		}
	}
	}