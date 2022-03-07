package main

import (
	"fmt"
)

// The naive approach swap does not work because arguments from calling
// code are copied to the parameters. The actual parameters are only
// valid inside the functions body.
func swapNaive(left, right int) {
	temp := left
	left = right
	right = temp
	fmt.Printf("left: %v, right: %v\n", left, right)
}

// This a simple solution for our task but it requires returning values
// and reassigning on the callers side.
func swapByReturn(left, right int) (int, int) {
	return right, left
}

// By passing addresses in the form of pointers we can modify the referenced
// values for the caller.
func swapByPointers(left, right *int) {
	// Note that go allows value swapping in a single line. This can be useful
	// sometimes, but of course one can also use a temporary variable and swap
	// step by step.

	// temp := *left
	// *left = *right
	// *right = temp

	*right, *left = *left, *right
}

// Slices are passed as references (pointers under the hood). Hence, we can
// change elements and see the change on the callers side, too.
func swapFirstElementsSlice(slice []int) {
	if len(slice) >= 2 {
		slice[0], slice[1] = slice[1], slice[0]
	}
}

// Maps are also passed as references. We can change elements and see the change
// on the callers side, too.
func swapMapElements(textToInt map[string]int, first string, second string) {
	temp := textToInt[first]
	textToInt[first] = textToInt[second]
	textToInt[second] = temp
}

// We define a struct type to check pointer semantics for struct types (see functions
// below).
type player struct {
	name  string
	level int
}

// Struct types behave the same as other types with regards to pointer
// semantics. Without a pointer, structs will be copied upon function
// call. Hence, changes are only visible inside the copied struct in
// the function body, not on the outside.
func levelupNaive(player player, delta int) {
	player.level += delta
	fmt.Println("player.level:", player.level)
}

// If we want to make changes durable for the caller, we can pass a
// pointer to the struct. Note that dereferencing happens automatically
// in this case, we don't need to write *player in the assignment.
func levelupByPointer(player *player, delta int) {
	player.level += delta
}

func main() {
	var message string = "Hello world!"
	fmt.Println("message:", message)

	// Pointer types are declared by using an asterik before the usual
	// type identifier. Pointer values are retrieved by using the address
	// operator & in front of a usual variable.
	// When printing a pointer variable you can see that it contains a memory
	// address. In our case it stores the address to our string message.
	var messagePointer *string = &message
	fmt.Println("messagePointer:", messagePointer)
	// We can dereference a pointer using the asterik operator * in order to
	// get the actual value that the pointer is referencing.
	fmt.Println("dereferenced messagePointer:", *messagePointer)

	// Note that messagePointer is referencing the same value that is held by
	// our usual variable message. Therefore, when we change the value of
	// message the changed value is also visible when dereferencing messagePointer.
	message = "Banana banana!"
	fmt.Println("message:", message)
	fmt.Println("messagePointer:", messagePointer)
	fmt.Println("dereferenced messagePointer:", *messagePointer)

	// Vice versa when we change the value that messagePointer is referencing
	// we change the same value that is stored in message.
	*messagePointer = "Grapefruit!"
	fmt.Println("message:", message)
	fmt.Println("messagePointer:", messagePointer)
	fmt.Println("dereferenced messagePointer:", *messagePointer)

	// Pointers can be created for any type. A pointer that is not pointing
	// anywhere contains nil which is a special value meaning "nothing".
	var numberPointer *int
	fmt.Println("numberPointer:", numberPointer)

	number := 42
	// Once we assign an actual address the pointers value changes from nil
	// to the address.
	numberPointer = &number
	fmt.Println("numberPointer:", numberPointer)
	// We can set it to nil if we want to signal that the pointer is pointing
	// nowhere. This can be useful, e.g. think of error handling where a nil
	// error signals that everything went fine.
	numberPointer = nil
	fmt.Println("numberPointer:", numberPointer)

	// Pointer arithmetic is not possible in Golang for safety reasons.
	// numberPointer++
	// numberPointer += 10
}
