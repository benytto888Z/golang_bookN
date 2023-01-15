package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){

	var tmpl *template.Template
	var err error
	
	// check to see if we already have template in our cache
	_, inMap := tc[t]

	if !inMap{
		// need to create the template
		log.Println("creating template and add it to the cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	}else{
		// we have the template in cache
		log.Println("using cache template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w,nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error{
 templates := []string{
	fmt.Sprintf("./templates/%s", t),
	"./templates/base.layout.tmpl",
 }
 // parse the template
 tmpl, err := template.ParseFiles(templates...)

 if err != nil{
	return err
 }
 // add template to cache (map)
 tc[t] = tmpl
 return nil
}