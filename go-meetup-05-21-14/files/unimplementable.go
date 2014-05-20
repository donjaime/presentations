package foo

type CantImplementMeBro {
  PublicMethod()
  packagePrivateMethod() // Only implementable and callable within the package // HL
}

func DependsOnConcreteType(obj CantImplementMeBro) {
  switch t := obj.(type) {  
  case Struct1:
      fmt.Printf("Struct1 %t\n", t)
  case Struct2:
      fmt.Printf("Struct2 %t\n", t)
  default: // HL
      fmt.Printf("BAD THIS SHOULDN'T HAPPEN! %t\n", t) // HL
  }
}