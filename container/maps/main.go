package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notabd",
	}

	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	m3 = make(map[string]int)
	m3["age"] = 20
	fmt.Println(m3)

	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values...")
	courseName := m["course"]
	fmt.Println(courseName)
	if caurseName, ok := m["caurse"]; ok {
		fmt.Println("值存在", caurseName)
	} else {
		fmt.Println("值不存在")
	}
	fmt.Println("Deleting values...")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
