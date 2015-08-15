// gostress -seed 1447328427 -want 210
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 9, 4, 1, 6, ); got != 210 {
fmt.Printf("f1_ssa(9, 9, 4, 1, 6, ) = %v, wanted 210\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, a3 int, a4 int64, a5 int8, ) uint8 {
switch {} // prevent inlining
a1 = (a1 * a1 - a1 - a1 * a1 & ((a1 >> a1) >> (0 << 3)) << a1) >> ((a2 * a2 + a2) << a1) // uint
a2 = a2 - (a2 + a2) & a2 | a2 ^ a2 | a2 // uint64
a4 *= (a4 - a4 ^ a4 - 0 * a4 >> a1 + a4 + a4 | a4 + a4 | a4 | a4 << 2) >> (0 << a2 | 3 * 3 - 3 | 0 + 3 + 2 + 0 - 0 * 2) // int64
a1 += ((a1 >> 2) & a1 & a1 | a1 >> a2) >> (a2 >> 2) ^ a1 & a1 ^ a1 >> 3 ^ 1 + a1 + a1 << 1 | a1 // uint
a3 -= int(a1 & uint(1 & 2)) << a1 // int
a3++
v1 := 3 | 1 * 3 - 0 + 1 & uint8(2 & 3) >> 3 * uint8(2) << 3 * 2 * 0 & 3 & 1 // uint8
_ = v1
v1--
v1 -= ((v1 * v1 + v1) << (uint16(3) << (a1 ^ a1)) | (v1 ^ v1)) // uint8
return v1 - v1 | v1 >> 3 - v1 << a1 ^ (v1 & v1 + v1 & v1) ^ (v1 | v1 | v1) >> ((v1 + v1) >> (uint32(0) << 3))
}
