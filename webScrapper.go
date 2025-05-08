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
func menu(){
	var link string
	for{ 
		fmt.Println("Ingrese una de las opciones:" )
		fmt.Println("\t -Url del foro que deseas subscribirte")
		fmt.Println("\t -0 para salir")
		fmt.Println("\t -1 para imprimir")
		fmt.Println("\t -att para recibir actualizaciones")
		fmt.Println("\t -rmv para remover alguna materia")
		fmt.Scanln(&link)
		link =strings.ToLower(link)
		if link == "0" {
			break
		}
		if link == "1"{
			imprimirInfo()
			continue
			
		}
		if link == "rmv"{
			imprimirInfo()
			fmt.Println("ingrese el nombre de la materia")
			fmt.Scanln(&link)
			removerMateria(link)
			continue
		}
		if link=="att" {
			actualizarDatos()
			continue
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
		for k:= range existingData{
			if existingData[k].Topico == "" {
				println("la materia posée una url privada:",existingData[k].Link)
			}else{
				fmt.Println(existingData[k])
			}
		}
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
	var infoPagina datoPagina
	ultimaActualizacion :=doc.Find("tr.discussion th a").Eq(0)
	doc.Find(".breadcrumb-item a").Each(func(i int, s *goquery.Selection){
		print(s.Text())
		href,exists := s.Attr("href")
		if exists && strings.Contains(href, "/course/view.php"){
			infoPagina.Materia = strings.TrimSpace(s.Text())
		}
	})
	infoPagina.Link = strings.TrimSpace(resp.Request.URL.String())
	infoPagina.Topico = strings.TrimSpace(ultimaActualizacion.Text())
	salvarEnJson(infoPagina)
	return true
}
}
func actualizarDatos(){
	datosExistentes := make(map[string]datoPagina)
	var cambios []datoPagina
	file,err := os.ReadFile("dados.json")
	if err != nil{
		fmt.Println("error")
	}
	json.NewDecoder(bytes.NewReader(file)).Decode(&datosExistentes)
	for k := range datosExistentes{
		resp,err := http.Get(datosExistentes[k].Link)
		if err != nil{
			fmt.Println("error")
			break
		}
		doc,err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil{
			fmt.Println("error")
			break
		}
		ultimaActualizacion := doc.Find("tr.discussion th a").Eq(0)
		if strings.TrimSpace(ultimaActualizacion.Text()) != datosExistentes[k].Topico {
			var cambio datoPagina
			temp := datosExistentes[k]
			temp.Topico = ultimaActualizacion.Text()
			datosExistentes[k] = temp
			cambio.Link = datosExistentes[k].Link
			cambio.Materia = strings.ToLower(datosExistentes[k].Materia)
			cambio.Topico = strings.TrimSpace(datosExistentes[k].Topico)
			cambios = append(cambios, cambio)
		}
	}
	data,err :=json.Marshal(&datosExistentes)
	if err != nil{
		fmt.Println("error")
		return
	}
	os.WriteFile("dados.json",data,0446)
	notificar(cambios)
}
func notificar(cambios []datoPagina){
	fmt.Println("hay ",len(cambios)," nuevas notificaciones")
	for i := 0; i < len(cambios); i++ {
		fmt.Println("En", cambios[i].Materia)
		fmt.Println("\t", cambios[i].Topico)
	}
}
func main() { 
menu()

}
