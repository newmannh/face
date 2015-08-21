package main

import (
	"fmt"

	"github.com/newmannh/foundyou/fpp"
)

func main() {
	resp, err := fpp.DetectFace("http://www.uni-regensburg.de/Fakultaeten/phil_Fak_II/Psychologie/Psy_II/beautycheck/english/schemaanpassungen/w(61-64).jpg")
	fmt.Printf("Error: %v\nResponse: \n%v", err, resp)
}
