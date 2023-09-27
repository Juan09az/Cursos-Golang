package main

import (
	connection "Taller_parcial/Connection"
	"Taller_parcial/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
// Hombres y mujeres mayores
var EHM []Models.Estudiante1
conn.Raw("select  e.nombre, e.apellido, e.edad from estudiantes e where e.edad=(select MAX(e2.edad) from estudiantes e2) and e.gender='male'").Scan(&EHM)
for i := range EHM {
	fmt.Println(EHM[i].Nombre, EHM[i].Apellido, EHM[i].Edad)
}

var EMM []Models.Estudiante1
conn.Raw("select  e.nombre, e.apellido, e.edad from estudiantes e where e.edad=(select MAX(e2.edad) from estudiantes e2) and e.gender='female'").Scan(&EMM)
for i := range EMM {
	fmt.Println(EMM[i].Nombre, EMM[i].Apellido, EMM[i].Edad)
}

	//Estudiante con mejor promedio
	var MAVG Models.Estudiante2
	conn.Raw("SELECT e.nombre, e.apellido, AVG(c.nota) as promedio, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id GROUP BY e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about ORDER BY promedio DESC LIMIT 1").Scan(&MAVG)
	fmt.Println(MAVG.Nombre, MAVG.Apellido, MAVG.Promedio, MAVG.Edad, MAVG.Gender, MAVG.Email, MAVG.Phone, MAVG.Address, MAVG.Matriculado, MAVG.About)

	var MAVGCursos []Models.Curso2
	conn.Raw("SELECT c.id, c.curso, c.nota FROM cursos c JOIN (SELECT e.id, AVG(c.nota) as promedio FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id GROUP BY e.id ORDER BY promedio DESC LIMIT 1) sub1 ON c.estudiante_id = sub1.id").Scan(&MAVGCursos)
	fmt.Println("Cursos y sus notas del mejor estudiante: ")
	for i := range MAVGCursos {
		fmt.Println(MAVGCursos[i].Curso, MAVGCursos[i].Nota)

	}

	//Info Algebra Lineal
	var AVGAlgebra Models.Curso3
	conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='Algebra lineal' group by c.curso").Scan(&AVGAlgebra)
	fmt.Println("Promedio Algebra lineal: ", AVGAlgebra.Promedio)

	var Peores10Algebra []Models.Estudiante3

	conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'Algebra lineal' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Algebra)
	for i := range Peores10Algebra {
		fmt.Println(Peores10Algebra[i].Nombre, Peores10Algebra[i].Nota)
	}

	//Info Calculo Diferencial
	var AVGCalculo Models.Curso3
	conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='Calculo diferencial' group by c.curso").Scan(&AVGCalculo)
	fmt.Println("Promedio Calculo diferencial: ", AVGCalculo.Promedio)
	var Peores10Calculo []Models.Estudiante3
	conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'Calculo diferencial' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Calculo)
	for i := range Peores10Calculo {
		fmt.Println(Peores10Calculo[i].Nombre, Peores10Calculo[i].Nota)

	}

	//Info POO
	var AVGPOO Models.Curso3
	conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='POO' group by c.curso").Scan(&AVGPOO)
	fmt.Println("Promedio POO: ", AVGPOO.Promedio)
	var Peores10POO []Models.Estudiante3

	conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'POO' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10POO)
	for i := range Peores10POO {
		fmt.Println(Peores10POO[i].Nombre, Peores10POO[i].Nota)

	}

	//Info CTD
	var AVGCTD Models.Curso3
	conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='CTD' group by c.curso").Scan(&AVGCTD)
	fmt.Println("Promedio CTD: ", AVGCTD.Promedio)
	var Peores10CTD []Models.Estudiante3

	conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'CTD' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10CTD)
	for i := range Peores10CTD {
		fmt.Println(Peores10CTD[i].Nombre, Peores10CTD[i].Nota)
	}
*/

func main() {
	//Creación de la base de datos, e insert de todos los datos extraidos del JSON

	//conexión a la base de datos
	conn, err := connection.GetConnetion()
	if err != nil {
		log.Panic(err)
	}

	var exists, exists2 bool
	err = conn.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'cursos');").Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'estudiantes');").Scan(&exists2).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	if exists && exists2 {
		fmt.Println("Ya existen las tablas cursos y estudiantes.")
	} else {
		//Se migran las estructuras a la base de datos (por cada estructura una nueva tabla)

		conn.AutoMigrate(&Models.Estudiante{}, &Models.Curso{})

		//Lee el JSON
		fileContent, err := os.Open("./Static/generated.json")

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("The File is opened successfully...")

		defer fileContent.Close()

		byteResult, _ := ioutil.ReadAll(fileContent)

		//Crea un arreglo de estudiantes
		var students Models.Estudiantes
		//A cada llave-posición del JSON le genero un nuevo estudiante en el arreglo de estudiantes
		json.Unmarshal(byteResult, &students)

		//Recorre el arreglo de estudiantes y va guardando cada posición de estudiante en la base de datos
		//A su vez, para cada estudiante va guardando sus 4 cursos en la tabla cursos (especificando a que estudiante pertenece)
		for _, e := range students.Estudiantes {

			estudiante := Models.Estudiante{Index: e.Index, Nombre: e.Nombre, Apellido: e.Apellido, Edad: e.Edad, Gender: e.Gender,
				Email: e.Email, Phone: e.Phone, Address: e.Address, About: e.About, Matriculado: e.Matriculado}

			conn.Create(&estudiante)
			for _, c := range e.Cursos {
				curso := Models.Curso{Id: c.Id, Curso: c.Curso, Nota: c.Nota, EstudianteID: estudiante.ID}
				conn.Create(&curso)
			}
		}
	}
	//--------------------------------------APIs---------------------------------------
	//---------------------------------------------------------------------------------
	//sexoM
	//---------------------------------------------------------------------------------
	http.HandleFunc("/apisexom", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		var EHM []Models.Estudiante1
		conn.Raw("select e.index, e.nombre, e.apellido, e.edad from estudiantes e where e.edad=(select MAX(e2.edad) from estudiantes e2) and e.gender='male'").Scan(&EHM)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(EHM)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//---------------------------------------------------------------------------------
	//sexoF
	//---------------------------------------------------------------------------------
	http.HandleFunc("/apisexof", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		var EMM []Models.Estudiante1
		conn.Raw("select  e.nombre, e.apellido, e.edad from estudiantes e where e.edad=(select MAX(e2.edad) from estudiantes e2) and e.gender='female'").Scan(&EMM)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(EMM)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})

	//---------------------------------------------------------------------------------
	//Algebra
	//---------------------------------------------------------------------------------
	//promedio
	http.HandleFunc("/apiavgalgebra", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//promedio Algebra
		var AVGAlgebra Models.Curso3
		conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='Algebra lineal' group by c.curso").Scan(&AVGAlgebra)
		//fmt.Println("Promedio Algebra lineal: ", AVGAlgebra.Promedio)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(AVGAlgebra)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//peores estudiantes
	http.HandleFunc("/apipeoresalgebra", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//peores estudiantes
		var Peores10Algebra []Models.Estudiante3
		conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'Algebra lineal' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Algebra)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(Peores10Algebra)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//---------------------------------------------------------------------------------
	//Calculo
	//---------------------------------------------------------------------------------
	//promedio
	http.HandleFunc("/apiavgcalculo", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//promedio calculo
		var AVGcalculo Models.Curso3
		conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='Calculo diferencial' group by c.curso").Scan(&AVGcalculo)
		//fmt.Println("Promedio Algebra lineal: ", AVGAlgebra.Promedio)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(AVGcalculo)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//peores estudiantes
	http.HandleFunc("/apipeorescalculo", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//peores estudiantes
		var Peores10Calculo []Models.Estudiante3
		conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'Calculo diferencial' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Calculo)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(Peores10Calculo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})

	//---------------------------------------------------------------------------------
	//CTD
	//---------------------------------------------------------------------------------
	//promedio
	http.HandleFunc("/apiavgctd", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//promedio calculo
		var AVGctd Models.Curso3
		conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='CTD' group by c.curso").Scan(&AVGctd)
		//fmt.Println("Promedio Algebra lineal: ", AVGAlgebra.Promedio)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(AVGctd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//peores estudiantes
	http.HandleFunc("/apipeoresctd", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//peores estudiantes
		var Peores10Ctd []Models.Estudiante3
		conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'CTD' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Ctd)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(Peores10Ctd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//---------------------------------------------------------------------------------
	//POO
	//---------------------------------------------------------------------------------
	//promedio
	http.HandleFunc("/apiavgpoo", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//promedio calculo
		var AVGpoo Models.Curso3
		conn.Raw("select c.curso, AVG(c.nota) as promedio from cursos c  where c.curso='POO' group by c.curso").Scan(&AVGpoo)
		//fmt.Println("Promedio Algebra lineal: ", AVGAlgebra.Promedio)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(AVGpoo)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//peores estudiantes
	http.HandleFunc("/apipeorespoo", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//peores estudiantes
		var Peores10Poo []Models.Estudiante3
		conn.Raw("SELECT e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about, c.nota FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id WHERE c.curso = 'CTD' and c.estudiante_id in (select e.id from estudiantes) ORDER BY c.nota ASC LIMIT 10").Scan(&Peores10Poo)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(Peores10Poo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//---------------------------------------------------------------------------------
	//Estudiante con mejores notas
	//---------------------------------------------------------------------------------
	//mejor estudiante
	http.HandleFunc("/apibestest", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//mejor estudiante
		var MAVG Models.Estudiante2
		conn.Raw("SELECT e.nombre, e.apellido, AVG(c.nota) as promedio, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id GROUP BY e.nombre, e.apellido, e.edad, e.gender, e.email, e.phone, e.address, e.matriculado, e.about ORDER BY promedio DESC LIMIT 1").Scan(&MAVG)
		// fmt.Println(MAVG.Nombre, MAVG.Apellido, MAVG.Promedio, MAVG.Edad, MAVG.Gender, MAVG.Email, MAVG.Phone, MAVG.Address, MAVG.Matriculado, MAVG.About)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(MAVG)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//notas
	http.HandleFunc("/apinotasbestest", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		//notas del mejor estudiante
		var MAVGCursos []Models.Curso2
		conn.Raw("SELECT c.curso, c.nota FROM cursos c JOIN (SELECT e.id, AVG(c.nota) as promedio FROM estudiantes e JOIN cursos c ON c.estudiante_id = e.id GROUP BY e.id ORDER BY promedio DESC LIMIT 1) sub1 ON c.estudiante_id = sub1.id").Scan(&MAVGCursos)

		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(MAVGCursos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})
	//---------------------------------------------------------------------------------
	//Listar todos los estudiantes
	//---------------------------------------------------------------------------------
	http.HandleFunc("/apiall", func(w http.ResponseWriter, r *http.Request) {
		// Crea una instancia de la estructura
		conn, err := connection.GetConnetion()
		_ = conn
		if err != nil {
			log.Panic(err)
		}
		var AllEstudiante []Models.Estudiante4
		conn.Raw("SELECT nombre,apellido,edad,gender,email,phone,address,about,matriculado from estudiantes").Scan(&AllEstudiante)
		// Convierte la estructura en un JSON
		jsonData, err := json.Marshal(AllEstudiante)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println(jsonData)
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")

		// Escribe la respuesta JSON
		w.Write(jsonData)
	})

	http.Handle("/", http.FileServer(http.Dir("Templates")))
	http.ListenAndServe("localhost:8080", nil)
}
