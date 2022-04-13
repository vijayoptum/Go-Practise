package main

import (
	"fmt"
	"strings"
)

func contains() {
	s := "abcdeabcdeabcde"
	r := "ab"
	println(strings.Contains(s, r))
}
func compare() {
	fmt.Println(strings.Compare("A", "B"))             // A < B
	fmt.Println(strings.Compare("B", "A"))             // B > A
	fmt.Println(strings.Compare("Japan", "Australia")) // J > A
	fmt.Println(strings.Compare("Australia", "Japan")) // A < J
	fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", ""))
	fmt.Println(strings.Compare("", " ")) // Space is less
}
func containsAny() {
	s := "abcdeabcdeabcde"
	r := "ab"
	println(strings.ContainsAny(s, r))
}
func count() {
	s := "abcdeabcdeabcde"
	r := "ab"
	println(strings.Count(s, r))
}
func equalfold() {
	s := "abcdeabcdeabcde"
	r := "abcdeabcdeabc"
	println(strings.EqualFold(s, r))
}

func fields() {
	s := "This is written by gaurav singh"

	test := (strings.Fields(s))
	for _, v := range test {
		println(v)
	}
}

func index() {
	s := "abcdeabcdeabcde"
	r := "cd"
	println(strings.Index(s, r))
}

func indexAny() {
	s := "abcdeabcdeabcde"
	r := "cd"
	println(strings.IndexAny(s, r))
}

func join() {
	textNum := []string{"1", "2", "3", "4", "5"}
	println(strings.Join(textNum, ""))

}
func main() {
	//compare()
	//contains()
	//containsAny()
	//count()
	//equalfold()
	//fields()
	//index()
	//indexAny()
	join()
}
