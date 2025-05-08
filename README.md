# goLang
//texto generado con IA.
Descripción

Este proyecto es un web scraper en Go que accede a páginas de foros de EVA (FING-UDELAR), obtiene el título del primer tópico publicado en el foro y guarda la información relevante (URL del curso, título del tópico y nombre de la materia) en un archivo JSON local para futuras comparaciones. El objetivo es detectar si hubo nuevas publicaciones.
Tecnologías utilizadas

    Lenguaje: Go (Golang)

    Librerías utilizadas:

        net/http: para hacer solicitudes HTTP

        github.com/PuerkitoBio/goquery: para parsear y navegar el HTML

        encoding/json: para guardar y leer los datos localmente

        os y io: para manipular archivos

Funcionalidad

    Ingreso de una materia:

        Realiza una solicitud HTTP a la URL del foro.

        Parsea el HTML con goquery.

        Extrae el título del primer tópico publicado.

        Extrae el nombre de la materia desde la barra breadcrumb.

        Construye un objeto con la información recolectada.

    Almacenamiento JSON:

        Verifica si el archivo "dados.json" existe.

        Si no existe, lo crea con el primer elemento.

        Si ya existe, carga los datos, evita duplicados y actualiza el archivo.

    Prevención de duplicados:

        Antes de agregar un nuevo elemento, se verifica si la URL ya está registrada en el JSON.

Ejemplo de JSON generado

{
"https://eva.fing.edu.uy/mod/forum/view.php?id=12345": {
"link": "https://eva.fing.edu.uy/mod/forum/view.php?id=12345",
"topico": "Resultados del examen de abril 2025"
}
}
Notas

    El programa detecta automáticamente si el link ya fue agregado anteriormente.

    El nombre de la materia se extrae desde el breadcrumb, ignorando su posición exacta ya que puede variar entre cursos.

    El archivo JSON se utiliza como base de comparación en futuras ejecuciones.