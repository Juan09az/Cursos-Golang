package Models

import "gorm.io/gorm"

type Estudiantes struct {
	Estudiantes []EstudianteJSON `json:"estudiantes"`
}

type Estudiantes2 struct {
	Estudiantes []Estudiante2
}

type Estudiante struct {
	gorm.Model
	Index       int
	Nombre      string
	Apellido    string
	Edad        int
	Gender      string
	Email       string
	Phone       string
	Address     string
	About       string
	Matriculado string
	Cursos      []Curso `gorm:"ForeignKey:EstudianteID"`
}

type Curso struct {
	gorm.Model
	Id           int
	Curso        string
	Nota         float64
	EstudianteID uint
}

type EstudianteJSON struct {
	Index       int     `json:"index"`
	Nombre      string  `json:"nombre"`
	Apellido    string  `json:"apellido"`
	Edad        int     `json:"edad"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	About       string  `json:"about"`
	Matriculado string  `json:"matriculado"`
	Cursos      []Curso `json:"cursos"`
}

type CursoJSON struct {
	Id    int     `json:"id"`
	Curso string  `json:"curso"`
	Nota  float64 `json:"nota"`
}

type Estudiante1 struct {
	Index       int    `json:"index"`
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Edad        int    `json:"edad"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	About       string `json:"about"`
	Matriculado string `json:"matriculado"`
}

type Estudiante2 struct {
	Index       int
	Nombre      string
	Apellido    string
	Promedio    float64
	Edad        int
	Gender      string
	Email       string
	Phone       string
	Address     string
	About       string
	Matriculado string
}
type Estudiante4 struct {
	Nombre      string
	Apellido    string
	Edad        int
	Gender      string
	Email       string
	Phone       string
	Address     string
	About       string
	Matriculado string
}

type Estudiante3 struct {
	Index       int
	Nombre      string
	Apellido    string
	Edad        int
	Gender      string
	Email       string
	Phone       string
	Address     string
	About       string
	Matriculado string
	Nota        float64
}

type Curso2 struct {
	Id    int
	Curso string
	Nota  float64
}

type Curso3 struct {
	Curso    string
	Promedio float64
}
