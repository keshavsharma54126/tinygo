package main

import (
	"testing"
)

func TestLoadBookworms_Success(t *testing.T) {
	test_cases := map[string]struct {
		bookWormFileName string
		want             []Bookworms
		wantErr          bool
	}{
		"file_exitsts": {
			bookWormFileName: "./testdata/bookworm.json",
			want: []Bookworms{
				{
					Name: "Fadi",
					Books: []Book{
						{
							Author: "Margaret Atwood",
							Title:  "The Handmaid's Tale",
						},
						{
							Author: "Sylvia Plath",
							Title:  "The Bell Jar",
						},
					},
				},
				{
					Name: "Peggy",
					Books: []Book{
						{
							Author: "Margaret Atwood",
							Title:  "Oryx and Crake",
						},
						{
							Author: "Margaret Atwood",
							Title:  "The Handmaid's Tale",
						},
						{
							Author: "Charlotte Brontë",
							Title:  "Jane Eyre",
						},
					},
				},
			},
			wantErr: false,
		},
		"file_not_exist": {
			bookWormFileName: "./bookworms.json",
			want: []Bookworms{
				{
					Name: "Fadi",
					Books: []Book{
						{
							Author: "Margaret Atwood",
							Title:  "The Handmaid's Tale",
						},
					},
				},

				{
					Name: "Peggy",
					Books: []Book{
						{
							Author: "Margaret Atwood",
							Title:  "Oryx and Crake",
						},
					},
				},
			},
			wantErr: true,
		},
		"invalid_json": {
			bookWormFileName: "./testdata/invalid.json",
			want: []Bookworms{
				{
					Name: "Fadil",
					Books: []Book{
						{
							Author: "sadfMargaret Atwood",
							Title:  "The Handmaid's Tale",
						},
					},
				},

				{
					Name: "asdfPeggy",
					Books: []Book{
						{
							Author: "Margaret Atwood",
							Title:  "Oryx and Crake",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for name, testcase := range test_cases {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testcase.bookWormFileName)
			if testcase.wantErr {
				if err == nil {
					t.Fatalf("expected an error got %s", got)
				}
				return
			}
			if !testcase.wantErr && err != nil {
				t.Fatalf("did not expected an error got error %s", err.Error())
			}

			if !isEqualBookworms(t, got, testcase.want) {
				t.Fatalf("expected bookmars to be %s \n but got %s", testcase.want, got)
			}
		})
	}
}

func isEqualBookworms(t *testing.T, got []Bookworms, want []Bookworms) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}
	for index := range got {
		if got[index].Name != want[index].Name {
			return false
		}
		if !isEqualBooks(t, got[index].Books, want[index].Books) {
			return false
		}
	}
	return true
}

func isEqualBooks(t *testing.T, got []Book, want []Book) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}
	for index := range got {
		if got[index].Author != want[index].Author {
			return false
		}
		if got[index].Title != want[index].Title {
			return false
		}
	}
	return true
}
