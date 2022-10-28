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
	var email_income = NewStack[Email]()

	email_income.Push(Email{From: "a", To: "b"})
	email_income.Push(Email{From: "c", To: "d"})
	email_income.Push(Email{From: "e", To: "f"})
	email_income.Push(Email{From: "g", To: "h"})

	email_income.Pop()
	email_income.Pop()

	expected, _ := email_income.Top()
	assert.EqualValues(t, expected, Email{From: "c", To: "d"})
}

func TestStackMemoryUsage(t *testing.T) {
	email_income := NewStack[Email]()

	assert.EqualValues(t, email_income.Memory(), 55)

	for i := 0; i < 1_000_000; i++ {
		email_income.Push(Email{From: "a", To: "b"})
	}

	assert.EqualValues(t, email_income.Memory(), 7000061)

	for i := 0; i < 1_000_000; i++ {
		if err := email_income.Pop(); err != nil {
			t.Fatal(err)
		}
	}

	assert.EqualValues(t, email_income.Memory(), 55)
}
