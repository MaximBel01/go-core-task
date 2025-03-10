package main

import (
	"fmt"
	"maps"
)

func Add(m map[string]int, key string, value int) {
	m[key] = value
	fmt.Printf("Added: %s -> %d\n", key, value)
}

func Remove(m map[string]int, key string) {
	delete(m, key)
	fmt.Printf("Removed: %s\n", key)
}

func Copy(original map[string]int) map[string]int {
	copyMap := make(map[string]int)
	maps.Copy(copyMap, original)
	return copyMap
}

func Exists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

func Get(m map[string]int, key string) (int, bool) {
	value, exists := m[key]
	return value, exists
}

func main() {
	strIntMap := make(map[string]int)
	Add(strIntMap, "Odin", 1)
	Add(strIntMap, "Dva", 2)
	Add(strIntMap, "Tri", 3)

	newMap := Copy(strIntMap)
	Remove(strIntMap, "Odin")
	fmt.Println("Final map: ", strIntMap)
	fmt.Println("Copied map: ", newMap)
	fmt.Printf("Does 'Dva' exist?: %v\n", Exists(strIntMap, "Dva"))
	val, exist := Get(strIntMap, "Dva")
	fmt.Printf("What number is 'Dva'? Number - %d, Exists - %v\n", val, exist)
}
