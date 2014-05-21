type Reader interface {
  Read(b []byte) (n int, err os.Error)
}

type Closer interface {  
  Close()
}

// Both a Reader and a Closer // HL
type ReadCloser interface {
  Read(b []byte) (n int, err os.Error)
  Close()
}
