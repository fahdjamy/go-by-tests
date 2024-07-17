package main

import (
	"reflect"
	"testing"
)

type ProgrammingLanguage struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Since   int
	Company string
}

func TestWalk(t *testing.T) {
	t.Run("walking complex functions", func(t *testing.T) {
		cases := []struct {
			Name     string
			Value    interface{}
			Expected []string
		}{
			{
				Name: "sending in a struct with one string value",
				Value: struct {
					Name string
				}{"go"},
				Expected: []string{"go"},
			},
			{
				Name: "sending in a struct with multiple string values",
				Value: struct {
					Name string
					City string
				}{"go", "goland"},
				Expected: []string{"go", "goland"},
			},
			{
				Name: "sending in a struct with multiple values of different types",
				Value: struct {
					Name string
					Age  int
				}{"go", 14},
				Expected: []string{"go"},
			},
			{
				Name: "sending in a struct with multiple nested struct types",
				Value: ProgrammingLanguage{
					Name: "go",
					Profile: Profile{
						Since:   2018,
						Company: "Google",
					},
				},
				Expected: []string{"go", "Google"},
			},
			{
				Name: "sending in a struct with nested pointer type",
				Value: &ProgrammingLanguage{
					Name: "go",
					Profile: Profile{
						Since:   2018,
						Company: "Google",
					},
				},
				Expected: []string{"go", "Google"},
			},
			{
				Name: "sending in a slice of struct values",
				Value: []Profile{
					{
						Since:   2018,
						Company: "Google",
					},
					{
						Since:   2015,
						Company: "Firefox",
					},
				},
				Expected: []string{"Google", "Firefox"},
			},
			{
				Name: "sending in a array of struct values",
				Value: [2]Profile{
					{
						Since:   2018,
						Company: "Google",
					},
					{
						Since:   2015,
						Company: "Firefox",
					},
				},
				Expected: []string{"Google", "Firefox"},
			},
			{
				Name: "sending in a struct containing a slice of values",
				Value: struct {
					Profiles []Profile
				}{
					[]Profile{
						{
							Since:   2018,
							Company: "Google",
						},
						{
							Since:   2015,
							Company: "Firefox",
						},
					},
				},
				Expected: []string{"Google", "Firefox"},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var got []string
				walk(c.Value, func(s string) {
					got = append(got, s)
				})
				if len(got) != len(c.Expected) {
					t.Errorf("called more than once")
				}
				if !reflect.DeepEqual(got, c.Expected) {
					t.Errorf("got %v, want %v", got, c.Expected)
				}
			})
		}
	})

	t.Run("walk assigned map type", func(t *testing.T) {
		value := map[string]string{"language": "go", "company": "Google"}

		var got []string

		walk(value, func(s string) {
			got = append(got, s)
		})
		assertContains(t, got, "go")
		assertContains(t, got, "Google")
	})

	t.Run("walk with channel type", func(t *testing.T) {
		ch := make(chan Profile)
		go func() {
			defer close(ch)
			ch <- Profile{1999, "Apple"}
			ch <- Profile{2000, "Google"}
		}()

		var got []string
		expected := []string{"Apple", "Google"}
		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})

	t.Run("walk with func type", func(t *testing.T) {
		testFn := func() []Profile {
			profiles := []Profile{
				{1999, "Apple"},
				{2000, "Google"},
			}

			return profiles
		}
		var got []string
		expected := []string{"Apple", "Google"}
		walk(testFn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})
}

func assertContains(t *testing.T, response []string, want string) {
	t.Helper()
	inSlice := false
	for _, v := range response {
		if v == want {
			inSlice = true
		}
	}
	if !inSlice {
		t.Errorf("value not in slice")
	}
}
