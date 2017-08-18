package matekasse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const helptext string = "Welcome to Matekasse!\nCheck http://wiki.c3re.de/index.php?title=Matekasse-API for more information."

func startServer() {
	http.HandleFunc("/", route)
	err := http.ListenAndServe(listen, nil)
	ce(err)
}

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL.Path)
	if res, _ := regexp.MatchString(`^/get/\d+$`, r.URL.Path); res {
		get(w, r)
	} else if res, _ = regexp.MatchString(`^/set/\d+/(\+|-)\d+$`, r.URL.Path); res {
		set(w, r)
	} else if res, _ = regexp.MatchString(`^/getallusers$`, r.URL.Path); res {
		handleGetAllUsers(w, r)
	} else if res, _ = regexp.MatchString(`^/sum$`, r.URL.Path); res {
		handleSum(w, r)
	} else {
		sendHelp(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	// This should give us "get" and a number
	id, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	respond(ID(id), w, r)
}

func set(w http.ResponseWriter, r *http.Request) {
	// should contain "set", "id" and "+-int"
	arr := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(arr[2])
	amount, _ := strconv.Atoi(arr[3])
	executeBooking(ID(id), amount)
	respond(ID(id), w, r)
}

func sendHelp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, helptext)
}

func respond(i ID, w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getUser(ID(i)))
	ce(err)
	fmt.Fprintf(w, string(json))
}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(getAllUsers())
	ce(err)
	fmt.Fprintf(w, string(json))
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"sum\":%d}", getSum())
}
