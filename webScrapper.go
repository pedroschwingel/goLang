package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("cibola")
	resp, err := http.Get("https://eva.fing.edu.uy/mod/forum/view.php?id=217947")

}
