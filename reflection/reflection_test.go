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

type BankAccount struct {
	Amount   int
	Currency string
}

type Person struct {
	Name    string
	Account BankAccount
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
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Foo", 33},
			ExpectedCalls: []string{"Foo"},
		},
		{
			Name: "nested structs",
			Input: Person{
				Name: "Foobarius",
				Account: BankAccount{
					Amount:   10,
					Currency: "BTC",
				},
			},
			ExpectedCalls: []string{"Foobarius", "BTC"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				"Foobarius",
				BankAccount{10, "USD"},
			},
			ExpectedCalls: []string{"Foobarius", "USD"},
		},
		{
			Name: "slices",
			Input: []Person{
				{
					Name: "Chase Bank",
					Account: BankAccount{
						Amount:   500,
						Currency: "USD",
					},
				},
			},
			ExpectedCalls: []string{"Chase Bank", "USD"},
		},
		{
			Name: "arrays",
			Input: [2]BankAccount{
				{10, "USD"},
				{10, "GBP"},
			},
			ExpectedCalls: []string{"USD", "GBP"},
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
	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Tap": "Tar",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertValContained(t, "Bar", got)
		assertValContained(t, "Tar", got)
	})
}

func assertValContained(t testing.TB, val string, slice []string) {
	t.Helper()
	contained := false
	for _, item := range slice {
		if val == item {
			contained = true
			break
		}
	}
	if !contained {
		t.Errorf("expexted %v to contain %q but it did not", slice, val)
	}
}
