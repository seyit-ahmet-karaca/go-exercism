package twelve

import (
	"fmt"
	"strings"
)

const beginOfVerse = "On the %s day of Christmas my true love gave to me:%s"

var dayVerse = map[int]string{
	1:"a Partridge in a Pear Tree.",
	2:"two Turtle Doves, and",
	3:"three French Hens,",
	4:"four Calling Birds,",
	5:"five Gold Rings,",
	6:"six Geese-a-Laying,",
	7:"seven Swans-a-Swimming,",
	8:"eight Maids-a-Milking,",
	9:"nine Ladies Dancing,",
	10:"ten Lords-a-Leaping,",
	11:"eleven Pipers Piping,",
	12:"twelve Drummers Drumming,",
}

var dayString = map[int]string{
	1:"first",
	2:"second",
	3:"third",
	4:"fourth",
	5:"fifth",
	6:"sixth",
	7:"seventh",
	8:"eighth",
	9:"ninth",
	10:"tenth",
	11:"eleventh",
	12:"twelfth",
}

func Song() string {
	var lines []string
	for i := 1 ; i <=12 ; i++ {
		lines = append(lines, Verse(i))
	}
	joined := strings.Join(lines, "\n")
	return joined
}

func Verse(verse int) string {
	var xx string
	for i:= verse ; i > 0 ; i--{
		xx = xx +" "+  dayVerse[i]
	}
	line := fmt.Sprintf(beginOfVerse,dayString[verse], xx)
	fmt.Println(verse," ", line)
	return line
}