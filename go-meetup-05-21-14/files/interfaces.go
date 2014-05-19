type Reader interface {
  Read(b []byte) (n int, err os.Error)
}

type Closer interface {  
  Close()
}

type ReadCloser interface { // Both a Reader and a Closer
  Read(b []byte) (n int, err os.Error)
  Close()
}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err os.Error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    r.Close()
    return
}