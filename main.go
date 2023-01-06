package main

import (
	"fmt"
	"html/template"
	Back "lib/GoAssets"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./styleAssets"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	http.HandleFunc("/", home)
	//	http.HandleFunc("/Recherche", cherche)
	http.HandleFunc("/test", carouselle)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}

var ListeArtist = Back.GetApiArtists()

/*func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./WebAssets/*.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.ExecuteTemplate(rw, "Test", ListeArtist)
	if err != nil {
		fmt.Println(err)
	}
}*/

func home(rw http.ResponseWriter, r *http.Request) {

	//Ensure that the endpoint only gets POST requests
	if r.Method != "POST" {
		// http.Error(w, "Bad request - Go away!", 405)

		tmpl2, _ := template.ParseGlob("./WebAssets/*.html")
		err := tmpl2.ExecuteTemplate(rw, "Test", ListeArtist)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := r.ParseForm()
		if err != nil {
			return
		}
		if r.FormValue("bouton_envoie") == "Submit" {
			envoie := ""
			envoie = r.FormValue("Artistesearch")
			fmt.Println(envoie)
			if envoie == "" {
				tmpl2, _ := template.ParseGlob("./WebAssets/*.html")
				err := tmpl2.ExecuteTemplate(rw, "Test", ListeArtist)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				ArtisteRecherche :=Back.Chercher(ListeArtist,envoie)
				tmpl2, _ := template.ParseGlob("./WebAssets/*.html")
				err := tmpl2.ExecuteTemplate(rw, "Test", ArtisteRecherche)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		http.Redirect(rw, r, "/", http.StatusSeeOther)
	}
}
func carouselle(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./WebAssets/*.html")
	if err != nil {
		fmt.Println(err)
	}
	t := Back.GetApiArtists()
	err = tmpl.ExecuteTemplate(rw, "Carouselle", t)
	if err != nil {
		fmt.Println(err)
	}
}
