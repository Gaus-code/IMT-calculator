package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))
	tmpl.Execute(w, nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)

	imt := calculateIMT(weight, height)
	result := fmt.Sprintf("Ваш ИМТ: %.0f (%s)", imt, interpretIMT(imt))

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, result)
}

func calculateIMT(userKg, userHeight float64) float64 {
	return userKg / math.Pow(userHeight/100, 2)
}

func interpretIMT(imt float64) string {
	switch {
	case imt < 18.5:
		return "Недостаточный вес"
	case imt < 25:
		return "Норма"
	case imt < 30:
		return "Избыточный вес"
	default:
		return "Ожирение"
	}
}
