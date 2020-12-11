package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var cache = make(map[string]int)

func input(elements map[string][]string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\s*\d* (.*)\sbag(s)*\.*`)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "contain")
		var bags []string
		for _, x := range strings.Split(s[1], ",") {
			bags = append(bags, re.ReplaceAllString(x, `$1`))
		}
		elements[strings.Split(s[0], " bags")[0]] = bags
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func check(key string, elements map[string][]string) int {
	if val, ok := cache[key]; ok {
		return val
	}
	for _, x := range elements[key] {
		if x == "shiny gold" {
			cache[key] = 1
			return 1
		}
		if x == "no other" {
			cache[key] = 0
			return 0
		}
		if check(x, elements) == 1 {
			cache[key] = 1
			return 1
		}
	}
	cache[key] = 0
	return 0
}

func main() {
	elements := make(map[string][]string)
	input(elements)
	var sum = 0
	for k := range elements {
		sum += check(k, elements)
	}
	fmt.Println(sum)
}
