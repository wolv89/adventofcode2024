package structures

/*
 * I think it's not idiomatic to have a "utilities" package or file but.... sometimes... maybe?
 */

func AbsInt(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
