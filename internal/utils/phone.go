package utils

func IsValidPhone(phone string) bool {
	if phone[0] != '0' {
		return false
	}
	for _, code := range phone[1:] {
		if code < '0' || code > '9' {
			return false
		}
	}
	if len(phone) != 11 {
		return false
	}
	return true
}
