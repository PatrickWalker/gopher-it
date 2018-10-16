package emailparse

import "strings"

//getEmailDomain is a bad function. It has problems
//Use table driven tests to determine what some of those problems are
//loneliness doesn't count
func getEmailDomain(email string) string {
	components := strings.Split(email, "@")
	domain := components[1]
	return domain
}
