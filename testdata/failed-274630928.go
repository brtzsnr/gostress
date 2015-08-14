// gostress -seed 274630928 -want 32768
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 8, 7, 8, ); got != 32768 {
fmt.Printf("f1_ssa(1, 8, 7, 8, ) = %v, wanted 32768\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, a3 uint, a4 uint, ) int {
switch {} // prevent inlining
a4 -= a4 // uint
a1 = a1 // int
a2++
a3 += ((a3 | 2) >> (a3 - a3)) >> 1 + a3 & a3 | (a4 << (a2 & a2)) >> (a3 * a3 ^ a3 ^ a3) ^ (a3 + a3) >> (0 * 3 ^ uint64(2) >> 0) // uint
a2 += a2 * a2 * a2 ^ a2 << a2 & (a2 | a2) << 0 | a2 // uint8
a3--
a2 += (0 - a2 | a2 | a2) // uint8
v1 := int8(2) >> (uint16(1 ^ 3) >> a4) ^ 2 ^ int8(2 | 3 | 3 | 0) << (0 ^ 0) // int8
_ = v1
a1 -= (a1 ^ a1 & a1 - a1 + a1) << (uint64(0) << 2 + 2) + (0 | a1 | a1 * a1) | a1 ^ a1 ^ a1 * a1 - a1 >> a3 ^ a1 // int
return (a1 - a1) << (1 - 0 & 0) | (a1 >> a3) << (3 | 2) | (a1 & a1) << (a4 + a3) - ((a1 * a1) >> (uint16(0) << a2) ^ a1 << (uint32(0) << 0)) << 0
}
