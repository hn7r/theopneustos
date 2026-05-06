package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type ApiResponse struct {
	Query       string        `json:"query"`
	Canonical   string        `json:"canonical"`
	Parsed      [][]int       `json:"parsed"`
	PassageMeta []PassageMeta `json:"passage_meta"`
	Passages    []string      `json:"passages"`
}

type PassageMeta struct {
	Canonical    string `json:"canonical"`
	ChapterStart []int  `json:"chapter_start"`
	ChapterEnd   []int  `json:"chapter_end"`
	PrevVerse    int    `json:"prev_verse"`
	NextVerse    int    `json:"next_verse"`
	PrevChapter  []int  `json:"prev_chapter"`
	NextChapter  []int  `json:"next_chapter"`
}

func getEsvText(book int, chapter int, verse int) string {

	ref := Reference{book, chapter, verse}

	q := decorateReference(ref)
	q = strings.ReplaceAll(q, " ", "+")

	baseUrl, _ := url.Parse("https://api.esv.org/v3/passage/text")
	queryParams := url.Values{}
	queryParams.Add("q", q)
	queryParams.Add("include-passage-references", "false")
	queryParams.Add("include-verse-numbers", "false")
	queryParams.Add("include-first-verse-numbers", "false")
	queryParams.Add("include-footnotes", "false")
	queryParams.Add("include-footnote-body", "false")
	queryParams.Add("include-headings", "false")
	queryParams.Add("include-short-copyright", "false")
	queryParams.Add("include-selahs", "false")
	queryParams.Add("indent-paragraphs", "0")
	queryParams.Add("indent-poetry", "false")
	queryParams.Add("indent-poetry-lines", "0")
	queryParams.Add("indent-declares", "0")
	queryParams.Add("indent-psalm-doxology", "0")

	baseUrl.RawQuery = queryParams.Encode()

	req, _ := http.NewRequest("GET", baseUrl.String(), nil)
	req.Header.Set("Authorization", "Token <TOKEN>")

	// Using the default client (part of std lib)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	var data ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return ""
	}

	responseText := ""
	for _, passage := range data.Passages {
		passage = strings.ReplaceAll(passage, "\n", "")
		responseText += passage
	}

	return responseText
}
