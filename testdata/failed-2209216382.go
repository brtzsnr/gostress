// gostress -seed 2209216382
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, ); got != 480 {
fmt.Printf("f1_ssa(4, ) = %v, wanted 480\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, ) uint64 {
switch {} // prevent inlining
a1++
v1 := ((0 & 1) ^ 0 & 1 + 2 ^ uint64(2) << 0 ^ 0 | 0) >> 0 & 0 // uint64
_ = v1
a1--
v1 = ((v1 >> 0 * v1) + v1 << v1 & 3 + v1 | (v1 - v1 + 2) << ((0 >> 0) << 2)) // uint64
v1 = v1 + 1 + v1 + v1 ^ v1 * (v1 ^ v1) >> (v1 & 1) - v1 + v1 + v1 | v1 - v1 & (v1 >> 1) >> (0 + 1) + v1 // uint64
v1++
v1--
v2 := v1 << 1 // uint64
_ = v2
v1++
v2++
v1++
v2 = v2 // uint64
v2--
v3 := ((a1 + a1) << (0 ^ 3) & a1 | a1 ^ a1 - a1 & a1) >> (3 & uint32(2) >> v1 + 0 * 3 & 2) + a1 // int64
_ = v3
a1++
v4 := 0 ^ 1 & int(0) >> ((2 << 1 & 3 << 2 | 1 | 1 | 1) << 3) // int
_ = v4
v5 := (1 * 3 - 0 | 0 + 3 & uint(2) >> (v2 * v1) & 2 | ((uint(2 + 2) << v2) >> (2 + 3 + 3 << 2))) >> ((v2 >> 3 * v2 << 1 + (v2 & v1) ^ v2) >> (0 * 2)) // uint
_ = v5
v2++
v6 := (v4 & v4 - v4 * v4 & v4 | v4 * v4 ^ v4 + v4 + v4 + int(2) << 2 * v4 & v4 & v4 * (v4 >> (v5 >> 0))) >> 0 // int
_ = v6
return v1 ^ (v1 << v5 + v1 * v2) + (v2 >> (uint16(3) >> 1)) << (3 - 1) & v1 * 0 & v2 & v1 ^ (v2 + v1) >> (uint(2) >> v5) * v1 >> 1
}
