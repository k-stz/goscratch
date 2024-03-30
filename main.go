package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	resident := &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return resident
}

func ShuffleAnimals() []string {
	animals := []string{
		"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog",
	}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

func closeThis(i int) func() string {
	return func() string {
		i++
		fmt.Println("closure called with i:", i)
		s := strconv.Itoa(i)
		return s
	}
}

func main() {
	fmt.Println(Application("sad❗sda"))
	fmt.Println(Application("sad🔍as"))
	fmt.Println(Application("asdf"))

	res := Replace("asdf#!##ok", '#', 'Z')
	fmt.Println("res:", res)
	fmt.Println(Replace("nooo 👎 bad!", '👎', '👍'))
}

func Replace(log string, oldRune, newRune rune) string {
	str := make([]rune, 0)
	for _, v := range log {
		if v == oldRune {
			str = append(str, newRune)
		} else {
			str = append(str, v)
		}
	}
	return string(str)
}

func Application(log string) string {
	for _, v := range log {
		switch v {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}
