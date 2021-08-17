package chatLib

import (
	"fmt"
	"testing"
)

func TestDecodeBas58(t *testing.T) {
	s := DecodeLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcReyHaWuwXm7cHgyw9KgQHWhiPHAVeAEXEqLzBVvAqxdx18SEtUbXXkoDnqMGPGzKNwmj42WyS3t7Cb2XwX9uud6CvPxiufd6Hhfazsuzf8yS1aEsMK1oNymzXjkiXskVwTxBUcdSki")

	fmt.Println(s)
}
