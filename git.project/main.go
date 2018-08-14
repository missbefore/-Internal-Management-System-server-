package main

import (
	"fmt"

	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "Your long raw text, it could be a book. Lorem ipsum..."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}