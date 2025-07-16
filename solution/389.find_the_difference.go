package solution

func FindTheDifference(s string, t string) byte {
	var sumS, sumT int

	for i := 0; i < len(s); i++ {
		sumS += int(s[i])
	}

	for i := 0; i < len(t); i++ {
		sumT += int(t[i])
	}
	return byte(sumT - sumS)
}
