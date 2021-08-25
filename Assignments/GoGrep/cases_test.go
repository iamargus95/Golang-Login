package gogrep

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
	flags       []string
	files       []string
	expected    []string
}{
	{ //Add an array of test cases. With flags, covering all possibilities.
		description: "1 file, 1 match, no flags",
		pattern:     "electronic",
		flags:       []string{},
		files:       []string{"a.txt"},
		expected:    []string{"the leap into electronic typesetting, remaining"},
	},
	{
		description: "1 file, 1 match, print line numbers flag",
		pattern:     "printing",
		flags:       []string{"-n"},
		files:       []string{"a.txt"},
		expected:    []string{"1:Lorem Ipsum is simply dummy text of the printing"},
	},
	{
		description: "1 file, 1 match, case-insensitive flag",
		pattern:     "VARIATIONS",
		flags:       []string{"-i"},
		files:       []string{"c.txt"},
		expected:    []string{"There are many variations of passages of Lorem Ipsum"},
	},
	{
		description: "1 file, 1 match,print filenames flag",
		pattern:     "normal",
		flags:       []string{"-l"},
		files:       []string{"b.txt"},
		expected:    []string{"b.txt"},
	},
	{
		description: "1 file, 1 match, match entire lines flag",
		pattern:     "There are many variations of passages of Lorem Ipsum",
		flags:       []string{"-x"},
		files:       []string{"c.txt"},
		expected:    []string{"There are many variations of passages of Lorem Ipsum"},
	},
}
