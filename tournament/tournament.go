package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type teamStatus struct {
	name                            string
	played, won, drawn, lost, point int
}

func Tally(r io.Reader, w io.Writer) error {
	table, _ := ioutil.ReadAll(r)
	rows := strings.Split(string(table), "\n")
	var tableMap = make(map[string]*teamStatus)

	//fmt.Println(string(table))

	for _, row := range rows {
		records := strings.Split(row, ";")
		if len(records) == 1 {
			continue
		} else if len(records) != 3 {
			return errors.New("Error !")
		}
		_, okRec0 := tableMap[records[0]]
		if !okRec0 {
			tableMap[records[0]] = &teamStatus{records[0], 0, 0, 0, 0, 0}
		}
		_, okRec1 := tableMap[records[1]]
		if !okRec1 {
			tableMap[records[1]] = &teamStatus{records[1], 0, 0, 0, 0, 0}
		}
		team1 := tableMap[records[0]]
		team1.played++

		team2 := tableMap[records[1]]
		team2.played++

		if records[2] == "win" {
			team1.point += 3
			team1.won++
			team2.lost++
		} else if records[2] == "loss" {
			team2.point += 3
			team1.lost++
			team2.won++
		} else if records[2] == "draw" {
			team1.point++
			team1.drawn++
			team2.point++
			team2.drawn++
		} else {
			return errors.New("error")
		}
		tableMap[records[0]] = team1
		tableMap[records[1]] = team2
	}
	var sorted []*teamStatus

	for _, v := range tableMap {
		sorted = append(sorted, v)
	}

	for i := 0; i < len(sorted); i++ {
		max := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[max].point < sorted[j].point {
				max = j
			}else if sorted[max].point == sorted[j].point {
				if sorted[max].name > sorted[j].name {
					max = j
				}
			}
		}
		sorted[i], sorted[max] = sorted[max], sorted[i]
	}

	output := "Team                           | MP |  W |  D |  L |  P\n"

	for _, v := range sorted {
		output += fmt.Sprintf(
			"%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.name,
			v.played,
			v.won,
			v.drawn,
			v.lost,
			v.point,
		)
	}
	_, err := w.Write([]byte(output))
	if err != nil {
		return nil
	}
	return err
}
