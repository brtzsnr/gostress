// gostress -seed 3029318246 -want 14016
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 0, 0, 5, 6, ); got != 14016 {
fmt.Printf("f1_ssa(0, 0, 0, 5, 6, ) = %v, wanted 14016\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int, a3 int64, a4 uint, a5 int64, ) uint {
switch {} // prevent inlining
a4 = a4 & a4 * a4 * a4 & a4 | (a4 & a4 >> a4 + a4 * a4 | a4) >> uint8(uint64(1 * 1) << a4) // uint
a4 *= (a4 | a4 * a4 | a4) >> (a4 + a4 & a4 >> a4) | ((a4 & a4 << 3) << 0) // uint
v1 := a4 << (0 * 2 + 0 & 2 + 3) ^ a4 // uint
_ = v1
a4 -= v1 // uint
a1++
v1 -= a4 * (v1 + a4 * v1 + v1) & (v1 + a4) * (v1 + v1) >> (uint64(3) << 2) // uint
v1 -= v1 | (v1 * v1) << 2 // uint
v2 := ((a1 * a1 - a1 & a1 | a1) << (uint16(0 & 2) << (2 >> 2) * 2 + 3 & uint16(2) << 1)) // uint8
_ = v2
return v1 * v1 - v1 * v1 - a4 + v1 ^ v1 & (a4 << 3)
}
