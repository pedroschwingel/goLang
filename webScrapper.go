package main

import (
	//"container/list"
	"fmt"
	"net/http"
	"strings"
	"gopkg.in/goquery.v1"
	"encoding/json"
	"os"
)

type datoPagina struct{
	link string 
	topico string 
}
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
func imprimirInfo(){

}

func salvarEnJson(info datoPagina) error{
	var existingData []datoPagina
	data,err := json.Marshal(info)
	if err != nil{		
			return err
	}
	file , err := os.Open("dados.json")
	if os.IsNotExist(err){
		err := os.WriteFile("dados.json",data,0644)
		if err !=nil{
			return err
		}
	}else{
		json.NewDecoder(file).Decode(existingData)
		existingData = append(existingData, info)
		data,err = json.Marshal(&existingData)
		if err != nil{
			return err
		}
		os.WriteFile("dados.json",data,0644)
	}
	defer file.Close()
	return nil
}

func ingresarMateria(url string)bool{
	resp,err := http.Get(url)
	if err != nil {
		return false
	}else{
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		return false
	}
	ultimaActualizacion :=doc.Find("tr.discussion th a").Eq(0)
	var infoPagina datoPagina
	infoPagina.link = resp.Request.URL.String()
	infoPagina.topico = ultimaActualizacion.Text()
	salvarEnJson(infoPagina)
	return true
}
}



func imprimirMaterias(pages []*http.Response){
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
func main() { 

	



	pages := obtenerUrls()
	fmt.Println(len(pages))
	for i := 0; i < len(pages); i++ {
		doc, err := goquery.NewDocumentFromReader(pages[i].Body)
		if err!=nil {
			fmt.Println("bad things")
		}
		nuevaMateria=ingresarMateria(doc)
		nombreCurso :=doc.Find(".breadcrumb-item").Eq(1)
		fmt.Println("nombre del curso es:", nombreCurso.Text())
		fmt.Println("las ultimas 5 actualizaciones fueron:")
		dis := doc.Find("tr.discussion th a").Slice(0,5)
		dis.Each(func(index int, element *goquery.Selection)  {
			fmt.Println(strings.TrimSpace(element.Text()))
		})
	}

}
