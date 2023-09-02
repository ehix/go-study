package hello

import "testing"

func TestSayHello(t *testing.T) {
	// creating structs on the fly
	subtests := []struct {
		items  []string
		result string
	}{
		// test cases, with 0, 1, and 1+ names:
		{
			result: "Hello, world",
		},
		{
			items:  []string{"Alex"},
			result: "Hello, Alex",
		},
		{
			items:  []string{"Alex", "Sue"},
			result: "Hello, Alex, Sue",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

	// want := "Hello, test"
	// got := Say([]string{"test"})

	// if want != got {
	// 	t.Errorf("wanted %s, got %s", want, got)
	// }
}
