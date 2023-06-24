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
			var got testValueCatcher
			walk(testCase.Input, returnAppenderFunc(&got))
			assertWithDeepEqual(t, got.values, testCase.ExpectedCalls)
		})
	}
	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Tap": "Tar",
		}
		var got testValueCatcher
		walk(aMap, returnAppenderFunc(&got))
		assertValContained(t, "Bar", got.values)
		assertValContained(t, "Tar", got.values)
	})
	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan BankAccount)
		go func() {
			aChannel <- BankAccount{10, "USD"}
			aChannel <- BankAccount{10, "BTC"}
			close(aChannel)
		}()
		var got testValueCatcher
		want := []string{"USD", "BTC"}
		walk(aChannel, returnAppenderFunc(&got))
		assertWithDeepEqual(t, got.values, want)
	})
	t.Run("functions", func(t *testing.T) {
		aFunc := func() []BankAccount {
			return []BankAccount{
				{10, "USD"},
				{20, "EUR"},
			}
		}
		var got testValueCatcher
		want := []string{"USD", "EUR"}
		walk(aFunc, returnAppenderFunc(&got))
		assertWithDeepEqual(t, got.values, want)
	})
}

type testValueCatcher struct{ values []string }

func returnAppenderFunc(slice *testValueCatcher) func(string) {
	return func(input string) {
		slice.values = append(slice.values, input)
	}
}

func assertWithDeepEqual(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
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
