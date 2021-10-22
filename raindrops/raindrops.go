package raindrops

import "fmt"

func Convert(input int) string {
	drop := ""
	drops := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for k,v := range drops{
		if input % k == 0 {
			drop = drop + v
		}
	}
	if drop == "" {
		return fmt.Sprintf("%d",input)
	}
	return drop
}