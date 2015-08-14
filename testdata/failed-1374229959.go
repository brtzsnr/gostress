// gostress -seed 1374229959 -want 66
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 7, 2, 4, 1, ); got != 66 {
fmt.Printf("f1_ssa(1, 7, 2, 4, 1, ) = %v, wanted 66\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int8, a3 uint64, a4 int64, a5 int, ) uint8 {
switch {} // prevent inlining
v1 := 2 | 2 + 2 + 1 & 1 | 1 ^ 2 * 3 & 0 * 0 + 2 & 2 - 3 - 0 - 2 | uint(0) << 0 ^ 3 // uint
_ = v1
v1++
a1++
a1 += ((a1 & uint8(v1) << (a1 | a1)) << ((a1 + a1) >> uint8(a3) | (a1 & a1) >> 0)) << (3 * uint32(0) << a3 - uint32(3) << v1 * uint32(1 * 3) >> (v1 << (a3 >> 0))) // uint8
v1 -= v1 * v1 - v1 & v1 | v1 - v1 ^ v1 // uint
v1--
v1--
a2 += a2 // int8
v1 += (v1 - v1 - v1 * v1 & v1) // uint
v1--
return a1
}
