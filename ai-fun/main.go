package main

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/tozd/go/errors"
	"gitlab.com/tozd/go/fun"
)

func main() {
	ctx := context.Background()

	// Perform addition, squaring, concatenation, and reversing
	adderResult := add(ctx, 38, 4)
	squaredResult := square(ctx, adderResult)
	concatResult := concatenate(ctx, "Hello", "World")
	reversedResult := reverse(ctx, concatResult)

	// Print final results
	fmt.Printf("Squared result: %d\n", squaredResult)
	fmt.Printf("Concatenated and reversed result: \"%s\"\n", reversedResult)
}

// Function to add two numbers
func add(ctx context.Context, num1, num2 int) int {
	adder := fun.Go[int, int]{
		Fun: func(_ context.Context, input ...int) (int, errors.E) {
			if len(input) < 2 {
				return 0, errors.New("insufficient arguments for addition")
			}
			sum := input[0] + input[1]
			log.Printf("Adding %d + %d = %d", input[0], input[1], sum)
			return sum, nil
		},
	}

	if errE := adder.Init(ctx); errE != nil {
		log.Fatalf("Error initializing adder: %v", errE)
	}

	output, errE := adder.Call(ctx, num1, num2)
	if errE != nil {
		log.Fatalf("Error calling adder: %v", errE)
	}

	fmt.Printf("Output of adder: %d\n", output)
	return output
}

// Function to square a number
func square(ctx context.Context, num int) int {
	squarer := fun.Go[int, int]{
		Fun: func(_ context.Context, input ...int) (int, errors.E) {
			if len(input) < 1 {
				return 0, errors.New("no input for squaring")
			}
			square := input[0] * input[0]
			log.Printf("Squaring %d = %d", input[0], square)
			return square, nil
		},
	}

	if errE := squarer.Init(ctx); errE != nil {
		log.Fatalf("Error initializing squarer: %v", errE)
	}

	output, errE := squarer.Call(ctx, num)
	if errE != nil {
		log.Fatalf("Error calling squarer: %v", errE)
	}

	fmt.Printf("Output after squaring: %d\n", output)
	return output
}

// Function to concatenate two strings
func concatenate(ctx context.Context, str1, str2 string) string {
	concatenator := fun.Go[string, string]{
		Fun: func(_ context.Context, input ...string) (string, errors.E) {
			if len(input) < 2 {
				return "", errors.New("insufficient arguments for concatenation")
			}
			result := input[0] + " " + input[1]
			log.Printf("Concatenating \"%s\" and \"%s\" = \"%s\"", input[0], input[1], result)
			return result, nil
		},
	}

	if errE := concatenator.Init(ctx); errE != nil {
		log.Fatalf("Error initializing concatenator: %v", errE)
	}

	output, errE := concatenator.Call(ctx, str1, str2)
	if errE != nil {
		log.Fatalf("Error calling concatenator: %v", errE)
	}

	fmt.Printf("Output of concatenator: \"%s\"\n", output)
	return output
}

// Function to reverse a string
func reverse(ctx context.Context, str string) string {
	reverser := fun.Go[string, string]{
		Fun: func(_ context.Context, input ...string) (string, errors.E) {
			if len(input) < 1 {
				return "", errors.New("no input for reversing")
			}
			reversed := reverse_str(input[0])
			log.Printf("Reversing \"%s\" = \"%s\"", input[0], reversed)
			return reversed, nil
		},
	}

	if errE := reverser.Init(ctx); errE != nil {
		log.Fatalf("Error initializing reverser: %v", errE)
	}

	output, errE := reverser.Call(ctx, str)
	if errE != nil {
		log.Fatalf("Error calling reverser: %v", errE)
	}

	fmt.Printf("Output after reversing: \"%s\"\n", output)
	return output
}

// Helper function to reverse a string
func reverse_str(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
