package main

import (
	"errors"
	"fmt"
	"gopkg.in/xmlpath.v2"
	"os"
	"regexp"
	"strconv"
	"strings"
	"theopneustos/pkg/constants"
)

type Reference struct {
	book    int
	chapter int
	verse   int
}

var FirstVerse = Reference{constants.GenesisID, 1, 1}
var LastVerse = Reference{constants.RevelationID, 22, 21}

var validQueryRegex = regexp.MustCompile(`^[ACDEFGHIJKLMNOPRSTZacdefghijklmnoprstz123][A-PR-Za-pr-z0-9:. ]{0,31}$`)
var validVerseIdRegex = regexp.MustCompile(`^(0[1-9]|[1-5][0-9]|6[0-6])(0[0-9]{2}|0[1-9][0-9]|1[0-4][0-9]|150)(0[0-9]{2}|0[1-9][0-9]|1[0-6][0-9]|17[0-6])$`)

func parseQuery(query string) Reference {
	if isVerseID(query) {
		book, _ := strconv.Atoi(query[0:2])
		chapter, _ := strconv.Atoi(query[2:5])
		verse, _ := strconv.Atoi(query[5:8])

		return Reference{book, chapter, verse}
	}

	query = strings.Replace(query, " ", "", -1)
	query = strings.Replace(query, ".", "", -1)
	query = strings.ToUpper(query)
	prefix := ""
	switch query[0] {
	case '1':
		prefix = "FIRST"
	case '2':
		prefix = "SECOND"
	case '3':
		prefix = "THIRD"
	}
	if prefix != "" {
		query = prefix + query[1:]
	}

	book := parseTextQueryForBook(query)

	var b strings.Builder
	for _, char := range query {
		if (char >= '0' && char <= '9') || char == ':' {
			b.WriteRune(char)
		}
	}

	chapterAndVerse := b.String()
	colonIndex := strings.Index(chapterAndVerse, ":")
	chapter, _ := strconv.Atoi(chapterAndVerse[:colonIndex])
	verse, _ := strconv.Atoi(chapterAndVerse[colonIndex+1:])

	return Reference{book, chapter, verse}
}

func parseTextQueryForBook(query string) int {

	switch query[0] {
	case 'A':
		if query[1] == 'C' {
			return constants.ActsID
		}
		return constants.AmosID
	case 'C':
		return constants.ColossiansID
	case 'D':
		switch query[1] {
		case 'A', 'N':
			return constants.DanielID
		case 'E', 'T':
			return constants.DeuteronomyID
		}
	case 'E':
		switch query[1] {
		case 'C':
			return constants.EcclesiastesID
		case 'P':
			return constants.EphesiansID
		case 'S':
			return constants.EstherID
		case 'X':
			return constants.ExodusID
		case 'Z':
			if query[2] == 'R' {
				return constants.EzraID
			}
			return constants.EzekielID
		}
	case 'G':
		if query[1] == 'A' {
			return constants.GalatiansID
		}
		return constants.GenesisID
	case 'H':
		switch query[1] {
		case 'A':
			if query[2] == 'B' {
				return constants.HabakkukID
			}
			return constants.HaggaiID
		case 'E':
			return constants.HebrewsID
		case 'O':
			return constants.HoseaID
		}
	case 'I':
		return constants.IsaiahID
	case 'J':
		switch query[1] {
		case 'B':
			return constants.JobID
		case 'L':
			return constants.JoelID
		case 'E', 'R':
			return constants.JeremiahID
		case 'N':
			return constants.JohnID
		case 'G':
			return constants.JudgesID
		}
		if strings.Contains(query, "G") {
			return constants.JudgesID
		}
		if strings.Contains(query, "S") ||
			strings.Contains(query, "JM") ||
			strings.Contains(query, "JAS") {
			return constants.JamesID
		}
		if strings.Contains(query, "D") {
			return constants.JudeID
		}
		if strings.Contains(query, "B") {
			return constants.JobID
		}
		if strings.Contains(query, "L") {
			return constants.JoelID
		}
		if strings.Contains(query, "S") {
			return constants.JoshuaID
		}
		if strings.Contains(query, "JON") ||
			strings.Contains(query, "JNH") {
			return constants.JonahID
		}
		return constants.JohnID
	case 'L':
		switch query[1] {
		case 'A':
			return constants.LamentationsID
		case 'E', 'V':
			return constants.LeviticusID
		}
		return constants.LukeID
	case 'M':
		switch query[1] {
		case 'C', 'I':
			return constants.MicahID
		case 'L':
			return constants.MalachiID
		case 'T':
			return constants.MatthewID
		case 'K', 'R':
			return constants.MarkID
		case 'A':
			switch query[2] {
			case 'L':
				return constants.MalachiID
			case 'R':
				return constants.MarkID
			case 'T':
				return constants.MatthewID
			}
		}
	case 'N':
		switch query[1] {
		case 'A':
			return constants.NahumID
		case 'E':
			return constants.NehemiahID
		}
		return constants.NumbersID
	case 'O':
		return constants.ObadiahID
	case 'P':
		switch query[1] {
		case 'P':
			return constants.PhilippiansID
		case 'M':
			return constants.PhilemonID
		case 'R':
			return constants.ProverbsID
		case 'S':
			return constants.PsalmsID
		}
		if strings.Contains(query, "M") || strings.Contains(query, "PHILE") {
			return constants.PhilemonID
		}
		return constants.PhilippiansID
	case 'R':
		switch query[1] {
		case 'T', 'U':
			return constants.RuthID
		case 'O', 'M':
			return constants.RomansID
		}
		return constants.RevelationID
	case 'S':
		return constants.SongOfSolomonID
	case 'T':
		return constants.TitusID
	case 'Z':
		if strings.Contains(query, "C") {
			return constants.ZechariahID
		} else if strings.Contains(query, "P") {
			return constants.ZephaniahID
		}
	case '1':
		switch query[1] {
		case 'S':
			return constants.FirstSamuelID
		case 'K':
			return constants.FirstKingsID
		case 'P':
			return constants.FirstPeterID
		case 'J':
			return constants.FirstJohnID
		case 'C':
			if query[2] == 'H' {
				return constants.FirstChroniclesID
			}
			return constants.FirstCorinthiansID
		case 'T':
			if query[2] == 'H' {
				return constants.FirstThessaloniansID
			}
			return constants.FirstTimothyID
		}
	case '2':
		switch query[1] {
		case 'S':
			return constants.SecondSamuelID
		case 'K':
			return constants.SecondKingsID
		case 'P':
			return constants.SecondPeterID
		case 'J':
			return constants.SecondJohnID
		case 'C':
			if query[2] == 'H' {
				return constants.SecondChroniclesID
			}
			return constants.SecondCorinthiansID
		case 'T':
			if query[2] == 'H' {
				return constants.SecondThessaloniansID
			}
			return constants.SecondTimothyID
		}
	case '3':
		return constants.ThirdJohnID
	default:
		return 0
	}

	return 0
}

func getPreviousVerse(ref Reference) Reference {
	if ref == FirstVerse {
		return LastVerse
	}

	newRef := Reference{ref.book, ref.chapter, ref.verse - 1}
	if isValidReference(newRef) {
		return newRef
	}

	newRef = Reference{ref.book, ref.chapter - 1, constants.VerseCounts[ref.book][ref.chapter-1]}
	if isValidReference(newRef) {
		return newRef
	}

	return Reference{ref.book - 1, constants.ChapterCount[ref.book-1], constants.VerseCounts[ref.book-1][ref.chapter-1]}
}

func getUrl(ref Reference) string {
	return fmt.Sprintf("https://api.theopneustos.bible/kjv/text?q=%02d%03d%03d", ref.book, ref.chapter, ref.verse)
}

func getNextVerse(ref Reference) Reference {
	if ref == LastVerse {
		return FirstVerse
	}

	newRef := Reference{ref.book, ref.chapter, ref.verse + 1}
	if isValidReference(newRef) {
		return newRef
	}

	newRef = Reference{newRef.book, newRef.chapter + 1, 1}
	if isValidReference(newRef) {
		return newRef
	}

	return Reference{newRef.book + 1, 1, 1}
}

func isValidReference(ref Reference) bool {
	if ref.book < constants.GenesisID || ref.book > constants.RevelationID {
		return false
	}

	chapterCount := constants.ChapterCount[ref.book]
	if ref.chapter == 0 || ref.chapter > chapterCount {
		return false
	}

	if ref.verse == 0 || ref.verse > constants.VerseCounts[ref.book][ref.chapter-1] {
		return false
	}

	return true
}

func isVerseID(str string) bool {
	return validVerseIdRegex.MatchString(str)
}

func isValidVersion(version string) bool {
	switch version {
	case "KJV":
		return true
	default:
		return false
	}
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
	return fmt.Sprintf("%s %d:%d", constants.BookNumberToName[ref.book], ref.chapter, ref.verse)
}
