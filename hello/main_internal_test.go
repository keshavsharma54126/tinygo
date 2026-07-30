package main

import (
	"fmt"
	"testing"
)

// func Example_main(){
// 	main()
// 	// Output:
// 	// hello world
// }

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello world"

	got1 := greet(lang)
	got2 := phrasebookgreet(lang)

	if got1 != want && got2 != want {
		t.Errorf("expected: %q,got1: %q,got2:%q", want, got1, got2)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	got1 := greet(lang)
	got2 := phrasebookgreet(lang)

	if got1 != want && got2 != want {
		t.Errorf("expected: %q,got1: %q,got2:%q", want, got1, got2)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := language("akk")
	want := ""

	got1 := greet(lang)
	got2 := phrasebookgreet(lang)

	if got1 != want && got2 != want {
		t.Errorf("expected: %q,got1: %q,got2:%q", want, got1, got2)
	}
}

func TestGreet_Hindi(t *testing.T) {
	lang := language("hi")
	want := "नमस्ते दुनिया"

	got1 := greet(lang)
	got2 := phrasebookgreet(lang)

	if got1 != want && got2 != want {
		t.Errorf("expected:%s,got1:%s, got2:%s", want, got1, got2)
	}
}

// a more robust way to write the aboove and support all languages is below

func TestGreet(t *testing.T) {
	type testcase struct {
		lang language
		want string
	}

	tests := map[string]testcase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language:"akk"`,
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Empty": {
			lang: "",
			want: `unsupported language:""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fmt.Printf("running test %s,\n", name)
			got := phrasebookgreet(tc.lang)
			if got != tc.want {
				t.Errorf("expected :%q,got: %q", tc.want, got)
			}
		})
	}
}
