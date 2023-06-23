package main

import (
	"reflect"
	"testing"
)

type TableTestCase struct {
	Name          string
	Input         interface{}
	ExpectedCalls []string
}

func TestWalk(t *testing.T) {
	cases := []TableTestCase{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Foobarius"},
			ExpectedCalls: []string{"Foobarius"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Foobarius", "Boston"},
			ExpectedCalls: []string{"Foobarius", "Boston"},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string
			walk(testCase.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, testCase.ExpectedCalls)
			}
		})
	}
}
