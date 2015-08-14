// gostress -seed 701124203 -want 18446744073709551607
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 4, 7, 0, 6, ); got != 18446744073709551607 {
fmt.Printf("f1_ssa(8, 4, 7, 0, 6, ) = %v, wanted 18446744073709551607\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int, a3 uint64, a4 uint, a5 uint, ) uint {
switch {} // prevent inlining
a5 = (a1 * a5 - a4) << (a3 << a5 | a3 >> a4) ^ a4 << a1 - a5 - a5 - 0 - a5 ^ (a4 * a4 << a3) & (a4 | a1 | a4 & a4) // uint
a1++
a3 = a3 << 2 // uint64
a4 = a4 - ((a1 ^ a1) * a1 - a4) >> a3 // uint
v1 := uint64(1) // uint64
_ = v1
v1 = (v1 << (uint16(2) << (2 - 1) * 0 - 0 | 2)) >> ((uint16(3 & 1 + 1) << 0) >> 2) // uint64
v2 := 3 + 0 - int8(1) << (1 ^ 1) | 0 // int8
_ = v2
a2 += a2 // int
v3 := a2 & (a2 - a2) >> 0 * (a2 - a2) | a2 // int
_ = v3
a2++
return a1 + a5
}
