package queue

type Queue []interface{}

func (q *Queue) Add(i interface{}) {
	*q = append(*q, i)
}

func (q *Queue) Pull() (ret interface{}, ok bool) {

	if len(*q) == 0 {
		ret = -1
		ok = false
		return
	}

	ret = (*q)[0]
	ok = true
	*q = (*q)[1:]
	return
}
