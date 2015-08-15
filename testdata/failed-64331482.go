// gostress -seed 64331482 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, ); got != 0 {
fmt.Printf("f1_ssa(4, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, ) uint64 {
switch {} // prevent inlining
a1 = a1 // uint64
a1--
a1 *= (a1 >> (0 * uint8(2 ^ 2 + 1) >> (a1 << 3 - a1))) >> (uint32(2 | 3 | 1 ^ 3) << 2 | (uint32(2) >> (0 | 3) ^ 2 & 2 * 2) >> a1) // uint64
a1 -= ((a1 ^ a1 & a1 | a1 ^ a1 >> a1) << (uint16(2 * 1 ^ 2 ^ 1) << ((3 << 0) << a1))) << ((uint(3 * 2 | 3) | (uint(1) >> 2 - uint(a1)) << 0) << (a1 | a1 & a1 ^ a1 * a1 & a1 * a1 | a1 - a1 | a1 & a1 << 3)) // uint64
a1 = (a1 * (((a1 * a1) >> (uint16(1) >> a1)) >> (3 & 3 & 2)) >> (3 + 1 + uint32(2) << 3)) >> (((uint32(0 & 2) >> (uint32(1) << 1 & 3 & 1)) >> 1) << a1) // uint64
a1 *= a1 // uint64
a1 += a1 & a1 >> 3 & ((a1 + a1 * a1 * a1 & a1 * a1 * a1) << 0) >> 2 // uint64
v1 := 0 * (int(0 - 1) >> (uint16(2) << 0)) >> ((uint32(0) >> 2) << 2 - 0 ^ 2) & (0 - 3 - int(1) << 0 & 3 - 0) & 0 | 1 & 1 ^ 0 & 0 - int(0 & 0 ^ 2) >> (a1 * a1 ^ a1 >> 2) // int
_ = v1
return ((a1 ^ a1 ^ a1) | a1 * a1 * a1 & a1 ^ a1 - a1 + a1 + a1 & 2 + a1 | a1 >> a1 - a1 + a1 + (a1 ^ a1) << (1 * 0) * a1) << 0
}
