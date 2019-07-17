package kata

import (
	"strings"
)

func NumUniqueEmails(emails []string) (area int) {
	u := make(map[string]struct{})

	for _, v := range emails {
		emailParts := strings.Split(v, "@")
		local := emailParts[0]
		domain := emailParts[1]

		if strings.Contains(local, "+") {
			local = strings.Split(local, "+")[0]
		}

		local = strings.Replace(local, ".", "", -1)

		email := local + "@" + domain

		u[email] = struct{}{}
	}

	return len(u)
}
