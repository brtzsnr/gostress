// gostress -seed 666618265 -want 18446744073709551456
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 7, 3, ); got != 18446744073709551456 {
fmt.Printf("f1_ssa(2, 7, 3, ) = %v, wanted 18446744073709551456\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int64, a3 int8, ) uint64 {
switch {} // prevent inlining
a2 += a2 // int64
a1 += a1 // uint8
a3++
a2--
v1 := 2 | uint(0) >> (uint32(2 + 2) >> (3 + 2)) ^ (uint(3 * 0) << ((uint32(3) << 2) << (a1 * a1))) // uint
_ = v1
v1++
v2 := a3 >> a1 // int8
_ = v2
v3 := uint64(2) >> ((v1 * 1) << (0 | 2)) | uint64(0) << (v1 >> 3) ^ 1 ^ 1 - 1 + (uint64(3 ^ 2) << (0 & 2)) << (uint16(v1) + uint16(0) >> a1 + 2 + 0 | 1) // uint64
_ = v3
return 1 * (2 - v3) - v3 | v3 - v3 << 2 ^ 3
}
