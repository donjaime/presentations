package lib

type LotsOfStuff interface  {
  A()
  B()
  C()
  D()
  E()
  ...
}
------------
package mycode

type ThePartsICareAbout interface {
  A() // HL
  C() // HL
}
