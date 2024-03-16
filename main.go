package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func executeTemplate(writer http.ResponseWriter, templatePath string) {
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(writer, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(writer, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(writer, "There was an error returning data", http.StatusInternalServerError)
		return
	}
}
func homeHandler(writer http.ResponseWriter, req *http.Request) {
	//Setting Response Header
	writer.Header().Set("Content-Type", "text/html")
	//fmt.Fprint(writer, "<h1>Welcome to my awesome site </h1>")
	//DONE TO HANDLE WINDOWS BACK SLASH
	tplFilePath := filepath.Join("templates", "home.gohtml")
	//tpl, err := template.ParseFiles(tplFilePath)
	////TODO Remove Pabic
	//if err != nil {
	//	log.Printf("Parsing template: %v", err)
	//	http.Error(writer, "There was an error parsing the template", http.StatusInternalServerError)
	//	return
	//}
	//err = tpl.Execute(writer, nil)
	//if err != nil {
	//	log.Printf("Executing template: %v", err)
	//	http.Error(writer, "There was an error returning data", http.StatusInternalServerError)
	//	return
	//}
	executeTemplate(writer, tplFilePath)
}

func contactHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	templatePath := filepath.Join("templates", "contact.gohtml")
	//tpl, err := template.ParseFiles(templatePath)
	//if err != nil {
	//	log.Printf("Error parsing template: %v", err)
	//	http.Error(writer, "Error occured parsing template", http.StatusInternalServerError)
	//
	//	return
	//
	//}
	//err = tpl.Execute(writer, nil)
	//if err != nil {
	//	log.Printf("Error occured at the presentation layer", http.StatusInternalServerError)
	//	return
	//}
	//fmt.Fprint(writer,
	//	"<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:philipokiokio@gmail.com\">my mail</p>")

	executeTemplate(writer, templatePath)

}

func faqHandler(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	//fmt.Fprint(writer, "<h1>FAQ</h1>\n <p>This is the FAQ page</p>\n<ul> <li>\n    <b>Is there a free version?</b> </li> </ul>")
	templatePath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(writer, templatePath)

}

func galleryHandler(writer http.ResponseWriter, req *http.Request) {
	galleryId := chi.URLParam(req, "id-here")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "This is the parameters: %v", galleryId)

}

//
//func pathHandler(writer http.ResponseWriter, req *http.Request) {
//
//	switch req.URL.Path {
//	case "/":
//		homeHandler(writer, req)
//	case "/contact":
//		contactHandler(writer, req)
//	case "/faq":
//		faqHandler(writer, req)
//
//	default:
//		writer.Header().Set("Content-Type", "text/html")
//		writer.WriteHeader(http.StatusNotFound)
//		http.Error(writer, "Page does not exist", http.StatusNotFound)
//		//fmt.Fprint(writer, "Page does not exist")
//
//	}
//
//	//if req.URL.Path == "/" {
//	//	homeHandler(writer, req)
//	//} else if req.URL.Path == "/contact" {
//	//	contactHandler(writer, req)
//	//}
//}

//type Router struct {
//}
//
//func (router Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
//
//	switch req.URL.Path {
//	case "/":
//		homeHandler(writer, req)
//	case "/contact":
//		contactHandler(writer, req)
//	case "/faq":
//		faqHandler(writer, req)
//	default:
//		writer.Header().Set("Content-Type", "text/html")
//		writer.WriteHeader(http.StatusNotFound)
//		http.Error(writer, "Page does not exist", http.StatusNotFound)
//		//fmt.Fprint(writer, "Page does not exist")
//
//	}
//}

func main() {
	//var router Router
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)

	fmt.Println("Listening on port :4000")

	//http.ListenAndServe(":4000", router)
	//http.ListenAndServe(":4000", nil)
	chi_router := chi.NewRouter()
	//middleware for logging
	chi_router.Use(middleware.Logger)

	chi_router.Get("/", homeHandler)
	chi_router.Get("/contact", contactHandler)
	chi_router.Get("/faq", faqHandler)
	chi_router.Get("/galleries/{id-here}", galleryHandler)

	chi_router.NotFound(func(writer http.ResponseWriter, req *http.Request) {
		http.Error(writer, "not found", http.StatusNotFound)
	})
	http.ListenAndServe(":4000", chi_router)

}
