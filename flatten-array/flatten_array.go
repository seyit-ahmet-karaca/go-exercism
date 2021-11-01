package flatten

func Flatten(list interface{}) []interface{} {
	switch list.(type) {
	case []interface{}:
		result := []interface{}{}
		for _, val := range list.([]interface{}) {
			flatten := Flatten(val)
			for _, element := range flatten {
				result = append(result, element)
			}
		}
		return result
	case nil:
		return []interface{}{}
	}
	return []interface{}{list}
}