func HigherOrderFunc(data interface{}, 
                     phase1 func(interface{}) interface{}, 
                     phase2 func(interface{}) interface{}) interface{} {
  return phase2(phase1(data))
}

---------

type Part func(interface{}) interface{}
type Phase1 Part
type Phase2 Part

func HigherOrderFunc(data interface{}, phase1 Phase1, phase2 Phase2) interface{} { // HL
  return phase2(phase1(data)) // HL
} // HL