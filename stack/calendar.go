package stack

type MyCalendar struct {
	dates [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		dates: [][]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.dates) == 0 {
		this.dates = append(this.dates, []int{start, end})
		return true
	}

	for i := 0; i < len(this.dates); i++ {
		if this.dates[i][1] <= start || end <= this.dates[i][0] {
			continue
		} else {
			return false
		}
	}
	this.dates = append(this.dates, []int{start, end})
	return true
}
