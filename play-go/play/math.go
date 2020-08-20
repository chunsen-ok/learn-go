package play

// Pos 表示三维空间中的一点
type Pos struct {
	X float64
	Y float64
	Y float64
}

// NewPos Pos构造函数
func NewPos(x float64, y float64, z float64) Pos {
	return Pos {
		X: x,
		Y: y,
		Z: x,
	}
}
