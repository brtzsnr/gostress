// gostress -seed 182708113 -want -36028797018963968
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 0, 5, 7, ); got != -36028797018963968 {
fmt.Printf("f1_ssa(0, 0, 5, 7, ) = %v, wanted -36028797018963968\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint8, a3 int, a4 uint, ) int64 {
switch {} // prevent inlining
v1 := (2 & 1 - 2 ^ int8(1) << 1 * 2) << a4 // int8
_ = v1
v1 += (v1 + 3) >> (a4 & a4 ^ a4 - a4 & a4 & a4 & (a4 | a4) >> (0 | 1) * (a4 - a4) >> (1 << a2)) // int8
a3 += (a3 | a3 - int(a4)) >> ((0 * 2) >> a2) + a3 - a3 >> 3 ^ a3 >> 1 - (a3 << (3 - 3 & 3)) << 1 // int
v1 += v1 << (2 << a2 ^ 2 | 3) * v1 - v1 ^ v1 >> 3 * (v1 - v1) << 1 * v1 * v1 & v1 | v1 ^ (v1 ^ 2) >> (2 + 3) // int8
v1 *= v1 * (v1 | v1 | v1) << (1 - 1 ^ 0) + v1 | v1 // int8
v1--
v1++
v1 *= v1 // int8
a1 *= a1 & (a1 | a1 + a2 & a2 + a1) << (a2 ^ a1 | a2 - a2 - a2 | 3 | a1) // uint8
return ((1 ^ 1 ^ 0 * 0) | 1 ^ 3 * 0 ^ 1 - int64(1) >> (a2 >> 2)) << (a4 | a4 & a4 * a4 & (a4 * a4))
}
