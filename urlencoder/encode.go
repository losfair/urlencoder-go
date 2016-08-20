package urlencoder

import "fmt"

func EncodeComponent(str string) string {
	chars := []byte(str)
	ret := ""
	for _,ch := range chars {
		if (ch>='a' && ch<='z') || (ch>='A' && ch<='Z') || (ch>='0' && ch<='9') || ch=='_' || ch=='-' {
			ret += string(ch)
		} else {
			ret += fmt.Sprintf("%%%2x",ch)
		}
	}
	return ret
}
