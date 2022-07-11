package byteconv

func byteconv(x interface{}) ([]byte, bool) {
	switch x.(type) {
	case []byte:
		return x.([]byte), true
	case *[]byte:
		return *x.(*[]byte), true
	}
	return nil, false
}

func byte2str(b byte) string {
	return string(b)
}
