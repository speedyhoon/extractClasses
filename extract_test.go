package extractClasses

import "testing"

func TestExtract(t *testing.T) {
	for rule, expected := range testData {
		output := Extract(rule)
		if len(output) != len(expected) {
			t.Errorf("Extract(%q)\nslice length %v != %v\n%q != \n%q", rule, len(output), len(expected), expected, output)
		} else {
			for i, s := range expected {
				if s != output[i] {
					t.Errorf("Extract(%q)\n%q != \n%q", rule, expected, output)
				}
			}
		}
	}
}

var testData = map[string][]string{
	//"CSS input string":				{expected slice of strings},
	"":                               {},
	".class-name":                    {".class-name"},
	".class-name, .class-name-other": {".class-name", ".class-name-other"},
	".class-name,.class-name-other":  {".class-name", ".class-name-other"},
	",.first-class.second-class,":    {".first-class", ".second-class"},
	"div.first-class.second-class":   {".first-class", ".second-class"},

	//class within tag
	"div .class-name":                             {".class-name"},
	"div .class-name ":                            {".class-name"},
	"div       .class-name        ":               {".class-name"},
	"div       .first-class.second-class        ": {".first-class", ".second-class"},

	//class within tag's child tag
	"div .class-name a":               {".class-name"},
	"div .first-class.second-class a": {".first-class", ".second-class"},

	//sandwiched class
	"div~!@$%^&*()+=,/';:\"?><[]{}|`.class-name~!@$%^&*()+=,/';:\"?><[]{}|`#": {".class-name"},

	//exclamation mark
	"div .class-name!a":            {".class-name"},
	"div.class-name!a":             {".class-name"},
	".class-name!a":                {".class-name"},
	"!.class-name!a":               {".class-name"},
	"!.first-class.second-class!a": {".first-class", ".second-class"},

	//ampersand
	"div .class-name&a":            {".class-name"},
	"div.class-name&a":             {".class-name"},
	".class-name&a":                {".class-name"},
	"&.class-name&a":               {".class-name"},
	"&.first-class.second-class&a": {".first-class", ".second-class"},

	//dollar
	"div .class-name$a":                          {".class-name"},
	"div.class-name$a":                           {".class-name"},
	".class-name$a":                              {".class-name"},
	"$.class-name$a":                             {".class-name"},
	"a[title~=name] .class-name$a":               {".class-name"},
	"a[title~=name] .first-class.second-class$a": {".first-class", ".second-class"},

	//percentage
	"div .class-name%a":                          {".class-name"},
	"div.class-name%a":                           {".class-name"},
	".class-name%a":                              {".class-name"},
	"%.class-name%a":                             {".class-name"},
	"[%~class-name] .class-name%a":               {".class-name"},
	"[%~class-name] .first-class.second-class%a": {".first-class", ".second-class"},

	//circumflex
	`a.class-name[href^="https"]`:               {".class-name"},
	`a.first-class.second-class[href^="https"]`: {".first-class", ".second-class"},

	//ampersand
	".class-name &":               {".class-name"},
	".first-class.second-class &": {".first-class", ".second-class"},

	//asterisk
	".class-name *":               {".class-name"},
	"*.class-name *":              {".class-name"},
	"*.first-class.second-class*": {".first-class", ".second-class"},

	//brackets
	"p.class-name:lang(it)":                             {".class-name"},
	"p.class-name:lang(it) p.class-name-other:lang(en)": {".class-name", ".class-name-other"},

	//plus
	"div.class-name + p":          {".class-name"},
	"div.class-name+p":            {".class-name"},
	"+.first-class.second-class+": {".first-class", ".second-class"},

	//equals
	`a.class-name[href*="npmjs"]`:    {".class-name"},
	`a.class-name [href *= "npmjs"]`: {".class-name"},
	"=.first-class.second-class=":    {".first-class", ".second-class"},

	//forward slash "/"
	".class-name/class-name-other": {".class-name"},
	"/.first-class.second-class/":  {".first-class", ".second-class"},

	//apostrophe
	".class-name'":                {".class-name"},
	"'.class-name":                {".class-name"},
	"'.first-class.second-class'": {".first-class", ".second-class"},

	//semicolon
	".class-name-1;.class-name-2;":  {".class-name-1", ".class-name-2"},
	";.class-name-1;.class-name-2;": {".class-name-1", ".class-name-2"},
	";.first-class.second-class;":   {".first-class", ".second-class"},

	//colon
	"input.class-name:read-only":                                              {".class-name"},
	"input:out-of-range .class-name input:out-of-range":                       {".class-name"},
	"input:out-of-range .class-name::selection input:out-of-range::selection": {".class-name"},
	":.first-class.second-class:":                                             {".first-class", ".second-class"},

	//double quote
	`.class-name a[href^="https"]`:                  {".class-name"},
	`a[href^="https"] .class-name a[href^="https"]`: {".class-name"},
	`"https".class-name"https"`:                     {".class-name"},
	`"https".first-class.second-class"https"`:       {".first-class", ".second-class"},

	//question mark
	".class-name ?":               {".class-name"},
	"? .class-name?":              {".class-name"},
	"?.class-name?":               {".class-name"},
	"?.first-class.second-class?": {".first-class", ".second-class"},

	//greater than sign
	".class-name> p":                                           {".class-name"},
	"* > .class-name > p > .class-name-other":                  {".class-name", ".class-name-other"},
	"*.class-name> .class-name-other> p > .class-name-another": {".class-name", ".class-name-other", ".class-name-another"},
	">.class1.class2> .class3.class4> p > .class5.class6":      {".class1", ".class2", ".class3", ".class4", ".class5", ".class6"},

	//square brackets
	"a[target=_blank] .class-name a[target=_blank]":            {".class-name"},
	"a[target=_blank] .class-name[target=_blank]":              {".class-name"},
	"a[target=_blank].class-name[target=_blank]":               {".class-name"},
	"a[target=_blank].first-class.second-class[target=_blank]": {".first-class", ".second-class"},

	//curly brackets
	"a{target=_blank} .class-name a{target=_blank}":            {".class-name"},
	"a{target=_blank} .class-name{target=_blank}":              {".class-name"},
	"a{target=_blank}.class-name{target=_blank}":               {".class-name"},
	"a{target=_blank}.first-class.second-class{target=_blank}": {".first-class", ".second-class"},

	//pipe
	"|.class-name|=en]":                 {".class-name"},
	"a[lang|=en] .class-name[lang|=en]": {".class-name"},
	"|.class-name|":                     {".class-name"},
	"|.first-class.second-class|":       {".first-class", ".second-class"},

	//grave accent, tick
	"`.class-name`":               {".class-name"},
	"`.first-class.second-class`": {".first-class", ".second-class"},

	//one letter class names
	".h":     {".h"},
	".a.b.c": {".a", ".b", ".c"},

	//id's
	"#id-name": {"#id-name"},

	//tag with id
	"div#id-name#whatever":      {"#id-name", "#whatever"},
	"div#id-name.class.another": {"#id-name", ".class", ".another"},

	//id within tag
	"div #id-name":                          {"#id-name"},
	"div #id-name ":                         {"#id-name"},
	"div       #id-name        ":            {"#id-name"},
	"div       #first-id#second-id        ": {"#first-id", "#second-id"},

	//id within tag's child tag
	"div #id-name a":              {"#id-name"},
	"div #id-name#second#third a": {"#id-name", "#second", "#third"},
	"div #id-name.second.third a": {"#id-name", ".second", ".third"},

	//id sandwiched
	"~!@$%^&*()+=,/';:\"?><[]{}|`#id-name#second#third[]yo~!@$%^&*()+=,/';:\"?><[]{}|`": {"#id-name", "#second", "#third"},

	//id exclamation mark
	"div #id-name!a":                 {"#id-name"},
	"!#id-name!":                     {"#id-name"},
	"!#id-name#second#third!":        {"#id-name", "#second", "#third"},
	"!#id-name.second#third.fourth!": {"#id-name", ".second", "#third", ".fourth"},

	//id ampersand
	"div #id-name&a":           {"#id-name"},
	"div#id-name&a":            {"#id-name"},
	"#id-name&a":               {"#id-name"},
	"&#id-name&a":              {"#id-name"},
	"&#id-name#second.third&a": {"#id-name", "#second", ".third"},
	"#id-name &":               {"#id-name"},
	"&#id-name&":               {"#id-name"},
	"&#id-name#second&":        {"#id-name", "#second"},

	//id dollar
	"div #id-name$a":            {"#id-name"},
	"div#id-name$a":             {"#id-name"},
	"#id-name$a":                {"#id-name"},
	"$#id-name$a":               {"#id-name"},
	"a[title~=name] #id-name$a": {"#id-name"},
	"$#id-name$":                {"#id-name"},
	"$#id-name#second$":         {"#id-name", "#second"},

	//percentage
	"div #id-name%a":            {"#id-name"},
	"div#id-name%a":             {"#id-name"},
	"#id-name%a":                {"#id-name"},
	"%#id-name%a":               {"#id-name"},
	"[%~class-name] #id-name%a": {"#id-name"},
	"%#id-name%":                {"#id-name"},
	"%#id-name#second%":         {"#id-name", "#second"},

	//circumflex
	`a#id-name[href^="https"]`: {"#id-name"},
	"^#id-name^":               {"#id-name"},
	"^#id-name#second^":        {"#id-name", "#second"},

	//asterisk
	"#id-name *":        {"#id-name"},
	"*#id-name *":       {"#id-name"},
	"*#id-name*":        {"#id-name"},
	"*#id-name#second*": {"#id-name", "#second"},

	//brackets
	"p#id-name:lang(it)":                          {"#id-name"},
	"p#id-name:lang(it) p#id-name-other:lang(en)": {"#id-name", "#id-name-other"},
	"()#id-name()":                                {"#id-name"},
	"(#id-name)":                                  {"#id-name"},
	"(#id-name#second.class)":                     {"#id-name", "#second", ".class"},

	//plus
	"div#id-name + p":   {"#id-name"},
	"div#id-name+p":     {"#id-name"},
	"+#id-name+":        {"#id-name"},
	"+#id-name#second+": {"#id-name", "#second"},

	//equals
	`a#id-name[href*="npmjs"]`:    {"#id-name"},
	`a#id-name [href *= "npmjs"]`: {"#id-name"},
	"=#id-name#second=":           {"#id-name", "#second"},

	//comma
	"#id-name, #id-name-other": {"#id-name", "#id-name-other"},
	"#id-name,#id-name-other":  {"#id-name", "#id-name-other"},
	",#id-name,":               {"#id-name"},
	",#id-name#second,":        {"#id-name", "#second"},

	//forward slash "/"
	"#id-name/#id-name-other":  {"#id-name", "#id-name-other"},
	"/#id-name/#id-name-other": {"#id-name", "#id-name-other"},
	"/#id-name/":               {"#id-name"},
	"/#id-name#second/":        {"#id-name", "#second"},

	//apostrophe
	"#id-name'":        {"#id-name"},
	"'#id-name":        {"#id-name"},
	"'#id-name#second": {"#id-name", "#second"},

	//semicolon
	"#id1;#id2":                 {"#id1", "#id2"},
	"#id-name;#id-name-other":   {"#id-name", "#id-name-other"},
	";#id-name;#id-name-other;": {"#id-name", "#id-name-other"},
	";#id1#id2;#id3#id4;":       {"#id1", "#id2", "#id3", "#id4"},

	//colon
	"input#id-name:read-only":                                                           {"#id-name"},
	"input:out-of-range #id-name input:out-of-range":                                    {"#id-name"},
	"input:out-of-range #id-name::selection input:out-of-range::selection":              {"#id-name"},
	"input:out-of-range #id-name#second.third::selection input:out-of-range::selection": {"#id-name", "#second", ".third"},

	//double quote
	`#id-name a[href^="https"]`:                         {"#id-name"},
	`a[href^="https"] #id-name a[href^="https"]`:        {"#id-name"},
	`a[href^="https"] #id-name#second a[href^="https"]`: {"#id-name", "#second"},

	//question mark
	"#id-name ?":                              {"#id-name"},
	"?#id-name?":                              {"#id-name"},
	"?#id-name#second?":                       {"#id-name", "#second"},
	"?#id-name? > p > #id-name-other":         {"#id-name", "#id-name-other"},
	"?#id-name-1? #id-name-2> p > #id-name-3": {"#id-name-1", "#id-name-2", "#id-name-3"},
	"?#id1#id2? #id3#id4> p > #id5#id6":       {"#id1", "#id2", "#id3", "#id4", "#id5", "#id6"},

	//square brackets
	"a[target=_blank] #id-name a[target=_blank]": {"#id-name"},
	"a[target=_blank] #id-name[target=_blank]":   {"#id-name"},
	"[zzz]#id-name#second[target=_blank]":        {"#id-name", "#second"},
	"zzz[#id-name#second]zzz":                    {"#id-name", "#second"},

	//curly brackets
	"a{target=_blank} #id-name a{target=_blank}": {"#id-name"},
	"a{target=_blank} #id-name{target=_blank}":   {"#id-name"},
	"aaa{bbb}#id-name#second{ccc}ddd":            {"#id-name", "#second"},
	"{#id-name#second}":                          {"#id-name", "#second"},
	"zz{#id-name#second}zzz":                     {"#id-name", "#second"},

	//pipe
	"|#id-name|=en]":                 {"#id-name"},
	"a[lang|=en] #id-name[lang|=en]": {"#id-name"},
	"|#id-name#second|":              {"#id-name", "#second"},

	//tick
	"`#id-name`":        {"#id-name"},
	"`#id-name#second`": {"#id-name", "#second"},

	//classes separated with a space should be recognised
	"div.first-class .second-class":    {".first-class", ".second-class"},
	"div.first-class div.second-class": {".first-class", ".second-class"},
	".first-class .second-class":       {".first-class", ".second-class"},

	//classes recognised after brackets
	"div.class1[lang|=en]#id1[something] .class2[lang|=en] #id2": {".class1", "#id1", ".class2", "#id2"},
	"div.first-class[lang|=en] div.second-class[lang|=en]":       {".first-class", ".second-class"},
	".first-class[lang|=en] .second-class[lang|=en]":             {".first-class", ".second-class"},

	//encoded line breaks
	"#unused-1\n\n\n\n\t\t\t\t\nz\t\ta":                   {"#unused-1"},
	"\n#id\ndiv.first-class\n.second-class\n.third-class": {"#id", ".first-class", ".second-class", ".third-class"},

	//recognises JS escaped strings & repeated dots & hashes
	"\naaa\n...    .unused-1\n\n\n.unused-2, .unused-3\n\t\t,,,\t###\t\nz\t\ta": {".unused-1", ".unused-2", ".unused-3"},
}
