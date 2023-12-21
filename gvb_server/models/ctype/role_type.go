package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 普通登录人
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用的用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "Administrator"
	case PermissionUser:
		str = "User"
	case PermissionVisitor:
		str = "Visitor"
	case PermissionDisableUser:
		str = "Disabled User"
	default:
		str = "Unknown user"
	}
	return str
}
