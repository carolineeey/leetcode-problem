package solution

import "strings"

func FullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	var line []string
	lineLen := 0

	for i := 0; i < len(words); {
		// try to fit as many words as possible into the line
		line = []string{}
		lineLen = 0

		for i < len(words) && lineLen+len(line)+len(words[i]) <= maxWidth {
			line = append(line, words[i])
			lineLen += len(words[i])
			i++ // only increment here after including the word, after finished building an entire line
		}

		spacesNeeded := maxWidth - lineLen
		justified := ""

		// check if the current line should be left-justified instead of fully justified
		// i == len(words) : we’ve reached the last word
		// len(line) == 1 : the current line only have one word
		if i == len(words) || len(line) == 1 {
			// last line or a line with only one word -> left-justified
			justified = line[0] // start with the first word
			for j := 1; j < len(line); j++ {
				justified += " " + line[j] // add a space before each subsequent word
			}
			justified += strings.Repeat(" ", maxWidth-len(justified)) // pad remaining space at the end
		} else { // fully justified (not the last line)
			spaces := spacesNeeded / (len(line) - 1)      // calculate the minimum number of spaces between each word
			extraSpaces := spacesNeeded % (len(line) - 1) // find how many extra spaces are left after even distribution

			for j := 0; j < len(line)-1; j++ { // loop through all words except the last one as the last word don't need any space
				justified += line[j]                     // add the word to the line
				justified += strings.Repeat(" ", spaces) // add the base number of spaces between words
				if j < extraSpaces {
					justified += " " // for the first extraSpaces words, add one more space to the gap
				}
			}
			justified += line[len(line)-1] // add the last word without any extra space afterward
		}

		res = append(res, justified)
	}

	return res
}
