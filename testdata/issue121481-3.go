// ./a31954.go:1: doasm: notfound ft=9 tt=16 00054 (/usr/local/google/home/mosoi/source/gocode/src/github.com/brtzsnr/gostress/driver/a31954.go:12)	XORW	$4402880119000, R8 9 16
package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() uint {
switch {} // prevent inlining
v1 := int8(0) // int8
v2 := v1 << 0 // int8
v3 := v2 // int8
v4 := uint16(3 ^ 2 | 3 | 3 - 0) << 1 >> (0 * (0 | 0) ^ (0 ^ 0) >> (2 | 0 | 2)) | uint16(int64(2 ^ 0 & 2) << (uint32(2) << 0 + 3 * 3) << 2) // uint16
v5 := (v4 >> v4) ^ v4 << v4 * 0 | (v4 | v4 << 0) & v4 ^ v4 * v4 * v4 + (v4 ^ v4) + uint16(0) >> 0 * v4 >> 0 & uint16(3) >> 1 ^ v4 << v4 | v4 - 1 ^ (v4 << 0) ^ (v4 * v4) ^ uint16(v3 * v1) ^ 0 & uint16(2) << (3 - 3) // uint16
v6 := v5 | (v4 >> 0 ^ uint16(1) ^ 0 - v4) & v5 + (v4 + v4) + v5 // uint16
v7 := v6 & v6 ^ v6 >> 0 - uint16(0) << 3 << 1 * uint16(uint8(3) >> 3) * v5 + v5 & v5 - v6 * 3 - v4 * v5 << v6 - v5 + v5 << 1 + v5 ^ v4 & v6 << (uint(1) << 0) // uint16
v8 := v3 << 0 >> (v7 ^ v5) >> (v6 - v4 & 3) ^ v2 | (v1 + v2 & v3) >> (uint32(0 & 1 * 2 | 3) >> (3 ^ v4 >> (3 * 2))) & (v1 ^ v1) | v1 // int8
v9 := v2 | v3 ^ v3 - v2 | v8 + v2 * v3 >> (v6 >> 1 ^ v5 & v7 + uint16(int(3) >> 2) - v5 >> v5 << (v7 + v7) ^ 3 - v7 >> (v6 - v5 * v5 * uint16(1) << 0 << (uint8(1) >> 0) & v4)) // int8
v10 := (v9 + v2 + v8 - v1 * v8 >> 0) | v2 ^ v3 ^ v8 * v3 * v3 + int8(0) << 3 | v3 * v9 - v9 + v1 >> 2 | (v9 * v8 * v2 * 2 - int8(1) << (2 | 3)) // int8
_ = v10
v11 := int64(3) << 2 // int64
_ = v11
v12 := v11 << (uint8(3) >> 2 << (1 ^ 3)) ^ 1 - v11 - v11 & v11 * v11 & 0 & v11 * v11 ^ v11 ^ (v11 - v11 >> 0) + v11 >> (uint64(2) >> (2 + 1)) * v11 | v11 + v11 >> (v5 * 0 & v4) // int64
_ = v12
v13 := int(0 | 1 ^ 2 - 0) << (2 ^ 3 + 0 - 1 | 0 << (v7 - v6)) | 2 | 0 + int(2) << 3 * 3 * 1 >> 2 // int
_ = v13
v14 := v12 ^ v11 << 1 + v11 ^ 3 + 0 - v12 << uint8(uint(1) >> 1) & v11 + ((v12 << 0 * 2 + v12) & (v12 << 0 ^ v12 * v12)) // int64
_ = v14
v5++
v6++
v15 := int(0 | 2 + 1 & 0 * 0 ^ 2 | 3 >> 1 ^ 3 * 3 | 2 ^ 1 >> (uint64(1) >> (v6 << 1) | 0 & 0 + 2 - 0 & 0 ^ uint64(3) >> 2 ^ (uint64(2) << 0 & 3 * 3 + uint64(3) << (2 + 0)))) // int
_ = v15
v4++
v2--
return (uint(0 - 0 * 2 | 2 | 3) >> (v5 + v4 ^ v7 + v7 + v7))
}
