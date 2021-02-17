package storage

import "github.com/Vlad1slavZhuk/TelegramBot/model"

type Memory struct {
	products []*model.Product
}

func NewMemory() *Memory {
	mem := &Memory{
		products: make([]*model.Product, 0),
	}

	mem.products = append(mem.products,
		&model.Product{Name: "Молоко"},
		&model.Product{Name: "Сметана"},
	)

	return mem
}

func (m *Memory) AddProduct(pr *model.Product) error {
	m.products = append(m.products, pr)
	return nil
}

func (m *Memory) RemoveProduct(name string) error {
	for id, v := range m.products {
		if v.Name == name {
			copy(m.products[id:], m.products[id+1:])
			m.products[len(m.products)-1] = nil
			m.products = m.products[:len(m.products)-1]
			break
		}
	}
	return nil
}

func (m *Memory) GetAll() []*model.Product {
	return m.products
}
