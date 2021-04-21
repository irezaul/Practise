package main

import "fmt"

func main() {

	emails := map[string]string{"Reza": "mtmartadd@gmail.com", "Nahid": "nahidbd@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}
         //Range key
	for key := range emails {
		fmt.Println("Name: " + key)

	}

}
