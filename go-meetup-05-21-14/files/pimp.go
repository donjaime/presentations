package lib

type Foo interface {
  DoSomething()
}

-----------

package mycode

type MoreThanFoo interface {
  lib.Foo // HL
  DoOtherThing()
}

type FooWrapper struct {
  lib.Foo // HL
}

func (f FooWrapper) DoOtherThing() {
  log.Println("Other Thing")
}
