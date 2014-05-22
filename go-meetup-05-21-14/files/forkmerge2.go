  mergedResults := make(map[string]string) 
  quit := make(chan bool)  
  errs := make([]error, 0)
  
  go func() { // The merger goroutine // HL
    for {
      select {
      case results := <-c:
        for k, v := range results {
          mergedResults[k] = v
        }        
        wg.Done() // HL
      case err := <-errorsChan:
        errs = append(errs, err)
        wg.Done()
      case <-quit:
        return
      }
    }
  }()

  ... Do stuff here while merger runs // HL

  wg.Wait() // HL
  quit <- true // Throw away the merger routine // HL