package main

//这就是最优解

// func main() {
// 	a := []string{"gin", "zen", "gig", "msg"}
// 	b := uniqueMorseRepresentations(a)
// 	fmt.Println(b)
// }

func uniqueMorseRepresentations(words []string) int {
	e := make(map[string]bool)
	passWords := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, v := range words {
		var c, d string
		for _, b := range v {
			c = passWords[rune(b)-97]
			d = d + c
		}
		e[d] = true
	}
	return len(e)
}
