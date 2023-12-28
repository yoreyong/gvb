package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	QiNiu ImageType = 2
)

func (t ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t ImageType) String() string {
	var str string
	switch t {
	case Local:
		str = "Local"
	case QiNiu:
		str = "QiNiu"
	default:
		str = "Unknown path of image"
	}
	return str
}
