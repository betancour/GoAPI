package main
 
import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
 
func main() {
 
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)
 
    log.Fatal(http.ListenAndServe(":3000", router))
}
 
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
 
func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}
 
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoID := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoID)
}