package services

type permissions struct {
	Roles []string
}

func (p *permissions) IsAllowedToExecute(role string) bool {
	for _, r := range p.Roles {
		if r == role {
			return true
		}
	}
	return false
}

var AdminPermission = &permissions{
	Roles: []string{"manager", "admin"},
}
