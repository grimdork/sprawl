// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// validChars contains a safe subset of printable symbols.
const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-.,"

// RandString generates a random string suitable for passwords.
func RandString(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
		if err != nil {
			return s.String()
		}

		s.WriteByte(validChars[n.Int64()])
	}
	return s.String()
}
