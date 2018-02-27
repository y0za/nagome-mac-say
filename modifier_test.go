package main

import "testing"

func TestOmitURL(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"normal text", "normal text"},
		{"http://google.com", "URL省略"},
		{"https://google.com", "URL省略"},
		{"text https://google.com text", "text URL省略 text"},
		{"text　https://google.com　text", "text　URL省略　text"},
	}

	for _, tt := range tests {
		isa := SayArgs{
			Text: tt.in,
		}
		osa, err := OmitURL(isa)
		if err != nil {
			t.Errorf("unexpected error occurred when the text is '%s': %s", tt.in, err)
		} else if osa.Text != tt.out {
			t.Errorf("expected text '%s' when the input text is '%s', but actual '%s'", tt.out, tt.in, osa.Text)
		}

	}
}
