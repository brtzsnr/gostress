// gostress -seed 203891762 -want 32
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, ); got != 32 {
fmt.Printf("f1_ssa(9, ) = %v, wanted 32\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) uint8 {
switch {} // prevent inlining
v1 := (2 | 2 - 3 - 0 | int(2 | 1) >> (uint64(3) >> 0) - 1) >> a1 // int
_ = v1
a1++
v1--
v1 += (v1 >> (uint16(0) >> ((a1 & a1) << 1))) << (a1 + a1 | a1 - (a1 ^ a1 ^ a1) << (2 - 1)) // int
a1 -= (a1 & (a1 + a1 << a1) >> (a1 << 2)) >> (uint32(3) << (uint64(1) - uint64(0) << (2 * 3))) // uint
v1 *= 2 & v1 + (v1 ^ v1 * v1) << (1 - 3 * uint64(3) << a1) // int
v2 := v1 // int
_ = v2
v2 -= v2 // int
a1 = uint((v1 | v1) << (3 & 3)) << (uint32(1) >> 1) & (a1 << 3) >> 3 ^ (a1 + a1) >> a1 & ((a1 * a1) >> (1 ^ 1)) >> 3 // uint
return (uint8(3 ^ 2) << (0 + 3) ^ uint8(1) << (a1 | a1) ^ uint8(2) >> (1 + 3)) << (2 - 0 ^ 0 + 1 & uint64(2) | 2)
}
