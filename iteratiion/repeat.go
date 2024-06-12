package iteratiion

func Repeat(char string, repeatedCount int) string {
	chrs := ""
	for i := 0; i < repeatedCount; i++ {
		chrs += char
	}

	return chrs
}
