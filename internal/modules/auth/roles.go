package auth

func HasPermission(userRole string, allowedRoles []string) bool {
	if userRole == RoleAdmin || userRole == RoleSuper {
		return true
	}
	for _, role := range allowedRoles {
		if role == userRole {
			return true
		}
	}
	if len(allowedRoles) == 0 {
		return true
	}
	return false
}

var rolesMap = map[string]int{
	RoleUser:  0,
	RoleAdmin: 1,
	RoleSuper: 2,
}

func HasHigherRole(a string, b string) bool {
	aInt := rolesMap[a]
	bInt := rolesMap[b]
	return aInt > bInt
}

func IsValidRole(a string) bool {
	_, ok := rolesMap[a]
	return ok
}
