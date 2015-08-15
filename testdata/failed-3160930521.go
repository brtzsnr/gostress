// gostress -seed 3160930521 -want 30819
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 1, 7, ); got != 30819 {
fmt.Printf("f1_ssa(8, 1, 7, ) = %v, wanted 30819\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int64, a3 uint64, ) int64 {
switch {} // prevent inlining
a1 -= uint((2 ^ 0 * 2) - 1 & 2 ^ int(1) >> a3 - int(2) << (a1 << 3)) * ((a1 >> 1) >> (uint8(1) >> 1) | a1 + a1 ^ a1 * a1 - (a1 | a1) >> (1 * 0) | a1 | a1 * a1 - a1) << (uint16(0) >> (3 & a3 + a3 & (a3 >> 3) << 0)) // uint
a3 += a3 ^ (a3 & a3 + a3 * a3 >> 3 ^ a3 + a3 + a3 ^ a3 * a3 >> 2) << 2 // uint64
a2 += (a2 + a2 & a2 - a2 << (1 & 2) - 0) >> 1 & (a2 * (a2 | a2) >> (a1 * a1 + a1 * a1)) << (((a3 ^ a3) >> 1) << (0 ^ 0 ^ 1 | 2) * a3 - uint64(1) + a3 - a3) // int64
a1--
a1--
a2 *= a2 << (2 + 3 >> (a1 & a1 | a1 * a1) * (2 ^ 3 * 3 | 0) << (a1 & a1 | a1 & a1)) // int64
a2 *= a2 << 3 ^ a2 - a2 & a2 + a2 >> 0 - a2 >> (a3 | a3) & a2 & a2 + a2 << 2 - a2 | a2 >> a1 - a2 | a2 ^ a2 << a3 | a2 | a2 + (a2 ^ a2) | 0 & a2 * a2 // int64
a2 -= (a2 - a2 ^ a2 * a2 * a2 ^ a2 * a2 + a2 | a2 * a2 | a2 & a2) << (uint32(3 | 0) << (3 * 3) - 1) & a2 // int64
a1 += (a1 ^ a1 - a1) & (a1 & a1 | a1 >> a3) >> (2 * 2 & 0 << a3) ^ a1 ^ a1 // uint
return (a2 * a2 - a2 | a2 + a2 + a2 - a2 + (a2 >> a1) << (3 & 3) | 3 ^ a2 & a2 - a2 & a2 & a2)
}
