package ctype

import "encoding/json"

type SignStatus int

const (
	SignWithQQ     SignStatus = 1
	SignWithGithub SignStatus = 2
	SignWithEmail  SignStatus = 3
)

// MarshalJSON 将 SignStatus 转换为 JSON
func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string

	switch s {
	case SignWithQQ:
		str = "Sign with QQ"
	case SignWithGithub:
		str = "Sign with Github"
	case SignWithEmail:
		str = "Sign with Email"
	default:
		str = "Unknown Account"
	}

	return str
}
