type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Easiest webserver ever")
}

func main() {
        http.Handle("/", &MyHandler{}) // HL
        log.Fatal(http.ListenAndServe(":8080", nil))
}

