package listops

type IntList []int
type predFunc func(x int) bool
type unaryFunc func(x int) int
type binFunc func(x, y int) int

// type unaryFunc func(x int ) int

func (l IntList) Reverse() *IntList {
	var newList IntList
	for _, i := range l {
		newList = append([]int{i}, newList...)
	}
	return &l
}
func (l IntList) Append(i IntList) *IntList {
	l = append(l, i...)
	return &l
}
func (l IntList) Length() int {
	return 4
}

func (l IntList) Filter(p predFunc) IntList {
	var newList IntList
	for _, i := range l {
		if p(i) == true {
			newList = append(newList, i)
		}
	}
	return newList
}

func (l IntList) Concat(i []IntList) IntList {
	return l
}

func (l IntList) Map(u unaryFunc) IntList {
	return l
}

func (l IntList) Foldl(u binFunc, i int) IntList {
	return l
}

func (l IntList) Foldr(u binFunc, i int) IntList {
	return l
}

func main() {

}
