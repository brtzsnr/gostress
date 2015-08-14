// gostress -seed 42032931 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 9, 2, ); got != 0 {
fmt.Printf("f1_ssa(6, 9, 2, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint64, a3 uint64, ) int8 {
switch {} // prevent inlining
a1--
a1++
v1 := uint(0) // uint
_ = v1
v2 := v1 // uint
_ = v2
v2 *= v2 | (v1 ^ v2) >> (uint(3) << v2) - (v1 & v1 + v1 & v2) & v1 // uint
v2++
a1++
a3 += a3 >> a2 - a2 | a3 | a2 >> a3 ^ a3 + a3 & 1 + a3 >> (uint16(3) >> (v1 | v2 & v1)) // uint64
return ((a1 << 1 | a1 | a1) * a1) >> (a2 + a3 + a3 * a2 >> 1 ^ a2 & a2 | a3 ^ a3 & a3 * a3 >> a3 ^ a3)
}
