package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

type Data struct {
	Email string
}
type response struct {
	ValidEmail bool `json:"validEmail"`
}

func momHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Error: Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hi mom")
}
func CheckEmail(w http.ResponseWriter, r *http.Request) {
	var p Data
	if r.Method != "POST" {
		http.Error(w, "Error: Bad Request", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("End Point hit: Post article")
	fmt.Println(p.Email)
	if validMailAddress(p.Email) {
		domainData := strings.SplitN(p.Email, "@", -1)
		fmt.Println(domainData[1])
		hasMX, hasSPF, _, hasDMARC, _ := checkDomain(domainData[1])
		if hasMX && hasDMARC && hasSPF {
			res := response{
				ValidEmail: true,
			}
			fmt.Fprintf(w, "data: %v", res)
		} else {
			res := response{
				ValidEmail: false,
			}
			fmt.Fprintf(w, "data: %v", res)
		}
	} else {
		http.Error(w, "Error: Wrong Email given", http.StatusBadRequest)
		return
	}

}
func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}
func Router() {
	http.HandleFunc("/", momHandle)
	http.HandleFunc("/email", CheckEmail)
	fmt.Println("Server running on PORT: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Router()
}
