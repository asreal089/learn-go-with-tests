package depinjection

import (
	"bytes"
	"fmt"
)

func Great(wr *bytes.Buffer, name string) {
	fmt.Fprintf(wr, "Hello, %s", name)
}
