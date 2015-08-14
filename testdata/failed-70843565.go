// gostress -seed 70843565 -want -33
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 9, 0, 6, ); got != -33 {
fmt.Printf("f1_ssa(1, 9, 0, 6, ) = %v, wanted -33\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, a3 int, a4 uint64, ) int8 {
switch {} // prevent inlining
a3 -= (a1 >> a2 ^ a1 & a1 + a1 >> a2) // int
a1 = 2 // int
a1 = (a3 >> (1 + 3)) << a2 & a1 << 3 & a1 | a1 >> 2 ^ a3 + a1 | a1 ^ a1 >> 3 - a3 >> 0 ^ a3 & a1 << (a2 ^ a2) // int
a4 -= uint64((1 ^ 0) | 2 + 2 | int64(0) << 1 - 2 + 0 ^ 2 & 2) // uint64
v1 := int8(0) // int8
_ = v1
v1++
a3--
return (v1 & v1) << (a4 ^ a4) + v1 | ((v1 | v1) << (a2 - a2)) >> a4 | v1 ^ v1 - v1 ^ v1 & v1 << (a4 & a4) ^ int8(a3)
}
