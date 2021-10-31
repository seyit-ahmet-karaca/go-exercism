package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (il IntList) Foldr(f binFunc, acc int) int {
	for i := len(il) - 1; i >= 0; i-- {
		acc = f(il[i], acc)
	}
	return acc
}

func (il *IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range *il {
		initial = fn(initial, v)
	}
	return initial
}

func (il *IntList) Filter(pred predFunc) IntList {
	list  := IntList{}
	for _, x := range *il {
		if pred(x) {
			list = append(list, x)
		}
	}
	return list
}

func (il *IntList) Length() int {
	return len(*il)
}

func (il *IntList) Map(fn unaryFunc) IntList {
	var list = IntList{}
	for _, x := range *il {
		list = append(list , fn(x))
	}
	return list
}

func (il *IntList) Reverse() IntList {
	var list = IntList{}
	for i:= len(*il) -1 ; i >= 0 ; i-- {
		list = append(list , (*il)[i])
	}
	return list
}

func (il *IntList) Append(list IntList) IntList {
	return append(*il, list...)
}

func (il *IntList) Concat(list []IntList) IntList {
	intList := *il
	for _,i := range list{
		intList = append(intList, i...)
	}
	return intList
}
