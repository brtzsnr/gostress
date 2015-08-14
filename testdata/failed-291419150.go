// gostress -seed 291419150 -want 4096
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 9, 2, 4, 0, ); got != 4096 {
fmt.Printf("f1_ssa(1, 9, 2, 4, 0, ) = %v, wanted 4096\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int64, a3 int, a4 uint, a5 uint, ) uint {
switch {} // prevent inlining
a2--
a1 = (a1 >> 3 - a1) >> (2 * 3 | 2) & a1 | a1 // uint8
a2 += a2 & a2 - a2 ^ a2 | a2 + a2 // int64
v1 := a5 // uint
_ = v1
v1 -= v1 // uint
a1 *= (a1 | a1 * a1 * a1 * (a1 * a1) >> (1 ^ 0)) // uint8
a5 *= v1 // uint
v1 *= (1 - v1 + 0 | v1) >> v1 // uint
v1++
return ((v1 + v1) ^ a5 & v1 | (v1 >> a1) >> (1 + 3) * 3 + a5 * v1 * a4) << (1 - 0 + 0 ^ (uint64(3) >> v1) << (2 | 3) ^ uint64(1) << a1 * uint64(2) >> 1 | (0 ^ 0))
}
