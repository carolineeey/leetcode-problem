package solution

import "strings"

func NomUniqueEmails(emails []string) int {
	seen := map[string]struct{}{}

	for _, email := range emails {
		email = strings.TrimSpace(email) // avoid hidden whitespace differences
		at := strings.Split(email, "@")  // split local and domain
		local := at[0]
		domain := at[1]

		// handle plus: ignore everything after +
		if p := strings.Index(local, "+"); p != -1 {
			local = local[:p]
		}

		// remove dots
		local = strings.ReplaceAll(local, ".", "")

		// reconstruct and add to set
		cleaned := local + "@" + domain
		seen[cleaned] = struct{}{}
	}

	return len(seen)
}
