package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var namesUsed = make(map[string]bool)
const maxNameCount = 26 * 26 * 10 * 10 * 10

func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	if len(namesUsed) == maxNameCount {
		return "",errors.New("No More Space  There is no more name could generate!")
	}
	if r.name == "" {
		for used := true; used; used = namesUsed[r.name] {
			r.name = generateID()
		}
		namesUsed[r.name] = true
	}

	return r.name, nil
}

func generateID() string {
	return fmt.Sprintf("%s%s%03d", randomLetter(), randomLetter(), rand.Intn(1000))
}

func randomLetter() string {
	return fmt.Sprintf("%c", rune(rand.Intn(26)+'A'))
}

func (r *Robot) Reset() {
	r.name = ""
}
