// gostress -seed 1340615643 -want 2
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 4, ); got != 2 {
fmt.Printf("f1_ssa(2, 4, ) = %v, wanted 2\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint, ) int64 {
switch {} // prevent inlining
a2++
v1 := uint8(1) << (uint8(2) << ((3 + 0) << a2)) // uint8
_ = v1
v2 := (a1 * a1 * a1) << uint32(0) + (a1 << a2) << (2 & 2) & a1 - a1 | a1 ^ a1 ^ (a1 | a1 + a1) << (v1 * v1 ^ v1) ^ a1 * a1 - a1 - a1 >> a2 // int64
_ = v2
a1--
v1 = v1 ^ v1 - 2 * v1 - v1 | v1 & 3 ^ v1 * v1 * v1 + v1 ^ v1 + v1 // uint8
v2 += (a1 & v2) // int64
v2 = a1 ^ v2 // int64
v2--
a1 = v2 // int64
return a1 - v2 & v2 & a1 - v2 << a2 & a1 + a1
}
