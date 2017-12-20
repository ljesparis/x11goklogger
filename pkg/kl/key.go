package kl

var keyMap = map[string]string{
	"space":        " ",
	"period":       ".",
	"slash":        "/",
	"exclam":       "!",
	"at":           "@",
	"numbersign":   "#",
	"dollar":       "$",
	"percent":      "%",
	"asciicircum":  "^",
	"ampersand":    "&",
	"asterisk":     "*",
	"parenleft":    "(",
	"parenright":   ")",
	"minus":        "-",
	"underscore":   "_",
	"equal":        "=",
	"plus":         "+",
	"bracketleft":  "[",
	"bracketright": "]",
	"braceleft":    "{",
	"braceright":   "}",
	"semicolon":    ";",
	"colon":        ":",
	"apostrophe":   "'",
	"quotedbl":     "\"",
	"grave":        "`",
	"asciitilde":   "~",
	"backslash":    "\\",
	"bar":          "|",
	"comma":        ",",
	"less":         "<",
	"greater":      ">",
	"question":     "?",
	"KP_Multiply":  "*",
	"KP_Subtract":  "-",
	"KP_Add":       "+",
	"KP_Decimal":   ".",
	"KP_Divide":    "/",
}

func getKey(s string) (ret string) {
	if len(s) == 1 {
		if (s[0] > 47 && s[0] < 58) || (s[0] > 64 && s[0] < 91) || (s[0] > 96 && s[0] < 124) {
			ret = s
		} else {
			ret = "\n"
		}
	} else {
		n := keyMap[s]
		if n != "" {
			ret = n
		} else {
			ret = "\n"
		}
	}

	return
}

func isCapsLock(s string) bool {
	return s == "Caps_Lock"
}

func isSpace(s string) bool {
	return s == "space"
}

func isShift(s string) bool {
	return (s == "Shift_L") || (s == "Shift_R")
}
