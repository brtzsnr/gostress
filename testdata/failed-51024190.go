// gostress -seed 51024190 -want 20
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 1, 9, ); got != 20 {
fmt.Printf("f1_ssa(5, 1, 9, ) = %v, wanted 20\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, a3 uint64, ) uint8 {
switch {} // prevent inlining
v1 := int64(1 + 3 + 2) << (a3 ^ a3 * a3 - a3) + int64(2 * 2) >> a2 * 2 + 2 - int64(2) & 0 * 3 & 2 - 2 // int64
_ = v1
a3++
v1 = ((v1 * int64(2) << a2 & v1) >> (a2 << (uint(1) >> (uint16(1) >> a2)))) >> a2 // int64
v1 += v1 // int64
v1++
v1 *= (v1 << 2 | 0) | v1 * v1 // int64
v1 += (v1 << (3 ^ 1) | v1 | v1 >> 1) >> (3 - 0 - uint16(1) << (0 + a3)) - v1 + (v1 | 1 + v1 | v1) // int64
v1 += ((2 + v1 * v1 + v1) << a2) >> 2 // int64
a3 += a3 // uint64
v1 += ((v1 | v1) >> ((2 | 2) << (1 * 3)) - (v1 + v1 * v1)) << (a3 - (a3 | a3) << (3 | 3) * a3) // int64
return a2 & a2 - a2 & uint8(v1) * a2 << 2 | a2 | a2 ^ uint8(a3) - a2
}
