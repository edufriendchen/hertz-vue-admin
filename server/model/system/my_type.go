package system

type MyType interface {
	GetValue() []string
}

type permission []string

// func (m *permission) Scan(val interface{}) error {
// 	s := val.([]uint8)
// 	ss := strings.Split(string(s), "|")
// 	*m = ss
// 	return nil
// }
