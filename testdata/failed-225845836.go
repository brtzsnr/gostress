// gostress -seed 225845836 -want 1024
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 9, 8, 9, 1, ); got != 1024 {
fmt.Printf("f1_ssa(1, 9, 8, 9, 1, ) = %v, wanted 1024\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int64, a3 uint64, a4 uint8, a5 int, ) int {
switch {} // prevent inlining
a2 += a2 & ((a2 << 3) >> 3) >> (uint16(0) << a4 - 3) + (a2 & a2 | a2 & a2) << (0 + uint16(0 * 2) << (1 + 2)) // int64
a4 -= a4 ^ a4 << 0 // uint8
a5 *= a1 ^ a5 >> a3 - a5 ^ a5 * (a5 << a3) >> 0 // int
a4--
a3 = ((a3 << (a3 ^ a3)) >> (a3 >> (a4 + a4))) << (((a3 << a4) << 2 ^ a3 >> 2 & a3) << 0) // uint64
a3 += (a3 >> 0 - a3 & a3) >> 0 * a3 | a3 << 0 - a3 & a3 - a3 - a3 ^ a3 * a3 ^ a3 ^ (a3 << 1) << (2 - 1) // uint64
a1 *= (a5 >> (uint16(2) >> a4)) << (2 | uint((3 - 2)) >> (a3 ^ a3 | a3 >> 0)) // int
a3 *= a3 - a3 ^ (a3 + a3) * a3 & a3 - uint64(2) + a3 >> 1 + a3 << 2 & a3 >> (a3 - a3) - a3 << (a4 >> 2) // uint64
return ((a5 & a1) ^ a1) >> (3 - uint32(0 + 2) >> (a3 << 1))
}
