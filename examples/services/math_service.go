package services

type MathService struct{}

func NewMathService() *MathService {
	return &MathService{}
}

func (m *MathService) CalcSquareArea(width int, height int) int {
	return width * height
}

func (m *MathService) CalcTriangleArea(width, height int) int {
	return width * height * (1 / 2)
}
