package main

import (
	"fmt"

	"github.com/koykov/dlqdump/encoder"
)

func main() {
	m := encoder.Marshaller{}
	itm0 := ItemNoPointerReceiver{1, 2}
	itm1 := ItemPointerReceiver{1, 2}
	_, err := m.Encode(nil, itm0)
	fmt.Println(err)
	_, err = m.Encode(nil, itm1)
	fmt.Println(err)
	_, err = m.Encode(nil, &itm1)
	fmt.Println(err)
}
