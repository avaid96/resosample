package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/gorilla/mux"
	"strings"
	"strconv"
)

const streamString = "stream"

func saveList(S []int, name string) {
	fo, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	e := json.NewEncoder(fo)
	if err := e.Encode(S); err != nil {
		panic(err)
	}
}

func loadList(name string) []int {
	var S []int
	fo, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	e := json.NewDecoder(fo)
	if err := e.Decode(&S); err != nil {
		panic(err)
	}
	return S
}

func reservoirSample(S []int, R *[]int) {
	// caching length in k to prevent double call
	k := len(*R)-1
	// filling the reservoir array
	for index:=0; index<=k; index++ {
		(*R)[index] = S[index]
	}
	// replacing elements with gradually increasing probability
	for index:=k+1; index<len(S); index++ {
		randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
		j := randGen.Intn(index+1)
		if j<=k {
			(*R)[j] = S[index]
		}
	}
}

func startSession(w http.ResponseWriter, r *http.Request) {
	// understanding parameters
	vars := mux.Vars(r)
	tempList := strings.Split(vars["list"], ",")
	var startList []int
	for _, i := range tempList {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		startList = append(startList, j)
	}
	fmt.Fprintf(w, "started session : %s\n", vars["session"])
	log.Printf("started session : %s\n", vars["session"])
	saveList(startList, vars["session"]+streamString)
	w.WriteHeader(http.StatusOK)
	log.Println("started session and loaded list")
}

func displace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tempVar, err := strconv.Atoi(vars["var"])
	if _, err := os.Stat(vars["session"]+streamString); os.IsNotExist(err) {
		fmt.Fprintln(w, "No such session active")
		return
	}
	if err!=nil {
		panic(err)
	}
	fmt.Fprintf(w, "giving %d a chance at displacement\n", tempVar)
	R := loadList(vars["session"]+streamString)
	S := append(R, tempVar)
	reservoirSample(S, &R)
	saveList(R, vars["session"]+streamString)
	log.Println("list is now = ", R)
}

func closeSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := os.Stat(vars["session"]+streamString); os.IsNotExist(err) {
		fmt.Fprintln(w, "No such session active")
		return
	}
	// just displaying the result
	fmt.Fprintln(w, loadList(vars["session"]+streamString))
	log.Println(loadList(vars["session"]+streamString))
	os.Remove(vars["session"]+streamString)
	w.WriteHeader(http.StatusOK)
	log.Println("closed session")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/start/{session}/{list}", startSession).Methods("POST")
	router.HandleFunc("/displace/{session}/{var}", displace).Methods("POST")
	router.HandleFunc("/close/{session}", closeSession).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
