package main

import (
	//"container/list"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"gopkg.in/goquery.v1"
)
type datoPagina struct{
	Materia string `json: materia"`
	Link    string `json:"link"`
    Topico string `json:"topico"`

}
//menu
func obtenerUrls(){
	var link string
	for{ 
		fmt.Println("Ingrese el Url del foro al cual desea suscribirse, 0 para salir, 1 para remover materias:")
		fmt.Scanln(&link)
		if link == "0" {
			break
		}
		if link == "1"{
			imprimirInfo()
			fmt.Println("ingrese el nombre de la materia")
			fmt.Scanln(&link)
			removerMateria(strings.ToLower(link))
		}else{
			
			err := ingresarMateria(link)
			if !err {
				println("error al ingresar la materia")
			}
		}
			
 	}
}
func removerMateria(materia string) error{
	existingData := make(map[string]datoPagina)
	file,err := os.ReadFile("dados.json")
	if err!=nil {
		fmt.Println("error")
		return err
	}
	json.NewDecoder(bytes.NewReader(file)).Decode(&existingData)
	delete(existingData,materia)
	filejs,err := json.Marshal(existingData)
	if err !=nil{
		return err
	}
	err = os.WriteFile("dados.json",filejs,0644)
	if err !=nil{
		return err
	}
	imprimirInfo()
	return nil
}
// obtiene y imprime toda la informacion guardada en el json
func imprimirInfo(){
	existingData := make(map[string]datoPagina)
	file,err := os.ReadFile("dados.json")
	if err != nil{
		println("error al imprimir:", err)
	}
	json.NewDecoder(bytes.NewReader(file)).Decode(&existingData)
	if len(existingData)>0 {
		fmt.Println(existingData)
		return
	}
	fmt.Println("no hay materias agregadas")
}

func salvarEnJson(info datoPagina) error{
	existingData := make(map[string]datoPagina)
	file , err := os.ReadFile("dados.json")
	if os.IsNotExist(err){
		existingData[info.Materia] = info
		data,err := json.Marshal(&existingData)
		if err != nil{		
			return err
		}
		err = os.WriteFile("dados.json",data,0644)
		if err !=nil{
			return err
		}
	}else{
		println("agreguei")
		json.NewDecoder(bytes.NewReader(file)).Decode(&existingData)
		_,existe:= existingData[info.Materia]
		if  existe{
			println("Url ya agregada")
			return nil
		}
		existingData[info.Materia] = info
		data,err := json.Marshal(&existingData)
		if err != nil{
			return err
		}
		os.WriteFile("dados.json",data,0644)
	}
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
	materia :=doc.Find(".breadcrumb-item").Eq(1)
	var infoPagina datoPagina
	infoPagina.Materia = strings.TrimSpace(materia.Text())
	infoPagina.Link = strings.TrimSpace(resp.Request.URL.String())
	infoPagina.Topico = strings.TrimSpace(ultimaActualizacion.Text())
	salvarEnJson(infoPagina)
	return true
}
}
func main() { 
obtenerUrls()
imprimirInfo()

}
