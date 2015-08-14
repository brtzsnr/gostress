// gostress -seed 201025545 -want 8
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 1, 3, ); got != 8 {
fmt.Printf("f1_ssa(6, 1, 3, ) = %v, wanted 8\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, a3 uint64, ) int8 {
switch {} // prevent inlining
a3 = (a3 - a1 | a3 & uint64(0)) >> 1 + a1 // uint64
v1 := (a1 * a3) >> 3 | ((a1 & a3 & a1 >> a3) >> (a3 - a3)) // uint64
_ = v1
v1++
v2 := a3 // uint64
_ = v2
v3 := 3 | 0 + int8(0) >> 0 * 3 & 2 * int8(3) >> 2 + 2 | 2 & int8(0) << 3 + int8(1 & 0) >> (1 * 0) // int8
_ = v3
a2++
v1 = ((a3 & v2 | v1 << 0) >> 0 - a3 + (v2 - v2) << 2) // uint64
v3 += (3 - v3 & v3 ^ v3 & int8(1)) << 0 * ((v3 << 3) << v1 * v3 >> a2 ^ v3) >> a2 // int8
v1++
v2++
return ((v3 << a2 - v3 + 3 & v3 << 2 * v3 * v3) << 1) << (2 * (0 + 2) >> (uint8(3) >> 1 ^ uint8(1) >> 2))
}
