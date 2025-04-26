package main

import (
	"errors"
	"fmt"
	"gopkg.in/xmlpath.v2"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Reference struct {
	book    int
	chapter int
	verse   int
}

func validateQuery(query string) bool {
	pattern := `^(0[1-9]|[1-5][0-9]|6[0-6])(0[0-9]{2}|0[1-9][0-9]|1[0-4][0-9]|150)(0[0-9]{2}|0[1-9][0-9]|1[0-6][0-9]|17[0-6])$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(query)
}

func validateVersion(version string) bool {
	return strings.Compare(version, "kjv") == 0
}

func parseQuery(query string) Reference {
	book, _ := strconv.Atoi(query[0:2])
	chapter, _ := strconv.Atoi(query[2:5])
	verse, _ := strconv.Atoi(query[5:8])

	return Reference{book, chapter, verse}
}

func getPassage(book int, chapter int, verse int, version string) (string, error) {
	path := xmlpath.MustCompile(
		fmt.Sprintf(
			"/bible/book[@num='%d']/chapter[@num='%d']/verse[@num='%d']",
			book,
			chapter,
			verse,
		),
	)

	bible, err := os.Open("./bible/kjv.xml")
	if err != nil {
		return "", err
	}

	root, err := xmlpath.Parse(bible)
	if err != nil {
		return "", err
	}

	if value, ok := path.String(root); ok {
		return value, nil
	}

	return "", errors.New("unable to locate text")
}

func decorateReference(ref Reference) string {
	var BookNumberToName = map[int]string{
		1:  "Genesis",
		2:  "Exodus",
		3:  "Leviticus",
		4:  "Numbers",
		5:  "Deuteronomy",
		6:  "Joshua",
		7:  "Judges",
		8:  "Ruth",
		9:  "1 Samuel",
		10: "2 Samuel",
		11: "1 Kings",
		12: "2 Kings",
		13: "1 Chronicles",
		14: "2 Chronicles",
		15: "Ezra",
		16: "Nehemiah",
		17: "Esther",
		18: "Job",
		19: "Psalms",
		20: "Proverbs",
		21: "Ecclesiastes",
		22: "Song of Solomon",
		23: "Isaiah",
		24: "Jeremiah",
		25: "Lamentations",
		26: "Ezekiel",
		27: "Daniel",
		28: "Hosea",
		29: "Joel",
		30: "Amos",
		31: "Obadiah",
		32: "Jonah",
		33: "Micah",
		34: "Nahum",
		35: "Habakkuk",
		36: "Zephaniah",
		37: "Haggai",
		38: "Zechariah",
		39: "Malachi",
		40: "Matthew",
		41: "Mark",
		42: "Luke",
		43: "John",
		44: "Acts",
		45: "Romans",
		46: "1 Corinthians",
		47: "2 Corinthians",
		48: "Galatians",
		49: "Ephesians",
		50: "Philippians",
		51: "Colossians",
		52: "1 Thessalonians",
		53: "2 Thessalonians",
		54: "1 Timothy",
		55: "2 Timothy",
		56: "Titus",
		57: "Philemon",
		58: "Hebrews",
		59: "James",
		60: "1 Peter",
		61: "2 Peter",
		62: "1 John",
		63: "2 John",
		64: "3 John",
		65: "Jude",
		66: "Revelation",
	}

	return fmt.Sprintf("%s %d:%d", BookNumberToName[ref.book], ref.chapter, ref.verse)
}
