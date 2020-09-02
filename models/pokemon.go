package models

// Pokemon ..,
type Pokemon struct {
	Nombre    string
	Tipo      map[int]string
	Stats     map[string]int
	Habilidad map[int]string
	Imagen    map[string]string
}
