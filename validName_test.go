package extractClasses

import (
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	for _, prefix := range []string{".","#"} {
		for rule, expected := range testRegex {
			if expected != ValidName.MatchString(prefix+rule){
				t.Errorf("rule `%s` expected %v got %v", prefix+rule, expected, !expected)
				return
			}
		}
	}
}

func TestRegexOneChar(t *testing.T){
	const space, del = 32, 127
	var expected bool
	var i uint8
	var rule string

	const symbols = "`~!@#$%^&*()=+[]\\{}|;':\",./<>?"
	const numbers = "0123456789"

	for _, prefix := range []string{".","#"} {
		for i = 0; i <= 254; i++ {
			rule = string(i)
			expected = i > space && i < del && !strings.ContainsAny(rule, symbols) && !strings.ContainsAny(rule, numbers)
			if expected != ValidName.MatchString(prefix+rule){
				t.Errorf("rule `%s` expected %v got %v", prefix+rule, expected, !expected)
				return
			}
		}
	}
}

func TestRegexTwoChars(t *testing.T){
	const space, del, lineFeed, forwardSlash, colon = 32, 127, 10, 92, 58
	var expected bool
	var i,j uint8
	var rule string

	const symbols = "`~!@#$%^&*()=+[]\\{}|;':\",./<>?"
	const numbers = "0123456789"

	for _, prefix := range []string{".","#"} {
		for i = 0; i <= 254; i++ {
			for j = 0; j <= 254; j++ {
				rule = string(i) + string(j)
				if i == colon || j == colon || i == lineFeed || j == lineFeed {
					expected = false
				}else if i == forwardSlash && strings.ContainsAny(string(j), symbols) {
					expected = i > space && i < del && j > space && j < del && !strings.ContainsAny(string(i), numbers)
				}else{
					expected = i > space && i < del && j > space && j < del && !strings.ContainsAny(rule, symbols) && !strings.ContainsAny(string(i), numbers)
				}

				if expected != ValidName.MatchString(prefix+rule) {
					t.Errorf("rule `%s` expected %v got %v\n[%d %d]", prefix+rule, expected, !expected, i, j)
				}
			}
		}
	}
}

var testRegex = map[string]bool{
	"":            false,
	"a": true,
	"A": true,
	"R\a": false,
	"R\b": false,
	"R\f": false,
	"R\n": false,
	"R\r": false,
	"R\t": false,
	"R\v": false,
	"abc": true,
	"ABC": true,
	"aBc": true,
	"a-c": true,
	"a-C": true,
	"a2c": true,
	"a2C": true,
	"_one": true,
	"_oNE": true,
	"class-name": true,
	"Class-Name": true,
	"a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z": true,
	"--a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z": true,
	"--A-B-C-D-E-F-G-H-I-J-K-L-M-N-O-P-Q-R-S-T-U-V-W-X-Y-Z": true,
	"--a-b-c-d-e-f-g-h-i-j-k-l-m-n-o-p-q-r-s-t-u-v-w-x-y-z-A-B-C-D-E-F-G-H-I-J-K-L-M-N-O-P-Q-R-S-T-U-V-W-X-Y-Z": true,
	"890890": false,
	"89High": false,
	`\38 90890`: true,
	`\38 9High`: true,

	"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~": false,
	"!": false,
	`"`:false,
	"#": false,
	"$": false,
	"%": false,
	"&": false,
	"'": false,
	"(": false,
	")": false,
	"*": false,
	"+": false,
	",": false,
	"-": true,
	".": false,
	"/": false,
	":": false,
	";": false,
	"<": false,
	"=": false,
	">": false,
	"?": false,
	"@": false,
	"[": false,
	"\\": false,
	"]": false,
	"^": false,
	"_": true,
	"`": false,
	"{": false,
	"|": false,
	"}": false,
	"~": false,

	"----": true,
	"___": true,
	"_-_-_-_": true,

	`\!`: true,
	`\"`:true,
	`\#`: true,
	`\$`: true,
	`\%`: true,
	`\&`: true,
	`\'`: true,
	`\(`: true,
	`\)`: true,
	`\*`: true,
	`\+`: true,
	`\,`: true,
	`\.`: true,
	`\/`: true,
	`\:`: false,
	`\;`: true,
	`\<`: true,
	`\=`: true,
	`\>`: true,
	`\?`: true,
	`\@`: true,
	`\[`: true,
	`\\`: true,
	`\]`: true,
	`\^`: true,
	"\\`": true,
	`\{`: true,
	`\|`: true,
	`\}`: true,
	`\~`: true,
	`\3A `:true,
	`\#\3A `:true,

	`##d`: false,
	`#.ff`: false,
	`fds\35 a432432`: false,
	`432432423423`: false,
	`\!"`: false,
	`.\!\"\#\$\%\&\'\(\)\*\+\,\.\/\;\<\=\>\?\@\[\\\]^\{\|\}\~`: false,
	`!"#$%&'()*+,-./:;<=>?@[\]^_{|}~`: false,
	`#$%&'\(\)\*\+,-\./;<=>\?@\[\]\^_{\|}~`: false,
	"fdsa`boo": false,
	":hover": false,
	`f:hover`:false,
	"\\3A hover": true,
	`_fd`:true,
	`re`:true,
	`\37 656467`:true,
	`nfdsafdsa`:true,
	`-fdsafsa`:true,
	`f5f555f5f5`:true,
	`fds-gfdsg`:true,
	`f-r`:true,
	`---`:true,
	`--tre-te-ss-`:true,
	`___fdsa-fsa-__e-`:true,
	`\39 abc43`:true,
	`\39 543`:true,
	`\!\"`:true,
	`q\3A 3`:true,
	`\*\3A \(`:true,
	`\3A`:false,
	`\3A trewtwe`:true,
}
