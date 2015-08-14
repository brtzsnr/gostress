// gostress -seed 1084410279 -want 65
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 8, 7, ); got != 65 {
fmt.Printf("f1_ssa(2, 8, 7, ) = %v, wanted 65\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int64, a3 uint, ) int64 {
switch {} // prevent inlining
a2 -= a2 ^ a2 ^ a2 ^ a2 + a2 << a1 * a2 >> a1 - a2 << a1 & a2 & a2 + a2 >> (2 & uint64(1) >> 1) + (a2 ^ a2 & a2) >> (0 & 0) // int64
a3 = a3 // uint
a3 *= a3 // uint
a1 = (uint(a2 + a2 ^ 2) ^ a3 + 3) << (uint32((0 ^ 2)) & uint32(3 + 2 ^ 0) << (2 & 0 | 2)) // uint
v1 := a2 ^ (a2 & a2) - a2 // int64
_ = v1
a1++
v1++
a3++
a1 += a1 // uint
return v1
}
