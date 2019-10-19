package main

import "fmt"
import "regexp"

func main() {
	var text string
	fmt.Print("masukan text : ")
	fmt.Scanln(&text)

	var regex, err = regexp.Compile(`[a-z]+\d`)

	if err != nil {
		fmt.Println(err.Error())
	}
	var hasil = regex.FindAllString(text, 5)
	fmt.Println(hasil)

	var cocok = regex.MatchString(text)
	fmt.Println(cocok)

	var index1 = regex.FindStringIndex(text)
	fmt.Println(index1)

	var stringbaru = regex.ReplaceAllString(text, "MANGGA")
	fmt.Println(stringbaru)
}
