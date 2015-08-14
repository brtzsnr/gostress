// ./a31954.go:1: doasm: notfound ft=9 tt=16 00044 (/usr/local/google/home/mosoi/source/gocode/src/github.com/brtzsnr/gostress/driver/a31954.go:13)	ORW	$-4396353028, DI 9 16
package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() int {
switch {} // prevent inlining
v1 := int(0) // int
_ = v1
v2 := (uint16(1) << 2 << (uint32(3) << 1) ^ 1 | 2 * 0 * 0) ^ 0 | 0 * 3 // uint16
v3 := v2 | v2 & v2 * v2 // uint16
v4 := v3 // uint16
v5 := v4 ^ v4 - 3 ^ v4 * v4 | v2 << 1 + (v2 - v3) << (2 + 0 | 2 | 0) * uint16(v1 - v1 | 3 << uint(0)) ^ v4 | v3 & v3 << 1 << 3 | v2 << (uint8(3) >> (v4 & v3)) * v3 // uint16
v6 := v5 // uint16
v7 := v6 ^ v2 - v3 & v2 * v5 & v4 * v2 + v5 * v6 << 2 & v2 ^ 3 // uint16
v8 := v7 * (v6 ^ 2 >> (0 | 2) | 3) ^ v3 // uint16
v9 := 2 | v8 // uint16
v10 := int(1 + 1 & 0 << (v9 + v5 - v4) + 3 * 3 << (v2 - v4 + uint16(3) >> v9) << 1 * 0) // int
_ = v10
v7--
v4--
v7++
v7--
v11 := int(3 & 1 * 0 + 3 ^ 1 - 3 * 3 + 1 & 2 - 0 ^ 3 >> 0 + 3 - 2 & 0 ^ 3 * 1 ^ 3 ^ 1 - 3 >> (v7 ^ v2 - v3 - v3 - v8 & v6 >> 0 ^ v4 - v6 | v9)) // int
_ = v11
v8++
v12 := int8((v1 & v11 - v10 + (1 + 1) - v1 * v10 - 2)) // int8
_ = v12
v11++
return v11 * int(1 - 3) << v5 * v1 >> (uint64(2) >> 2) << ((v3 * v4) - v5 ^ 1 >> (2 & 0)) * v1 + v10 * v1 | v1 >> 2 - v10 * v10 - 0 ^ v1 << uint16(v12 & v12 * 2 + 3 ^ v12 * v12 ^ v12)
}
