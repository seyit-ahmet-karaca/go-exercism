package strain

type Ints []int
type Lists [][]int
type Strings []string




//    (Strings) Keep(func(string) bool) Strings
func (i Strings) Keep(f func(string) bool) Strings {
	var newList = Strings{}
	for _, x := range i {
		b := f(x)
		if b {
			newList = append(newList, x)
		}
	}
	return newList
}

//    (Lists) Keep(func([]int) bool) Lists
func (i Lists) Keep(f func([]int) bool) Lists {
	var newList = Lists{}
	for _, x := range i {
		b := f(x)
		if b {
			newList = append(newList, x)
		}
	}
	return newList
}

//    (Ints) Keep(func(int) bool) Ints
func (i Ints) Keep(f func(int) bool) Ints {
	if i == nil {
		return nil
	}
	var newList = Ints{}
	for _, x := range i {
		b := f(x)
		if b {
			newList = append(newList, x)
		}
	}
	return newList
}

//    (Ints) Discard(func(int) bool) Ints
func (i Ints) Discard(f func(int) bool) Ints {
	if i == nil {
		return nil
	}
	var newList = Ints{}
	for _, x := range i {
		b := f(x)
		if !b {
			newList = append(newList, x)
		}
	}
	return newList
}
