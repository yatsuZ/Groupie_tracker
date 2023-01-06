package Back

//objectif trouvé un moyen de gagné de l'argent dés cette anné

//recupuré format json decodé etc CHEC
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// -1 Partie Artists ////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
structure de l'api artists
struct to artists api
*/
type ArtistsApi struct {
	ID           int      `json:"id"` //l'identifiant
	Image        string   `json:"image"`
	Name         string   `json:"name"`         //Nom du Groupe
	Members      []string `json:"members"`      //Membre du groupe
	CreationDate int      `json:"creationDate"` //Date de la creation du groupe
	FirstAlbum   string   `json:"firstAlbum"`   //Date du premier albulm
	Locations    string   `json:"locations"`    //Localisation des concert
	ConcertDates string   `json:"concertDates"` //Date des concert
	Relations    string   `json:"relations"`    //Mets en relation les autres API
}

//click+ ctrl
const urlArtists = "https://groupietrackers.herokuapp.com/api/artists"

/*
function to see if the lib module works with Main
fonction qui permet de voir si le module lib marche avec Main
*/
func Test() {
	print("test")
}

/*
[fr] la fonction remplie la structure  ArtistsApi par l'api artists url = "https://groupietrackers.herokuapp.com/api/artists"
[en] the function fulfills the ArtistsApi structure by the artists api url = "https://groupietrackers.herokuapp.com/api/artists"
*/
func GetApiArtists() []ArtistsApi {
	/*
		remarque : ma difficulté es que j'ai mis du temps à comprendre qu'il faut []Artists et non Artists
		note: my difficulty is that it took me a long time to understand that it is necessary to []Artists and not Artists
	*/
	var Intermediere []ArtistsApi
	/*
		étape 1 utilisez la requête GET pour transformer l'url sur http.Response et vérifiez qu'il n'y a pas d'erreur
		step 1 use requet GET for transformed the url on http.Response and check that there is no error
	*/
	Resp, err := http.Get(urlArtists)
	if err != nil {
		panic(err)
	}

	/*
		étape 2 nous lisons le HTTP.response.Body et nous l'avons transformé en byte et vérifions qu'il n'y a pas d'erreur
		step 2 we read the HTTP.response.Body and we transformed that on byte and check that there is no error
	*/
	bytes, err := ioutil.ReadAll(Resp.Body)
	if err != nil {
		panic(err)
	}
	/*
		+ dont forgot to close the request that will execute on laste thanks to "defer"
		+ n'oubliez pas de fermer la requête qui s'exécute en dernier grâce à "defer"
	*/
	defer Resp.Body.Close()

	/*
		dernière étape, utilisez json.Unmarshal pour transformer automatiquement les byte en structure la structure Artists[]
		last step use json.Unmarshal for automatically transform byte to struct in question
	*/
	err = json.Unmarshal(bytes, &Intermediere)
	if err != nil {
		panic(err)
	}
	return Intermediere
}

//-2 Partie Location///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
structure de l'api Locations
struct to Locations api
*/
type LocationsApi struct {
	Index []struct {
		ID        int      `json:"id"`        //l'identifiant de Artiste
		Locations []string `json:"locations"` //Localisation des concert
		Dates     string   `json:"dates"`     //Date des concert
	}
}

const urlLocations = "https://groupietrackers.herokuapp.com/api/locations"

/*
[fr] la fonction remplie la structure  LocationsApi par l'api locations url = "https://groupietrackers.herokuapp.com/api/locations"
[en] the function fulfills the LocationsApi structure by the locations api url = "https://groupietrackers.herokuapp.com/api/locations"
*/
func GetApiLocations() LocationsApi {
	var Intermediere LocationsApi
	Resp, err := http.Get(urlLocations)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(Resp.Body)
	if err != nil {
		panic(err)
	}
	defer Resp.Body.Close()
	err = json.Unmarshal(bytes, &Intermediere)
	if err != nil {
		panic(err)
	}
	return Intermediere
}

//-3 Partie Dates///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
structure de l'api Dates
struct to Dates api
*/
type DatesApi struct {
	Index []struct {
		ID    int      `json:"id"`    //l'identifiant de Artiste
		Dates []string `json:"dates"` //dates des concert
	}
}

const urlDates = "https://groupietrackers.herokuapp.com/api/dates"

/*
[fr] la fonction remplie la structure  DatesApi par l'api Dates url = "https://groupietrackers.herokuapp.com/api/locations"
[en] the function fulfills the DatesApi structure by the Dates api url = "https://groupietrackers.herokuapp.com/api/locations"
*/
func GetApiDates() DatesApi {
	var Intermediere DatesApi
	Resp, err := http.Get(urlDates)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(Resp.Body)
	if err != nil {
		panic(err)
	}
	defer Resp.Body.Close()
	err = json.Unmarshal(bytes, &Intermediere)
	if err != nil {
		panic(err)
	}
	return Intermediere
}

//-4 Partie Realations///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
structure de l'api Dates
struct to Dates api
*/
type RelationApi struct {
	Index []struct {
		ID             int                 `json:"id"`             //l'identifiant de Artiste
		DatesLocations map[string][]string `json:"datesLocations"` //dates des concert avec leur localisation
	}
}

const urlRelations = "https://groupietrackers.herokuapp.com/api/relation"

/*
[fr] la fonction remplie la structure  RelationApi par l'api locations url = "https://groupietrackers.herokuapp.com/api/relation"
[en] the function fulfills the RelationApi structure by the locations api url = "https://groupietrackers.herokuapp.com/api/relation"
*/
func GetApiRelation() RelationApi {
	var Intermediere RelationApi
	Resp, err := http.Get(urlRelations)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(Resp.Body)
	if err != nil {
		panic(err)
	}
	defer Resp.Body.Close()
	err = json.Unmarshal(bytes, &Intermediere)
	if err != nil {
		panic(err)
	}
	return Intermediere
}


/*Permé de trie les nom de groupe qui corespond à BarreDeRecherche je peut l'amelioré en enlenvant les majuscul 
et etre plus permissife pour la recherche*/
func Chercher(L []ArtistsApi, BarreDeRecherche string) []ArtistsApi {
	fmt.Println(L)
	var ArtisteRecherche []ArtistsApi
	fmt.Println(ArtisteRecherche)
	for i := range L {
		for j := range L[i].Members{
			if L[i].Members[j] == BarreDeRecherche{
				ArtisteRecherche = append(ArtisteRecherche, L[i])
				break
			}
		}
		if L[i].Name == BarreDeRecherche  {
			ArtisteRecherche = append(ArtisteRecherche, L[i])
			break//je break car je cherche une chose précis
		}
	}
	fmt.Println(ArtisteRecherche)
	if ArtisteRecherche == nil {
		return L
	} else {
		return ArtisteRecherche
	}
}
