package helper

func GuardLastEnabled(sibling1, sibling2 bool, requested bool) bool {
	if !requested && !sibling1 && !sibling2 {
		return true
	}

	return requested
}
