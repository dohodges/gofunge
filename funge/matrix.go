package funge

type Matrix struct {
	rows int
	cols int
	data [][]int32
}

func NewMatrix(rows, cols int) *Matrix {
	matrix := &Matrix{
		rows: rows,
		cols: cols,
		data: make([][]int32, rows),
	}

	for r := 0; r < rows; r++ {
		matrix.data[r] = make([]int32, cols)
	}

	return matrix
}

func IdentityMatrix(size int) *Matrix {
	matrix := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		matrix.data[i][i] = 1
	}

	return matrix
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Columns() int {
	return m.cols
}

func (m *Matrix) Get(row, column int) int32 {
	return m.data[row][column]
}

func (m *Matrix) Set(row, column int, value int32) {
	m.data[row][column] = value
}

func (m *Matrix) Add(o *Matrix) *Matrix {
	if m.rows != o.rows || m.cols != o.cols {
		panic("gofunge.Matrix: cannot add matrices of inequal dimensions")
	}

	sum := NewMatrix(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			sum.data[r][c] = m.data[r][c] + o.data[r][c]
		}
	}

	return sum
}

func (m *Matrix) Multiply(o *Matrix) *Matrix {
	if m.cols != o.rows {
		panic("gofunge.Matrix: cannot multiply matrices of inequal columns/rows")
	}

	product := NewMatrix(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < o.cols; c++ {
			for i := 0; i < m.cols; i++ {
				product.data[r][c] += m.data[r][i] * m.data[i][c]
			}
		}
	}

	return product
}

type Vector struct {
	*Matrix
}

func NewVector(size int) Vector {
	return Vector{NewMatrix(size, 1)}
}

func (v Vector) Size() int {
	return v.Matrix.Rows()
}

func (v Vector) Get(axis Axis) int32 {
	return v.Matrix.Get(int(axis), 0)
}

func (v Vector) Set(axis Axis, value int32) {
	v.Matrix.Set(int(axis), 0, value)
}

func (v Vector) Add(w Vector) Vector {
	sum := v.Matrix.Add(w.Matrix)
	return Vector{sum}
}

func (v Vector) Transform(transform *Matrix) Vector {
	product := transform.Multiply(v.Matrix)
	return Vector{product}
}
