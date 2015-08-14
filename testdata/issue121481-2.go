// ./a31954.go:1: doasm: notfound ft=9 tt=16 00074 (/usr/local/google/home/mosoi/source/gocode/src/github.com/brtzsnr/gostress/driver/a31954.go:11)	ORL	$-72057592695750653, R13 9 16
package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() int {
switch {} // prevent inlining
v1 := uint32(2) << (3 * 3 * 3 - 1 * 0 & uint16(0) >> (1 * 0) << (uint(2) >> 1 | 1 + 2 & 1 * 3)) // uint32
v2 := v1 >> 1 << (1 & v1 | v1 | v1) * (v1 & 0 + v1) - v1 * (uint32(1) >> v1) ^ v1 ^ 3 + 1 * v1 + v1 - v1 * v1 + v1 + v1 // uint32
v3 := v2 // uint32
v4 := v3 * v2 + v1 * v1 >> (1 * 1) ^ v3 & 2 | v3 ^ v1 >> (uint16(2 - 0 & 0 - 0) << (2 - 2)) | v2 // uint32
v5 := v4 * uint32(int64(0) >> 3) >> 1 << (uint16(1) - 1 + 3 ^ 3 ^ 3) * v1 * v3 | v4 >> 0 * v1 // uint32
v6 := (v5 * v5 & v4 >> (uint64(1) >> 1 & 0)) << (1 + 1) >> (3 - 3 * 1 & 0 * 2 + 2 | 3 & uint8(0) >> 3) // uint32
v7 := v6 // uint32
v8 := v7 - v6 // uint32
v9 := int64(2) << (v8 | v8) << (uint(0) >> 3 | 1) << (uint64(2) >> (uint(2) >> 1) * 2 ^ 3 + 3 ^ 0) ^ int64(0 - 2 & 1 ^ 0) >> 0 & 3 // int64
v10 := v9 * v9 >> (uint8(0 & 0) >> (uint16(3) << 3) + 1 * 0 ^ 2 * 0) * v9 // int64
_ = v10
v7++
v7--
v2++
v11 := uint(int8(1 - 0) << 3 & int8(1 - 3 + 3 | 0) << (uint16(1) << 1 + 2)) >> uint64(3) // uint
_ = v11
v12 := uint64(3) // uint64
_ = v12
v2--
return (3 - 1 + int(3) << (v6 | v8 - v1 ^ v6) * 3 ^ int(1) >> v5 ^ 3 - 3) * 2
}
