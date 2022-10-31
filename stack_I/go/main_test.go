package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Email struct {
	From string
	To   string
}

func TestStack(t *testing.T) {
	emailIncome := NewStack[Email]()

	emailIncome.Push(Email{From: "a", To: "b"})
	emailIncome.Push(Email{From: "c", To: "d"})
	emailIncome.Push(Email{From: "e", To: "f"})
	emailIncome.Push(Email{From: "g", To: "h"})

	if err := emailIncome.Pop(); err != nil {
		panic(err)
	}

	if err := emailIncome.Pop(); err != nil {
		panic(err)
	}

	expected, _ := emailIncome.Top()
	assert.EqualValues(t, expected, Email{From: "c", To: "d"})
}

func TestStackMemoryUsage(t *testing.T) {
	emailIncome := NewStack[Email]()

	assert.EqualValues(t, emailIncome.Memory(), 55)

	for i := 0; i < 1_000_000; i++ {
		emailIncome.Push(Email{From: "a", To: "b"})
	}

	assert.EqualValues(t, emailIncome.Memory(), 7000061)

	for i := 0; i < 1_000_000; i++ {
		if err := emailIncome.Pop(); err != nil {
			t.Fatal(err)
		}
	}

	assert.EqualValues(t, emailIncome.Memory(), 55)
}
