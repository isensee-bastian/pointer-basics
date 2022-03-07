package main

import (
	"testing"
)

func TestSwapNaive(t *testing.T) {
	left := 1
	right := 2
	swapNaive(left, right)

	if left != 1 || right != 2 {
		t.Errorf("Expected left 2 and right 1 but got %v and %v\n", left, right)
	}
}

func TestSwapByReturn(t *testing.T) {
	left := 1
	right := 2
	left, right = swapByReturn(left, right)

	if left != 2 || right != 1 {
		t.Errorf("Expected left 2 and right 1 but got %v and %v\n", left, right)
	}
}

func TestSwapByPointers(t *testing.T) {
	left := 1
	right := 2
	swapByPointers(&left, &right)

	if left != 2 || right != 1 {
		t.Errorf("Expected left 2 and right 1 but got %v and %v\n", left, right)
	}
}

func TestSwapFirstElementsSlice(t *testing.T) {
	slice := []int{10, 20}
	swapFirstElementsSlice(slice)

	if slice[0] != 20 || slice[1] != 10 {
		t.Errorf("Expected slice elements 20, 10 but got %v and %v\n", slice[0], slice[1])
	}
}

func TestSwapMapElements(t *testing.T) {
	textToInt := map[string]int{"apple": 10, "banana": 20}
	swapMapElements(textToInt, "apple", "banana")

	if textToInt["apple"] != 20 || textToInt["banana"] != 10 {
		t.Errorf("Expected value for key apple 20, value for key banana 10, but got %v\n", textToInt)
	}
}

func TestLevelupNaive(t *testing.T) {
	player := player{name: "Karl", level: 1}
	levelupNaive(player, 2)

	if player.level != 1 {
		t.Errorf("Expected level 1 but got %v\n", player.level)
	}
}

func TestLevelupByPointer(t *testing.T) {
	player := player{name: "Karl", level: 1}
	levelupByPointer(&player, 2)

	if player.level != 3 {
		t.Errorf("Expected level 3 but got %v\n", player.level)
	}
}
