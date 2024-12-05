package solution202404

import "errors"

func GetMapLetter(input []string, pos *Vector) (string, error) {
	if pos.X < 0 || pos.Y < 0 {
		return "", errors.New("invalid pos")
	}

	if pos.Y >= len(input) {
		return "", errors.New("invalid pos")
	}

	if pos.X >= len(input[pos.Y]) {
		return "", errors.New("invalid pos")
	}

	return string(input[pos.Y][pos.X]), nil
}

func CheckFromPositionAndDirection(word string, input []string, pos *Vector, dir *Vector) bool {
	for i, letter := range word {
		curPos := pos.Add(dir.Scale(i))
		wordLetter := string(letter)
		mapLetter, err := GetMapLetter(input, curPos)

		if err != nil {
			return false
		}

		if mapLetter != wordLetter {
			return false
		}
	}

	return true
}

func IsMas(word string) bool {
	return word == "MAS" || word == "SAM"
}
