package utility

func IsSubset[T comparable](subset, superset []T) bool {
	set := make(map[T]struct{})
	for _, v := range superset {
		set[v] = struct{}{}
	}

	for _, v := range subset {
		if _, found := set[v]; !found {
			return false
		}
	}
	return true
}
