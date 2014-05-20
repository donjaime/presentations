func doWithTimeout(work func(), timeout time.Duration) error {
  done := make(chan bool, 1) // Non Blocking... buffered
  go func() {
      work()
      done <- true
  }()
  select {
  case <-done:
    return nil
  case <-time.After(timeout): // HL
    return ErrTimeout
  } 
}

// You would call it above via:
err := doWithTimeout(func() {
    ... // Do some long running work. // HL
}, 3 * time.Millisecond)