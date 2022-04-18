package iface_copy

type T struct {
	id int32
	ct float32
	nm string
	pl []byte
}

type TI struct{}

func (i TI) Copy(x interface{}) (interface{}, error) {
	var t T
	if t1, ok := x.(*T); ok {
		t = *t1
		return t, nil
	}
	return x, nil
}

type SI struct{}

func (i SI) Copy(x interface{}) (interface{}, error) {
	var t interface{}
	switch x.(type) {
	case *bool:
		t = *x.(*bool)
	case *int:
		t = *x.(*int)
	case *int8:
		t = *x.(*int8)
	case *int16:
		t = *x.(*int16)
	case *int32:
		t = *x.(*int32)
	case *int64:
		t = *x.(*int64)
	case *uint:
		t = *x.(*uint)
	case *uint8:
		t = *x.(*uint8)
	case *uint16:
		t = *x.(*uint16)
	case *uint32:
		t = *x.(*uint32)
	case *uint64:
		t = *x.(*uint64)
	case *float32:
		t = *x.(*float32)
	case *float64:
		t = *x.(*float64)
	case []byte:
		p := x.([]byte)
		t = append([]byte(nil), p...)
	case *[]byte:
		p := *x.(*[]byte)
		t = append([]byte(nil), p...)
	case *string:
		t = *x.(*string)
	default:
		return x, nil
	}
	return t, nil
}
