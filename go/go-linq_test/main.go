package main

import (
	"fmt"
	"strings"

	. "github.com/ahmetb/go-linq"
)

func main() {
	type Test struct {
		ID    int
		Email string
		Names []string
	}

	TestIds := []Test{}
	n := Test{ID: 1001, Email: "test@example.com", Names: []string{"name_1", "sub_name_1"}}
	TestIds = append(TestIds, n)
	n = Test{ID: 999, Email: "test2@example.com", Names: []string{"name_2"}}
	TestIds = append(TestIds, n)

	var output1 []string
	From(TestIds).Where(func(c interface{}) bool {
		return c.(Test).ID <= 1000
	}).Select(func(c interface{}) interface{} {
		return c.(Test).Email
	}).ToSlice(&output1)
	fmt.Println(output1)

	output2 := From(TestIds).SelectMany(
		func(test interface{}) Query {
			return From(test.(Test).Names)
		}).GroupBy(
		func(name interface{}) interface{} {
			return name
		}, func(name interface{}) interface{} {
			return name
		}).OrderByDescending(
		func(group interface{}) interface{} {
			return len(group.(Group).Group)
		}).Select(
		func(group interface{}) interface{} {
			return group.(Group)
		}).First()
	fmt.Println(output2)

	var output3 []string
	testSlice := []string{"a bb", "a ccc dddd eeeee ffffff"}
	From(testSlice).
		SelectManyT(func(testSlice string) Query {
			return From(strings.Split(testSlice, " "))
		}).
		GroupByT(
			func(word string) string { return word },
			func(word string) string { return word },
		).
		OrderByDescendingT(func(wordGroup Group) int {
			return len(wordGroup.Group)
		}).
		ThenByT(func(wordGroup Group) string {
			return wordGroup.Key.(string)
		}).
		Take(5).
		SelectT(func(wordGroup Group) string {
			return fmt.Sprintf("[%s / %d]", wordGroup.Key, len(wordGroup.Group))
		}).
		ToSlice(&output3)
	fmt.Println(output3)
}
