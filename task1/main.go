package main

import "fmt"

func main() {
	var data interface{} = map[string]interface{}{
		"person": map[string]string{
			"name":  "Zahid",
			"city":  "Dhaka",
			"hobby": "something",
		},
	}

	// Exercise:
	// 1. Use this data in your code and find out the city of the person.
	// 2. Use type assertion in one line. (Not in multiple line and not part by part)
	// 3. Print the value.

	d1 := data.(map[string]interface{})["person"].(map[string]string)["city"]
	fmt.Println(d1)

}
