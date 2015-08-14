// gostress -seed 808431353 -want -4
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, ); got != -4 {
fmt.Printf("f1_ssa(1, ) = %v, wanted -4\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) int8 {
switch {} // prevent inlining
a1 -= a1 // uint
v1 := uint8(2) << 2 // uint8
_ = v1
a1 = (a1 ^ a1 & a1) & 1 + a1 | a1 * ((a1 - a1) >> a1) >> 1 & a1 + a1 | a1 + a1 & a1 << 3 & a1 // uint
a1 *= a1 // uint
v2 := 2 - 0 ^ 0 // int
_ = v2
a1 -= a1 // uint
v3 := 3 & int8(1 ^ 1) << (3 ^ v1) - int8(3 & 2) << 2 | int8(0) << (0 & 3 + 1) // int8
_ = v3
a1 += a1 >> (a1 >> (a1 & a1 * 0 ^ 3) ^ a1 << 2 * a1 + a1 * (a1 - a1) >> (0 ^ 1)) // uint
v2 = v2 // int
v1 = (v1 * v1 ^ v1 * v1 - v1 & v1) // uint8
return (((int8(1) << a1 & v3 | v3) >> (1 - 0)) >> a1) << ((uint16(2 ^ 2) << (0 >> a1)) + 0)
}
