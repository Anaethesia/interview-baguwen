package golang


const (
	c = 10 // iota = 0
	d = iota // iota = 1
	e = iota // iota = 2
	f = "hello" // iota = 3
	g = "hello" // iota = 4
	h = iota // iota = 5
	i = iota // iota = 6
	j = 0 // iota = 7
	k = 0 // iota = 8
	l, m = iota, iota // iota = 9
	n, o = iota, iota // iota = 10
	p = iota + 1 // iota = 11
	q = iota + 1 // iota = 12
	_ = iota + 1 // iota = 13
	r = iota * iota // iota = 14
	s = iota * iota // iota = 15
	t = r // iota = 16
	u = r // iota = 17
	v = 1 << iota // iota = 18
	w = 1 << iota // iota = 19
	x = iota * 0.01 // iota = 20
	y float32 = iota * 0.01 // iota = 21
	z float32 = iota * 0.01 // iota = 22
)