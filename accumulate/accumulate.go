package accumulate

func Accumulate(list []string, transform func(string) string) []string {
	var newList []string

	for _,v := range list{
		newList = append(newList, transform(v))
	}
	return newList
}
