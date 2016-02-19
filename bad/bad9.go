// gostress -seed 11048
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(2, 3, true, false, 0, 3, ))
}
func f1_ssa(a1 uint, a2 uint64, a3 bool, a4 bool, a5 int8, a6 uint, ) bool {
switch {} // prevent inlining
a1--
v1 := a2 ^ (a2 | (a2 - (a2 * a2))) // uint64
_ = v1
v2 := a4 // bool
_ = v2
a1--
v3 := int64(3 * 3) // int64
_ = v3
a1--
return (uint8(3 - 2) << 2) == (2 + uint8(v3 >> 1))
}
