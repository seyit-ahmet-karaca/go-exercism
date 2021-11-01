package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

var plants = map[rune]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	g := Garden{}
	if diagram[0] != '\n' {
		return nil, errors.New("wrong format")
	}
	diagram = strings.ReplaceAll(diagram, "\n", "")
	diagramRows := []string{diagram[:len(diagram)/2], diagram[len(diagram)/2:]}
	var copyChildren = make([]string, len(children))
	copy(copyChildren, children)
	sort.Strings(copyChildren)

	if len(diagramRows) != 2 || len(diagramRows[0])/2 != len(children) || len(diagramRows[0]) != len(diagramRows[1]) ||
		len(diagramRows[0])%2 != 0 {
		return nil, errors.New("mismatched rows")
	}

	for i, child := range copyChildren[1:] {
		if copyChildren[i] == child {
			return nil, errors.New("Child is listed twice ")
		}
	}

	for _, row := range diagramRows {
		for i, seed := range row {
			seedStr, ok := plants[seed]
			if !ok {
				return nil, errors.New("invaid cup codes")
			}
			g[copyChildren[i/2]] = append(g[copyChildren[i/2]], seedStr)
		}
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	garden := *g
	plants, ok := garden[child]
	return plants, ok
}
