package composite

import "fmt"

type MenuComponent interface {
	Print()
}

type MenuItem struct {
	name  string
	price int
}

func (mi *MenuItem) Print() {
	fmt.Println("Name of the dish:", mi.name, "\nPrice is $", mi.price)
}

type Menu struct {
	name  string
	items []MenuComponent
}

func (menu *Menu) Print() {
	fmt.Println("Name of the menu:", menu.name)

	for _, item := range menu.items {
		item.Print()
	}
}

func main() {
	menu := &Menu{
		name: "Main Menu",
		items: []MenuComponent{
			&Menu{
				name: "Dinner",
				items: []MenuComponent{
					&MenuItem{
						name:  "Hamburger",
						price: 12,
					},
					&MenuItem{
						name:  "Shawarma",
						price: 8,
					},
				},
			},
			&MenuItem{
				name:  "Sides",
				price: 24,
			},
		},
	}

	menu.Print()
}
