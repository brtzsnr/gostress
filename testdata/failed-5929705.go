// gostress -seed 5929705 -want 5
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 1, 8, 7, ); got != 5 {
fmt.Printf("f1_ssa(0, 1, 8, 7, ) = %v, wanted 5\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, a3 uint8, a4 int64, ) int {
switch {} // prevent inlining
v1 := ((a1 | a1) << (1 ^ 0) | (a1 | a1)) << (2 + 1) ^ a1 // uint64
_ = v1
v2 := 2 // int
_ = v2
v2 *= v2 >> ((a1 + v1 + v1 >> a3) >> 0 | a1) // int
a2++
v3 := a3 | (a3 & 0) ^ a3 | (a3 + a3 ^ a3) >> ((a2 * a2) << 0) // uint8
_ = v3
a2 *= a2 // uint
v3 = (v3 & a3) >> (1 & (uint16(0 * 0) << (0 | 3))) // uint8
v1 += a1 + a1 ^ a1 | a1 & a1 + (v1 - a1 * v1 * v1 & a1) >> (0 + uint32(3 * 1) >> (2 ^ 0)) // uint64
v1 *= ((v1 + v1 << 2 | a1) >> v1) // uint64
a2++
return v2 & (v2 >> v1 - v2 << a3) >> (3 * 0 ^ 1) ^ int(a4 >> (3 & 0))
}
