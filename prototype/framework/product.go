package framework

type Product interface {
	Use(str string)
	CreateClone() Product
}

func NewManager() Manager {

	return Manager{showCase: make(map[string]Product)}
}

type Manager struct {
	showCase map[string]Product
}

func (m *Manager) Register(name string, product Product) {
	m.showCase[name] = product
}

func (m *Manager) Create(protName string) Product {
	p := m.showCase[protName]
	return p
}
