package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	numDecimal := 42
	numOctal := 052
	numHexadecimal := 0x2A
	nice := 6.9
	name := "Golang"
	isActive := true
	complexNum := complex(float32(1), float32(2))

	variables := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		nice,
		name,
		isActive,
		complexNum,
	}

	fmt.Println("Variable types:")
	for _, v := range variables {
		fmt.Printf("%v: %s\n", v, reflect.TypeOf(v))
	}

	concatenated := concatenateVariables(variables)
	fmt.Println("\nConcatenated string:", concatenated)

	runes := []rune(concatenated)

	salt := "go-2024"
	hash := addSaltAndHash(runes, salt)
	fmt.Println("\nSHA256 Hash:", hash)
}

func concatenateVariables(vars []any) string {
	var result string
	for _, v := range vars {
		result += variableToString(v)
	}
	return result
}

func variableToString(v any) string {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case complex64:
		return fmt.Sprint(val)
	default:
		return ""
	}
}

func addSaltAndHash(runes []rune, salt string) string {
	saltRunes := []rune(salt)
	mid := len(runes) / 2
	saltedRunes := append(runes[:mid], append(saltRunes, runes[mid:]...)...)
	saltedStr := string(saltedRunes)
	hash := sha256.Sum256([]byte(saltedStr))
	return hex.EncodeToString(hash[:])
}
