package main

import "fmt"

func maps() {

	fmt.Println("========== GO MAP PRACTICE ==========")

	// 1. Empty Map
	emptyMap := map[string]int{}
	fmt.Println("Empty Map:", emptyMap)

	// 2. Map Initialization
	user := map[string]int{
		"age":    23,
		"salary": 50000,
	}
	fmt.Println("User Map:", user)

	// 3. Using make()
	ages := make(map[string]int)

	ages["Korim"] = 44
	ages["Fokier"] = 55
	ages["Rahim"] = 30

	fmt.Println("Ages:", ages)

	// 4. Insert New Data
	ages["Arko"] = 25
	fmt.Println("After Insert:", ages)

	// 5. Update Existing Value
	ages["Arko"] = 26
	fmt.Println("After Update:", ages)

	// 6. Access Value
	fmt.Println("Arko Age:", ages["Arko"])

	// 7. Check Key Exists
	value, exists := ages["Arko"]

	fmt.Println("Value:", value)
	fmt.Println("Exists:", exists)

	// 8. Check Non-existing Key
	value, exists = ages["Sakib"]

	fmt.Println("Sakib Value:", value)
	fmt.Println("Sakib Exists:", exists)

	// 9. Delete Key
	delete(ages, "Korim")
	fmt.Println("After Delete:", ages)

	// 10. Map Length
	fmt.Println("Map Length:", len(ages))

	// 11. Loop Through Map
	fmt.Println("Loop Through Map:")

	for key, value := range ages {
		fmt.Println(key, ":", value)
	}

	// 12. Only Keys
	fmt.Println("Only Keys:")

	for key := range ages {
		fmt.Println(key)
	}

	// 13. Only Values
	fmt.Println("Only Values:")

	for _, value := range ages {
		fmt.Println(value)
	}

	// 14. Empty Map Check
	if len(ages) == 0 {
		fmt.Println("Map is Empty")
	} else {
		fmt.Println("Map is Not Empty")
	}

	// 15. Clear Entire Map (Go 1.21+)
	clear(ages)
	fmt.Println("After Clear:", ages)

	studentdffd := map[string][]string{
		"Arko":  {"Go", "Node.js", "TypeScript"},
		"Rahim": {"Java", "Python", "C++"},
		"Karim": {"PHP", "Laravel", "MySQL"},
	}

	fmt.Println(studentdffd)

	student := map[int]map[string]string{
		1: {
			"name":    "Arko",
			"email":   "arko@gmail.com",
			"address": "Dhaka",
		},
		2: {
			"name":    "Rahim",
			"email":   "rahim@gmail.com",
			"address": "Rangpur",
		},
	}

	fmt.Println(student)

	// Access
	fmt.Println(student[1]["name"])
	fmt.Println(student[2]["email"])
}
