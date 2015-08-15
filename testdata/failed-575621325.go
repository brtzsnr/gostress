// gostress -seed 575621325 -want -196635
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, ); got != -196635 {
fmt.Printf("f1_ssa(0, ) = %v, wanted -196635\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, ) int64 {
switch {} // prevent inlining
a1 *= a1 // uint64
a1 -= 2 & a1 | (a1 ^ a1) << (1 << (0 | 3)) * a1 << 3 & a1 + uint64(2) // uint64
a1 -= a1 | a1 + a1 & a1 ^ (a1 >> 1) & a1 | a1 ^ (a1 ^ a1) + a1 * a1 << 0 ^ 0 * a1 * (a1 | a1) >> (a1 - a1) | a1 // uint64
a1 *= a1 | uint64((3 * 0 * 0) + 2 + 3 ^ 3 - 0 - 0) // uint64
a1 = a1 // uint64
v1 := ((a1 >> 3) - (a1 * a1)) >> (uint16(uint8(2) >> 0) & (2 & 3)) - (a1 * a1) << a1 & a1 - a1 ^ a1 - a1 * a1 & a1 * (a1 >> (a1 * a1)) << (uint(3) << a1 & 0 | 0) & uint64(3) << 3 ^ a1 // uint64
_ = v1
a1 -= uint64(int8(0 * 0) << 1 * 0 & 1 * int8(0) << 3 ^ 0 * 3 | 0 ^ 1 + 2 + 3 & 2 - 1 + 3 - 1 + 1 - 2 | 0 ^ 1) // uint64
v1 *= a1 & (v1 - a1 * a1 | v1 & (v1 ^ v1) << (a1 << 1)) >> (0 * 2 + 2) ^ v1 + a1 - v1 & a1 << 2 - a1 - v1 * a1 // uint64
a1 = v1 ^ v1 & a1 * v1 >> 1 - v1 - a1 * a1 ^ a1 ^ v1 + ((a1 * v1) >> (uint16(2) << 3) | (v1 + a1) << v1) >> ((a1 - a1 + a1 << 3) << (uint8(3) >> v1 * 1 | 2)) // uint64
return (3 + int64(3) ^ int64(1) << v1) >> (0 - 1 & 1 & uint32(0) >> 1) - ((int64(2) << 1) << a1 ^ 2 | 3 | int64(2) >> a1) >> (2 | 0 - 0 + uint32(3) << v1 + 2 * 2) - (int64(1 | 0) >> v1 + int64(3 - 0) << (2 << 3)) - 1
}
