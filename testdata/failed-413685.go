// gostress -seed 413685 -want 46179488366604
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 4, 3, 7, ); got != 46179488366604 {
fmt.Printf("f1_ssa(1, 4, 3, 7, ) = %v, wanted 46179488366604\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, a3 uint64, a4 int, ) uint64 {
switch {} // prevent inlining
a1 = a1 ^ a1 | a1 * a1 >> 3 | a1 + a1 | a1 | a1 ^ a1 >> a1 | a1 ^ a1 * a1 | a1 - a1 * a1 | a1 // uint
a1 = a1 // uint
v1 := uint8(1) // uint8
_ = v1
a3 = (uint64(3) >> 2 * a3 | a2) >> ((v1 << a1) >> (a2 * a3)) + a3 << a2 - a2 >> (a1 + a1 - uint(1) * a1 & a1 - a1) // uint64
v1 += 0 - v1 & (v1 | v1 | 3 + v1 * v1 & v1 << a1) >> a1 // uint8
a2 = a2 // uint64
a3 *= (a2 | a2 & 1 * a3 + a3 ^ a3 - a2 + a2) << (a1 + a1 * a1 ^ a1 & a1 + a1 << v1 * a1 ^ a1 | a1 & a1 & a1) // uint64
a1--
a3 = a3 // uint64
return a2 + a2 + a3 << 3 ^ a2 ^ a2 + a2 & a2 - a3 << (a1 ^ a1 & 3)
}
