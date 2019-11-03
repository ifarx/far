package search

type Order uint8

const DIS = Order(0)

const ASC = Order(1)

const DESC = Order(2)

type Rule struct {
	_f string /*field*/
	_o Order  /*order*/
}

func (r Rule) Field() string {
	return r._f
}

func (r Rule) Order() Order {
	return r._o
}

type Sort struct {
	_rs []Rule /*rules*/
}

type Range struct {
	_s uint64 /*start*/
	_c uint64 /*count*/
}

func (r Range) Start() uint64 {
	return r._s
}

func (r Range) Count() uint64 {
	return r._c
}
