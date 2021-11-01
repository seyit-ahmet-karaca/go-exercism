package beer

import (
	"errors"
	"fmt"
)

const verseOther = "%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n"
const verse_0 = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
const verse_1 = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
const verse_2 = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"

func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}

func Verses(start, stop int) (string, error) {
	var verses string
	for i := start; i >= stop; i-- {
		if verse, err := Verse(i); err != nil {
			return "", err
		} else {
			verses += verse + "\n"
		}
	}
	if verses == "" {
		return "", fmt.Errorf("Empty verses from %d to %d", start, stop)
	}
	return verses, nil
}

func Verse(n int) (string, error) {
	if n > 2 && n < 100 {
		return fmt.Sprintf(verseOther, n, n, n-1), nil
	} else if n == 0 {
		return verse_0, nil
	} else if n == 1 {
		return verse_1, nil
	} else if n == 2 {
		return verse_2, nil
	}
	return "", errors.New("N should not be negative")
}
