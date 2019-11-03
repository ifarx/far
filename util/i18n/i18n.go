package i18n

type MessageSource struct {
	_l string            /*language*/
	_m map[string]string /*map*/
}

func (m *MessageSource) Get(k string) string {
	return m._m[k]
}

func (m *MessageSource) Push(k, v string) {
	m._m[k] = v
}

func CreateMessageSource(l string) *MessageSource {
	return &MessageSource{l, make(map[string]string)}
}

type I18n struct {
	_ms map[string]*MessageSource /*messgae_source_map*/
}

func (i *I18n) Message(l, k string) string {
	m := i._ms[l]
	if m == nil {
		return k
	}
	v := m.Get(k)
	if v == "" {
		return k
	}
	return v
}

func (i *I18n) PushMessage(l, k, v string) {
	m := i._ms[l]
	if m == nil {
		m = CreateMessageSource(l)
		i._ms[l] = m
	}
	m.Push(k, v)
}

func CreateI18n() *I18n {
	return &I18n{map[string]*MessageSource{}}
}

var __d_i18n = CreateI18n()

func Message(l, k string) string {
	return __d_i18n.Message(l, k)
}

func PushMessage(l, k, v string) {
	__d_i18n.PushMessage(l, k, v)
}
