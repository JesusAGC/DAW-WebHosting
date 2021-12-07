package api

import (
	"BackendOrdinario/database"
	"BackendOrdinario/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	requestBody, _ := ioutil.ReadAll(r.Body)
	var Lg models.Login
	json.Unmarshal(requestBody, &Lg)

	loginU, _ := database.GetLoginUserData(Lg.UserName)
	loginA, _ := database.GetLoginAdminData(Lg.UserName)

	if loginU.UserName == Lg.UserName {
		if loginU.Password == Lg.Password {
			session, _ := store.Get(r, "session")
			session.Values["username"] = Lg.UserName
			session.Values["usertype"] = "user"
			session.Save(r, w)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("User Login")

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect Password")

		}

	} else if loginA.UserName == Lg.UserName {
		if loginA.Password == Lg.Password {
			session, _ := store.Get(r, "session")
			session.Values["username"] = Lg.UserName
			session.Values["usertype"] = "admin"
			session.Save(r, w)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Admin Login")

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect Password")

		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not exists")
	}

}

func TestSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Allowing CORS to any server requests.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Specifying HTTP methods allowed.
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		return
	}
	username, ok := untyped.(string)
	if !ok {
		return
	}
	untypedUtype, ok := session.Values["usertype"]
	if !ok {
		return
	}
	utype, ok := untypedUtype.(string)
	if !ok {
		return
	}

	//w.Write([]byte(username))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(username + "\n")
	json.NewEncoder(w).Encode(utype)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	// removing session
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		fmt.Println("failed to delete session", err)
		return
	}
	fmt.Println("logged out")
	//http.Redirect(w, r, "http://localhost:1000/compania", http.StatusFound)
}

func IsAdmin(r *http.Request, session *sessions.Session) bool {
	untyped, ok := session.Values["usertype"]
	if !ok {
		return false
	}
	usertype, ok := untyped.(string)
	if !ok {
		return false
	}
	if usertype == "admin" {
		return true
	} else {
		return false
	}
}
