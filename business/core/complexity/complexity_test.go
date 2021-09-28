package complexity

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestCountLinesIndentations(t *testing.T) {
	c := Core{}

	tests := []struct {
		name string
		line string
		want int
	}{
		{name: "empty string", line: "", want: 0,},
		{name: "one line with no indentation", line: "some string", want: 0,},
		{name: "one line with one space indentation", line: "  some string", want: 1,},
		{name: "one line with two spaces indentations", line: "    some string", want: 2,},
		{name: "one line with one tab indentation", line: "\tsome string", want: 1,},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := c.CountLineIndentations(test.line, 2)
			if got != test.want {
				t.Errorf("Want %d, Got %d", test.want, got)
			}
		})
	}
}

func TestComplexityAnalysis(t *testing.T) {
	file, err := ioutil.TempFile("", "gocan")
	if err != nil {
		t.Log(err)
		t.Fatalf("Cannot create temp file")
	}

	const s = `
  line 1
    line 2
    line 3
  line 4
`
	file.WriteString(s)
	file.Close()

	c := Core{}
	now := time.Now()
	got, err := c.AnalyzeComplexity(file.Name(), now)
	if err != nil {
		t.Log(err)
		t.Fatalf("Cannot analyze complexity")
	}

	want := Complexity{
		Indentations: 6,
		Date: now,
	}

	if got != want {
		t.Errorf("Want %v, Got %v", want, got)
	}
}

