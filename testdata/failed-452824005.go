// gostress -seed 452824005 -want -3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 6, 3, ); got != -3 {
fmt.Printf("f1_ssa(7, 6, 3, ) = %v, wanted -3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 int, ) int {
switch {} // prevent inlining
v1 := a3 // int
_ = v1
v1 *= v1 ^ v1 | v1 ^ v1 >> 2 ^ v1 ^ v1 << a1 | v1 & v1 * int(2 - 3) | (v1 + a3 + v1) << (a2 | a2) ^ v1 // int
v1 = (v1 ^ v1 & a3 & v1) + int(1) << ((2 * 3) >> (uint64(1) << a2)) - v1 // int
a2--
v1 += v1 | a3 | v1 - a3 ^ a3 | v1 & v1 * (v1 ^ a3 & v1 >> 0) >> (3 & uint16(3) << 3) - (v1 ^ v1 >> 0) // int
a2--
a2 += 3 | 2 ^ a2 - a2 & uint8(0) << 1 | 1 + a2 >> a1 | a2 * a2 // uint8
a1 += ((a1 << 2) >> 3 + (a1 - a1) << (3 ^ 3) ^ 3 ^ a1 * a1 * 2) << (a2 >> (0 << 0) + (a2 & a2) >> (3 | 1)) // uint
v1--
return (a3 & (v1 << 2) >> (a2 << a2)) >> (a1 >> uint64(2) * a1 | a1 * a1 + a1) | v1
}
