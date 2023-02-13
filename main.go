package main

func main() {
	card2 := []string{"!23", newCard()}
	card2 = append(card2, "123223", "124242")
	for i := range card2 {
		println(card2[i])
	}
}
func newCard() string {
	return "ngu xi"
}
