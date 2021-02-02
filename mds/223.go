/*
223. 矩形面积

在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。

每个矩形由其左下顶点和右上顶点坐标表示，如图所示。

Rectangle Area

示例:

输入: -3, 0, 3, 4, 0, -1, 9, 2
输出: 45

说明: 假设矩形面积不会超出 int 的范围。
*/
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if B >= H || C <= E || D <= F || G <= A {
		return abs(C-A)*abs(D-B) + abs(G-E)*abs(H-F)
	}
	s, x, z, y := min(D, H), max(B, F), max(A, E), min(C, G)
	return abs(C-A)*abs(D-B) + abs(G-E)*abs(H-F) - abs(s-x)*abs(z-y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}