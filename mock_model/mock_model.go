package mock_model

type Ctx struct {
	buf []string
}

type MockModel struct {
	ctx   Ctx
	inner Model
}

func (m MockModel) GetID(id int) int {
	return m.inner.GetID(id)
}

type Model struct {
	ctx Ctx
}

func (m Model) GetID(_ int) int {
	return 0
}
