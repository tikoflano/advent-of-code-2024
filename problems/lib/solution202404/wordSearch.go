package solution202404

import "errors"

type WordSearch struct {
	input []string
}

func NewWordSearch(input []string) *WordSearch {
	return &WordSearch{
		input,
	}
}

func (wordSearch *WordSearch) GetLetterAt(pos *Vector) (string, error) {
	if pos.X < 0 || pos.Y < 0 {
		return "", errors.New("invalid pos")
	}

	if pos.Y >= len(wordSearch.input) {
		return "", errors.New("invalid pos")
	}

	if pos.X >= len(wordSearch.input[pos.Y]) {
		return "", errors.New("invalid pos")
	}

	return string(wordSearch.input[pos.Y][pos.X]), nil
}

func (wordSearch *WordSearch) CheckWordFromPositionAndDirection(word string, pos *Vector, dir *Vector) bool {
	for i, letter := range word {
		curPos := pos.Add(dir.Scale(i))
		wordLetter := string(letter)
		mapLetter, err := wordSearch.GetLetterAt(curPos)

		if err != nil {
			return false
		}

		if mapLetter != wordLetter {
			return false
		}
	}

	return true
}
