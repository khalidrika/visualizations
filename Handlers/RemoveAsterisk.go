package groupie

func RemoveAsterisk(T []string) []string {
	for i := 0; i < len(T); i++ {
		if len(T[i]) > 1 && T[i][0] == '*' {
			T[i] = T[i][1:]
		}
	}
	return T
}
