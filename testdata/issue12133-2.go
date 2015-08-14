package main
import "fmt"
func main() {
fmt.Println(f1())
}
func f1() int64 {
switch {} // prevent inlining
v1 := uint16(7) // uint16
v2 := v1 * uint16(int8(v1)) + 8 // uint16
v3 := v2 // uint16
v4 := v3 // uint16
v5 := v4 // uint16
return 1 * (2 >> 7) - 8 << (v5 * 5 * v2 + v2 + 1 << v5) + int64(v4) * int64(v5) - int64(v5) >> uint(8 >> v1) - int64(v3) - 6 + 8 << uint8(v3) - 3 << uint64(1 << 6 + 5 + 3 * uint(v2) + 6 - uint(v4) << (9 * 8))
}
