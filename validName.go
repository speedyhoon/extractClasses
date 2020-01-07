package extractClasses

import (
	"regexp"
)

/* symbols:
	\3A		colon :				(must contain a trailing space)
	\3[0-9]	leading digit 0-9	(must contain a trailing space)
	\x60		grave accent `
*/

const anySymbol = `[!\"#$%&'()*+,./;<=>?@[\\\]^\x60{|}~]`	//excluding colon :, hyphen - and underscore _

//ValidName can be used to determine if a CSS class or id is valid or not.
//This regular expression is much more strict than the W3C recommendation //www.w3.org/TR/CSS21/grammar.html#scanner
var ValidName = regexp.MustCompile(
	`^[.#]([a-zA-Z_-]|(\\((3[0-9A] )|`+ anySymbol +`)))([0-9a-zA-Z_-]|(\\((3A )|`+ anySymbol +`)))*$`,
)