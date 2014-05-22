
func main() {
        http.Handle("/other", http.HandlerFunc(handle)) // HL
        log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w ResponseWriter, r *Request) { // HL
        fmt.Fprintf(w, "Easiest webserver ever.... got even easier") // HL
}

--------
// Enabled by this type alias in the http package
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
}