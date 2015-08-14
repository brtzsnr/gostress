// gostress -seed 180242551 -want 1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, ); got != 1 {
fmt.Printf("f1_ssa(5, ) = %v, wanted 1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) int64 {
switch {} // prevent inlining
a1 = (a1 | a1 ^ a1 ^ a1 + a1 ^ a1) << 0 // uint
a1 *= (a1 | a1 * a1 << 0 + a1 << 3 | a1 & a1 ^ a1 * a1 - a1 & a1) // uint
a1--
a1 *= a1 ^ a1 >> a1 * a1 << 2 | a1 ^ a1 * (a1 * a1 * a1 - 3) >> 0 // uint
a1 = (2 - a1 + (a1 + a1 * a1 - a1) << (uint16(1) << 0 | 2 & 3)) // uint
v1 := (1 | uint64(3) << 3 ^ 0 & 3) >> (a1 * a1 + a1) - 2 - 1 ^ 3 - 3 | 1 & 0 ^ uint64(1) >> 1 // uint64
_ = v1
a1 *= a1 // uint
v2 := int64(2) >> (uint16(3) >> (0 * 3)) - 1 // int64
_ = v2
a1 -= a1 * a1 - a1 | 0 * a1 ^ a1 - a1 & a1 >> v1 // uint
return int64(1) * v2 & v2 * ((v2 >> v1) << (3 & a1))
}
