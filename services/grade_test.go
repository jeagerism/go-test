package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "A", score: 80, expected: "A"},
		{name: "B", score: 70, expected: "B"},
		{name: "C", score: 60, expected: "C"},
		{name: "D", score: 50, expected: "D"},
		{name: "F", score: 40, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			expect := c.expected

			assert.Equal(t, expect, grade)
			// if grade != expect {
			// 	t.Errorf("got %v expected %v", grade, expect)
			// }
		})
	}

}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	//Output: A
}
