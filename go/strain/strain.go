package strain

type Ints []int
type Lists [][]int
type Strings []string

func (c Ints) Keep(fun func(int) bool) Ints {
	newInts := Ints{}
	for _, i := range c {
		if fun(i) == true {
			newInts = append(newInts, i)
		}
	}
	return newInts
}
func (c Ints) Discard(fun func(int) bool) Ints {
	newInts := Ints{}
	for _, i := range c {
		if fun(i) == false {
			newInts = append(newInts, i)
		}
	}
	return newInts
}
func (c Strings) Keep(fun func(string) bool) Strings {
	newStrings := Strings{}
	for _, i := range c {
		if fun(i) == true {
			newStrings = append(newStrings, i)
		}
	}
	return newStrings
}
func (c Strings) Discard(fun func(string) bool) Strings {
	newStrings := Strings{}
	for _, i := range c {
		if fun(i) == false {
			newStrings = append(newStrings, i)
		}
	}
	return newStrings
}
func (c Lists) Keep(fun func([]int) bool) Lists {
	newLists := Lists{}
	for _, i := range c {
		if fun(i) == true {
			newLists = append(newLists, i)
		}
	}
	return newLists
}
func (c Lists) Discard(fun func([]int) bool) Lists {
	newLists := Lists{}
	for _, i := range c {
		if fun(i) == false {
			newLists = append(newLists, i)
		}
	}
	return newLists
}
