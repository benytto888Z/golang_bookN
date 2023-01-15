package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){

	// create template cache

	tc, err := createTemlateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache

	t, ok := tc[tmpl]

	if  !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err =t.Execute(buf,nil)
	if err != nil {
		log.Println(err)
	}

	
	// render the template

	_,err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func createTemlateCache() (map[string]*template.Template, error)  {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil{
		return myCache,err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts,err := template.New(name).ParseFiles(page)

		if err != nil{
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil{
			return myCache,err
		}

		if len(matches) > 0{
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return myCache,err
			}
		}

		myCache[name] = ts
	}
	return myCache,nil
}
