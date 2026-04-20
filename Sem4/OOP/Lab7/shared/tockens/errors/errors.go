package errors

import (
	"fmt"
	"golang.org/x/crypto/nacl/auth"
)

var ErrInvalidSize = fmt.Errorf("invalid key syze: it must be %s", auth.KeySize)
