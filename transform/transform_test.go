package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	expected := []string{"DOG", "CAT", "SHE3P", "SNAKE"}
	validateTranform(expected, applyTransform(getAnimals(), strings.ToUpper), t)
}

func TestCapitalize(t *testing.T) {
	expected := []string{"Dog", "Cat", "She3p", "Snake"}
	validateTranform(expected, applyTransform(getAnimals(), capitalize), t)
}

func TestReverse(t *testing.T) {
	expected := []string{"GoD", "tac", "P3ehs", "EKANS"}
	validateTranform(expected, applyTransform(getAnimals(), reverse), t)
}

func TestEcho(t *testing.T) {
	expected := []string{"DoG", "cat", "she3P", "SNAKE"}
	validateTranform(expected, applyTransform(getAnimals(), echo), t)
}

func validateTranform(expected []string, transformed []string, t *testing.T) {
	if len(expected) != len(transformed) {
		t.Errorf("Transformed slice has different length than the expected")
	}

	for i, item := range expected {
		if item != transformed[i] {
			t.Errorf("Value not transformed correctly")
		}
	}
}

func getAnimals() []string {
	return []string{"DoG", "cat", "she3P", "SNAKE"}
}

func capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}

func echo(s string) string {
	fmt.Println(s)
	return s
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
