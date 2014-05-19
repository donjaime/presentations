type Inner struct {
}

func (i *Inner) DoesSomething() {
  ...
}

// Outer "inherits" the methods on Inner
type Outer struct {
  Inner // HL
}
