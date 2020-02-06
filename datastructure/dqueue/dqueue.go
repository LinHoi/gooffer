package dqueue

type Dqueue struct {
	elememts []int
}

type IDqueue interface {
	IsEmpty() bool
	PopFront() int
	PopBack() int
	PushFront(int)
	PushBack(int)
	Front() int
	Back() int
}

func dDueueInstance() IDqueue {
	return &Dqueue{}
}

func  New() *Dqueue {
	var elements []int
	return &Dqueue{elememts:elements}
}
func (d *Dqueue) IsEmpty() bool {
	if len(d.elememts) == 0 {
		return true
	}
	return false
}
func (d *Dqueue) PopFront() int {
	front := d.elememts[0]
	d.elememts = d.elememts[1:]
	return front
}
func (d *Dqueue) PopBack() int {
	dLen := len(d.elememts)
	back := d.elememts[dLen-1]
	d.elememts = d.elememts[:dLen-1]
	return back
}

func (d *Dqueue) PushFront(ele int) {
	d.elememts = append(d.elememts,0)
	for i := len(d.elememts)-1; i > 0; i-- {
		d.elememts[i] = d.elememts[i-1]
	}
	d.elememts[0] = ele
}

func (d *Dqueue) PushBack(ele int) {
	d.elememts =append(d.elememts,ele )
}

func (d *Dqueue) Front() int {
	return d.elememts[0]
}
func (d *Dqueue) Back() int {
	return d.elememts[len(d.elememts)-1]
}