// gostress -seed 27988553 -want 38091
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 6, 0, 8, 4, ); got != 38091 {
fmt.Printf("f1_ssa(9, 6, 0, 8, 4, ) = %v, wanted 38091\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int64, a3 uint8, a4 uint, a5 int8, ) uint {
switch {} // prevent inlining
a1--
a1 *= (int(a2 ^ a2 + a2) >> 0) << (uint16(2 | 3 * 2 & 0) << 0) // int
a4 = a4 * a4 ^ a4 ^ a4 & (a4 * a4) * a4 << 3 - a4 + a4 * a4 & a4 // uint
a5--
a5++
a4 += (1 ^ a4 & (0 ^ a4 << 2) << (0 & 1 ^ uint64(1) << a4)) // uint
a1++
a4 += a4 & a4 << ((a4 ^ a4) >> (a3 ^ a3)) | a4 + a4 & a4 // uint
a5++
return a4 - a4 ^ a4 & a4 * a4 | a4 + ((a4 - 2 + a4) << a3) << (uint(uint8(a2)) ^ a4 & a4 + a4 << a4)
}
