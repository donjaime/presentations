func DoStuffInParallel(things []Foo) map[string]string {
  c := make(chan map[string]string)

  for _, thing := range things {
    go func(thing) { // HL
      c <- doWorkOnSingleThing(thing)
    }(thing) // HL
  }
  
  mergedResults := make(map[string]string)  
  for i := 0; i < len(things); i++ {
    results := <-c
    // Merge them into allRoots
    for k, v := range results {
      mergedResults[k] = v
    }
  }

  return mergedResults
}