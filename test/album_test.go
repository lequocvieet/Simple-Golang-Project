package test

import (
	"fmt"
	"golang_standard/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var err = godotenv.Load("local.env")
var router = mux.NewRouter()

func TestRegister(t *testing.T) {
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	router.Handle("/auth/register", http.HandlerFunc(routes.Register))
	// var user = models.User{
	// 	Username: "new",
	// 	Email:    "asdasd@gmail.com",
	// 	Password: "213423",
	// }
	body := "username=new&email=hsadh@gmail.com&password=12"
	req, err := http.NewRequest("POST", "/auth/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}

}

func TestLogin(t *testing.T) {
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	router.Handle("/auth/login", http.HandlerFunc(routes.Login))
	body := "username=vieet&password=06012000"
	req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestGetAllAlbums(t *testing.T) {
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	router.Handle("/albums", http.HandlerFunc(routes.GetAllAlbums))
	//Token from login
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjQyMTg3ODksInVzZXJuYW1lIjoidmllZXQifQ.68ylT8H41TdGBMUJKKdd8lNBVKqW0OV95ZYlvorz5L0"
	req, err := http.NewRequest("GET", "/albums", nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func executeRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
