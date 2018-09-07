package exercise

import (
	"bytes"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// It is not efficient to concatenate the string since string is an immutable type
	var buffer bytes.Buffer
	for i, v := range ip {
		if i > 0 {
			buffer.WriteByte('.')
		}
		// common mistake: buffer.WriteString(v)
		// https://stackoverflow.com/questions/29614943/stringer-implementation-without-using-sprintf
		buffer.WriteString(strconv.Itoa(int(v)))
	}

	return buffer.String()
	// alternative solution
	// return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
