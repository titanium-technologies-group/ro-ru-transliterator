package transliteration

import (
	"strings"
)

var finishingReplacer *strings.Replacer
var firstPriorityReplacer *strings.Replacer
var vowelLookupTable map[string]bool

/**
Transliterates from romanian to russian
All uppercase letters will be transformed to lower case

 */
func TransliterateInRussian(text string) string {
	words := strings.Split(text, " ")
	var result string
	for _, v := range words {
		result += transliterateWord(v) + " "
	}
	return strings.TrimSpace(result)
}

func transliterateWord(word string) string {
	word = strings.ToLower(word)
	word = firstPriorityReplacer.Replace(word)
	word = replaceSuffix(word, "ci", "ч")
	word = replaceSuffix(word, "ii", "и")
	word = replaceSuffix(word, "iu", "иу")
	word = replaceSuffix(word, "iii", "ий")
	word = replaceSuffix(word, "ia", "ия")
	word = replaceSuffix(word, "ie", "ие")
	word = replacePrefix(word, "î", "и")
	word = replacePrefix(word, "ia", "я")
	word = replacePrefix(word, "ie", "е")
	word = replacePrefix(word, "io", "йо")
	word = replacePrefix(word, "iu", "ю")
	word = replaceInTheMiddleIfBefore(word, "ia", "я", isVowel)
	word = replaceForEandI(word, "c", "ч")
	word = replaceForEandI(word, "g", "дж")
	word = replaceInTheMiddleIfBefore(word, "ie", "е", isVowel)
	word = replaceInTheMiddleIfBefore(word, "i", "й", isVowel)
	word = replaceInTheMiddleIfBefore(word, "ia", "ья", isConsonant)
	word = replaceInTheMiddleIfBefore(word, "ie", "ье", isConsonant)
	word = replaceInTheMiddleIfBefore(word, "io", "йо", isVowel)
	word = replaceInTheMiddleIfBefore(word, "io", "ьо", isConsonant)
	word = replaceInTheMiddleIfBefore(word, "iu", "ю", isVowel)
	word = replaceInTheMiddleIfBefore(word, "iu", "ью", isConsonant)
	return finishingReplacer.Replace(word)
}

func replaceInTheMiddleIfBefore(text string, replacement string, replacementWith string, checker func(previous string) bool) string {
	index := strings.Index(text, replacement)
	if index > 0 && checker(string(text[index-1])) {
		result := strings.Replace(text, replacement, replacementWith, -1)
		return replaceInTheMiddleIfBefore(result, replacement, replacementWith, checker)
	}
	return text
}

func replaceForEandI(text string, replacement string, replacementWith string) string {
	index := strings.Index(text, replacement)
	if index > -1 && index != len(text)-1 && isEorI(string(text[index+1])) {
		result := strings.Replace(text, replacement, replacementWith, -1)
		return replaceForEandI(result, replacement, replacementWith)
	}
	return text
}

func replaceSuffix(text string, suffix string, replacement string) string {
	if strings.HasSuffix(text, suffix) {
		return strings.TrimSuffix(text, suffix) + replacement
	}
	return text
}

func replacePrefix(text string, prefix string, replacement string) string {
	if strings.HasPrefix(text, prefix) {
		return replacement + strings.TrimPrefix(text, prefix)
	}
	return text
}

func init() {
	vowels := [7]string{"a", "e", "i", "o", "u", "î", "ă"}
	vowelLookupTable = make(map[string]bool)
	for _, v := range vowels {
		vowelLookupTable[v] = true
	}
	firstPriorityReplacer = strings.NewReplacer(
		"cea", "ча",
		"cia", "ча",
		"cio", "чо",
		"ciu", "чу",
		"gea", "джа",
		"gia", "джа",
		"geo", "джо",
		"gio", "джо",
		"giu", "джу",
		"ch", "к",
		"gh", "г",
	)
	finishingReplacer = strings.NewReplacer(
		"ea", "я",
		"ș", "ш",
		"ă", "э",
		"â", "ы",
		"îi", "ый",
		"î", "ы",
		"m", "м",
		"p", "п",
		"n", "н",
		"j", "ж",
		"q", "к",
		"k", "к",
		"r", "р",
		"ț", "ц",
		"d", "д",
		"t", "т",
		"i", "и",
		"f", "ф",
		"e", "е",
		"b", "б",
		"u", "у",
		"o", "о",
		"l", "л",
		"c", "к",
		"s", "с",
		"v", "в",
		"w", "в",
		"x", "кс",
		"y", "и",
		"z", "з",
		"a", "а",
		"g", "г",
		"h", "х",
	)
}

func isVowel(x string) bool {
	return vowelLookupTable[x]
}

func isConsonant(x string) bool {
	return !vowelLookupTable[x]
}

func isEorI(x string) bool {
	return x == "e" || x == "i"
}
