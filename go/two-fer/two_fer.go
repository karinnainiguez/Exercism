package twofer

func ShareWith(name string) string {

	user := name
	if user == "" {
		user = "you"
	}

	return "One for " + user + ", one for me."
}
