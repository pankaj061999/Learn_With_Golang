package main

import (
	"fmt"
	"sort"
)

func PrintMapValue(m map[string]int64) {

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("Key sort value ==>>", keys)

	// sorted format
	for _, key := range keys {
		value := m[key]
		fmt.Println("Print map key ===>>", key, "map value ==>>", value)
	}

	// for key, value := range m {
	// 	fmt.Println("Print map key ===>>", key, "map value ==>>", value)
	// }
}

func main() {

	// syntax here key type is string and map value type is int
	myMap := map[string]int64{
		"one":   123,
		"two":   213,
		"three": 312,
		"four":  412,
	}

	// myMap["four"] = 412
	fmt.Println("map file", myMap)
	PrintMapValue(myMap)

	// for i, v := range myMap {
	// 	fmt.Println("Print map value ", v, "Key value ==>> ", i)
	// }

}
