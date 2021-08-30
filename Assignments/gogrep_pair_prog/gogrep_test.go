package gogrep2

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var fileContentData = []string{
	"												      ",
	"a.txt											      ",
	"-----------------------------------------------------",
	"|Lorem Ipsum is simply dummy text of the printing   |",
	"| and typesetting industry. Lorem Ipsum has been    |",
	"|the industry's standard dummy text ever since the  |",
	"| 1500s, when an unknown printer took a galley of   |",
	"|type and scrambled it to make a type specimen book.|",
	"| It has survived not only five centuries, but also |",
	"|the leap into electronic typesetting, remaining    |",
	"|essentially unchanged.							 |",
	"-----------------------------------------------------",
	"													    ",
	"b.txt												    ",
	"-------------------------------------------------------",
	"|It is a long established fact that a reader will be  |",
	"|distracted by the readable content of a page when    |",
	"|looking at its layout. The point of using Lorem Ipsum|",
	"| is that it has a more-or-less normal distribution of|",
	"|letters, as opposed to using 'Content here, content  |",
	"|here', making it look like readable English.         |",
	"-------------------------------------------------------",
	"                                                        ",
	"c.txt                                                   ",
	"--------------------------------------------------------",
	"|There are many variations of passages of Lorem Ipsum  |",
	"|available, but the majority have suffered alteration  |",
	"|in some form, by injected humour, or randomised words |",
	"| which don't look even slightly believable. If you are|",
	"| going to use a passage of Lorem Ipsum, you need to be|",
	"| sure there isn't anything embarrassing hidden in the |",
	"|middle of text.                                       |",
	"-------------------------------------------------------|",
}

var testCases = []struct {
	description string
	pattern     string
	files       []string
	expected    string
}{
	{ //Add an array of test cases. With flags, covering all possibilities.
		description: "1 file, 1 match",
		pattern:     "electronic",
		files:       []string{"a.txt"},
		expected:    "the leap into electronic typesetting, remaining",
	},
	{
		description: "1 file, 1 match",
		pattern:     "printing",
		files:       []string{"a.txt"},
		expected:    "Lorem Ipsum is simply dummy text of the printing",
	},
}

func createFiles(content []string) (filenames []string) {
	// Parse fileContentData, making the list of filenames
	// with their content.
	var filename string
	var f *os.File
	for _, d := range content {
		t := strings.TrimSpace(d)
		if len(t) == 0 {
			if len(filename) == 0 {
				continue
			}
			// Close file
			f.Close()
			filenames = append(filenames, filename)
			filename = ""
			f = nil
			continue
		}
		if strings.Contains(t, ".txt") {
			filename = t
			// Open file
			var err error
			f, err = os.Create(filename)
			if err != nil {
				panic(err)
			}
			continue
		}
		fields := strings.Split(t, "|")
		if len(fields) == 3 {
			// Write string into file with newline.
			_, err := f.WriteString(strings.TrimRight(fields[1], " ") + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	if f != nil {
		f.Close()
		filenames = append(filenames, filename)
	}
	return
}

func deleteFiles(filenames []string) {
	for _, file := range filenames {
		os.Remove(file)
	}
}

func TestSearchString(t *testing.T) {
	files := createFiles(fileContentData)
	defer deleteFiles(files)

	for _, tc := range testCases {
		actual := SearchString(tc.files[0], tc.pattern)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("FAIL: %s\nSearch for pattern %q\nexpected %v\nactual %v.",
				tc.description, tc.pattern, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
