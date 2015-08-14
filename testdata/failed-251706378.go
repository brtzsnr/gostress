// gostress -seed 251706378 -want 24
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 5, 9, 9, 2, ); got != 24 {
fmt.Printf("f1_ssa(8, 5, 9, 9, 2, ) = %v, wanted 24\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 uint64, a4 int8, a5 uint8, ) int {
switch {} // prevent inlining
a1 += a1 & a1 - a1 * a1 + 2 + a1 & a1 | a1 - a1 // uint
a3 *= ((a3 & a3 & a3 ^ a3) >> (a5 * a5 + a2 - (a2 | a2) << (a3 >> 0))) << (uint32(0 + 0) >> (a3 | a3 * a3 & 1)) // uint64
a5 = a5 | a5 - a2 - a5 | a2 * a5 & a5 * a2 * a5 << (a5 >> 1 & a5 ^ a2) // uint8
a1 += ((a1 << ((a1 | a1) >> (2 + 3))) >> 2) << (a3 & 2 - a3 | a3 & (a3 * a3) >> a5 - a3 >> (0 + 2 + 1 * 0)) // uint
a1++
a2--
v1 := a2 & a2 << a3 | (a2 | a5 - a2 - a5 >> (a1 - a1)) // uint8
_ = v1
return (3 * (int(1) >> a3) << (3 >> a3) + int(1) >> (2 << a3))
}
