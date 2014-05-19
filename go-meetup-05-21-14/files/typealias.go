type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Easiest webserver ever")
}

func main() {
        http.Handle("/", &MyHandler{})
        log.Fatal(http.ListenAndServe(":8080", nil))
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
}