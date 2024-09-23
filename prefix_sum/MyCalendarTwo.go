package prefixsum

type MyCalendarTwo struct {
	single [][]int
	double [][]int
}

func MyCalendarTwoConstructor() MyCalendarTwo {
	return MyCalendarTwo{
		single: [][]int{},
		double: [][]int{},
	}
}

func interset(x0, x1, y0, y1 int) []int {
	if x0 > y0 {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if y0 >= x1 {
		return []int{}
	}
	return []int{max(x0, y0), min(x1, y1)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.double {
		z := interset(start, end, v[0], v[1])
		if len(z) > 0 {
			return false
		}
	}
	for _, v := range this.single {
		z := interset(start, end, v[0], v[1])
		if len(z) > 0 {
			this.double = append(this.double, z)
		}
	}

	this.single = append(this.single, []int{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
