package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"hexadecimal et binaire", "Simply add 42 (hex) and 10 (bin)", "Simply add 66 and 2"},
		{"majuscule", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"minuscules", "SHOUTING (low)", "shouting"},
		{"capitalisation", "harold wilson (cap, 2)", "Harold Wilson"},
		{"plusieurs derniers mots", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"espaces de ponctuation", "over there ,and then BAMM !!", "over there, and then BAMM!!"},
		{"points de suspension", "I was thinking ... You were right", "I was thinking... You were right"},
		{"apostrophes", "they said: ' I am ready '", "they said: 'I am ready'"},
		{"article a", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"accent", "élan (cap)", "Élan"},
		{"marqueur sans mot", "(up) hello", "hello"},
		{"exemple complet 1", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"exemple complet 2", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"exemple complet 3", "Don not be sad ,because sad backwards is das . And das not good .", "Don not be sad, because sad backwards is das. And das not good."},
		{"exemple complet 4", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"exemple complet 5", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"exemple complet 6", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Process(test.in); got != test.want {
				t.Errorf("Process(%q) = %q, want %q", test.in, got, test.want)
			}
		})
	}
}

func TestProcessTextKeepsLines(t *testing.T) {
	input := "go (up)\nhello ,world"
	want := "GO\nhello, world"
	if got := processText(input); got != want {
		t.Errorf("processText() = %q, want %q", got, want)
	}
}

func TestToDecimalKeepsInvalidValue(t *testing.T) {
	if got := toDecimal("zz", 16); got != "zz" {
		t.Errorf("toDecimal() = %q, want %q", got, "zz")
	}
}
