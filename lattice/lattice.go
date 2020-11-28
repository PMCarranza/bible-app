package lattice

import (
	"strings"
)

type TagType int

const (
	None    TagType = iota
	Book
	Comment
	Chapter
	ChapterVerse
	CharacterNote
	CharacterQuote
	Parenthesis
	SearchResultHiLite
)

type TextTag struct {
	StartIndex int
	EndIndex   int // Optional -- some tag values are in the text already like CharacterQuote
	TagValue   string
	TagType    TagType
}

//export QueryResponse
type QueryResponse struct {
	Body       strings.Builder
	TextTags   []TextTag
}