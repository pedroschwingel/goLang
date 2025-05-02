package main

import (
	"fmt"
	"net/http"
	"gopkg.in/goquery.v1"
)

func main() {
	resp, err = http.Get("https://eva.fing.edu.uy/mod/forum/view.php?id=217947")
	doc, err= goquery.NewDocumentFromReader(resp.Body)
	doc.Find("td.topic.starter.a").Each(func(i int, s *goquery.Selection)){
		titulo := s.Text()
		fmt.PrintLn("titulo: ",titulo)

	}

}
