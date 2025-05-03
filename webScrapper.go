package main

import (
	//"container/list"
	"fmt"
	"net/http"
	"strings"
	"gopkg.in/goquery.v1"
)


func obtenerUrls()[]*http.Response{
	var paginas []*http.Response
	var link string
	for{ 
		fmt.Println("Ingrese el Url del foro al cual desea suscribirse o 0 para salir:")
		fmt.Scanln(&link)
		if link == "0" {
			break
		}
		resp,err := http.Get(link)
		if err != nil {
			fmt.Println("error, intente nuevamente")
			}else{
				paginas = append(paginas, resp)
			}
		}
		return paginas
}

func main() {
	pages := obtenerUrls()
	fmt.Println(len(pages))
	for i := 0; i < len(pages); i++ {
		doc, err := goquery.NewDocumentFromReader(pages[i].Body)
		if err!=nil {
			fmt.Println("bad things")
		}
		nombreCurso :=doc.Find(".breadcrumb-item").Eq(1)
		fmt.Println("nombre del curso es:", nombreCurso.Text())
		fmt.Println("las ultimas 5 actualizaciones fueron:")
		dis := doc.Find("tr.discussion th a").Slice(0,5)
		dis.Each(func(index int, element *goquery.Selection)  {
			fmt.Println(strings.TrimSpace(element.Text()))
		})
	}

}
 