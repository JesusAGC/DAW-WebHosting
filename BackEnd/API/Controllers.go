package api

import (
	database "BackendOrdinario/database"
	"BackendOrdinario/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func MyRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Suggestions/In", GetAllInternSuggestions).Methods("GET")
	router.HandleFunc("/Suggestions/Ex", GetAllExternSuggestions).Methods("GET")
	router.HandleFunc("/Users/{id}", GetUserByID).Methods("GET")
	router.HandleFunc("/Users", RegisterUser).Methods("POST")
	router.HandleFunc("/Contracts", GetAllContracts).Methods("GET")
	router.HandleFunc("/Contracts/{id}", GetContractByID).Methods("GET")
	router.HandleFunc("/Contracts", RegisterContract).Methods("POST")
	router.HandleFunc("/Contracts/{id}", EditContract).Methods("PUT")
	router.HandleFunc("/Contracts/{id}", RemoveContract).Methods("DELETE")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/loginT", TestSession).Methods("GET")
	router.HandleFunc("/logout", Logout).Methods("GET")
	// router.HandleFunc("/secret", secret).Methods("GET")

	return router
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var Usr models.User
	json.Unmarshal(requestBody, &Usr)
	usrDB, _ := database.GetUserByMail(Usr.Email)

	if usrDB.Email == "" {
		err := database.NewUser(Usr)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode("Error")

		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Success")

		}
	} else {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("Error, ya existe el correo")

	}

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ID := mux.Vars(r)["id"]

	UserDB, err := database.GetUser(ID)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.Marshal(UserDB)
		json.NewEncoder(w).Encode(UserDB)

	} else {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Error al registrar contrato")

	}

}

func GetAllInternSuggestions(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	suggs, err := database.GetAllInSuggestions()
	if err != nil {
		fmt.Println(err)
	}

	json.Marshal(suggs)
	json.NewEncoder(w).Encode(suggs)
}

func GetAllExternSuggestions(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	suggs, err := database.GetAllExSuggestions()
	if err != nil {
		fmt.Println(err)
	}

	json.Marshal(suggs)
	json.NewEncoder(w).Encode(suggs)
}

func GetAllContracts(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	contrs, err := database.GetAllContractsFromDB()
	if err != nil {
		fmt.Println(err)
	}

	json.Marshal(contrs)
	json.NewEncoder(w).Encode(contrs)
}

func RegisterContract(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var Ctr models.Contract
	json.Unmarshal(requestBody, &Ctr)
	Ctr.DateOfContract = time.Now()
	fmt.Println(Ctr.DateOfContract.String())
	Ctr.DateOfExpiration = Ctr.DateOfContract.AddDate(1, 0, 0)

	err := database.InsertContract(Ctr)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Success in Contract")
	} else {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Error al registrar contrato")

	}
}

func RemoveContract(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ID := mux.Vars(r)["id"]

	err := database.DeleteContract(ID)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Se borró el contrato")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Error al borrar el contrato")
		fmt.Println(err.Error())
	}
}

func GetContractByID(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ID := mux.Vars(r)["id"]

	Ctr, err := database.ReadContract(ID)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.Marshal(Ctr)
		json.NewEncoder(w).Encode(Ctr)

	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Error al registrar contrato")
		fmt.Println(err.Error())

	}

}

func EditContract(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita estar loggeado")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if !IsAdmin(r, session) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Se necesita ser Admin")
		return

	}

	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var Ctr models.Contract
	json.Unmarshal(requestBody, &Ctr)

	err := database.UpdateContract(Ctr)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Se editó el contrato")
	} else {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Error al editar el contrato")
	}
}
