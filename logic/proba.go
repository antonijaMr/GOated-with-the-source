package logic

import (
	"fmt"
	"reflect"
)

type Map_value1 struct {
	Values     []int
	Free_space []int
}
type Item1[T any] struct {
	Content T
	Index   int
	Tag     []string
}
type GoMap[T any] struct {
	Baza          map[string]Map_value1
	Item_list     []Item1[T]
	MainFreeSpace []int
}

func (m *GoMap[T]) ReadTag(tag string) []Item1[T] {
	b := m.Baza[tag].Values
	a := []Item1[T]{}
	for i := 0; i < len(b); i++ {
		temp := m.Item_list[b[i]].Content
		if !reflect.ValueOf(&temp).Elem().IsZero() {
			a = append(a, m.Item_list[b[i]])
		}
	}

	return a
}

func (m *GoMap[T]) Delete(index int) {
	m.MainFreeSpace = append(m.MainFreeSpace, index)
	keys := m.Item_list[index].Tag //lista stringova
	//trebamo izbrisati item
	//prvo prolazimo kroz sve tagove
	//prolzimo kroz values od itema koji brisemo
	//u free space ubacujemo slobodan index iz value
	for num := range keys {
		//go through all tags of item at given inde
		a_key := keys[num]
		values := m.Baza[a_key].Values
		for y := range values {
			if values[y] == index {
				//u values postoji neiskoristen index
				//u free space stavljamo y
				x := m.Baza[a_key]
				x.Free_space = append(x.Free_space, y)
				m.Baza[keys[num]] = x
			}
		}
	}
	temp := *new(T)
	m.Item_list[index].Content = temp //free up space of deleted item
	m.Item_list[index].Tag = nil
}
func (m *GoMap[T]) CreateTag(tagName string) {
	//ukoliko tag vec postoji
	_, ok := m.Baza[tagName]
	if ok {
		return
	}
	//ukoliko tag ne postoji
	fmt.Println(tagName)
	if m.Baza == nil {
		m.Baza = make(map[string]Map_value1)
	}
	m.Baza[tagName] = Map_value1{}
}

/*createItem*/
func (m *GoMap[T]) Add(newItem Item1[T]) {
	// kao index newItema se uvijek salje nula
	//jer ne mozemo znati na kojem se mjestu nalazi
	//dodajemo item u glavni array
	//u prvo slobodno mjesto ako postoji
	if len(m.MainFreeSpace) != 0 {
		indexIzFreeSpace := m.MainFreeSpace[len(m.MainFreeSpace)-1]
		//mainFreeSpace popback
		if len(m.MainFreeSpace) > 0 {
			m.MainFreeSpace = m.MainFreeSpace[:len(m.MainFreeSpace)-1]
		}
		m.Item_list[indexIzFreeSpace] = newItem
		newItem.Index = indexIzFreeSpace
	} else {
		//normalno stavljamo item u item list
		newItem.Index = len(m.Item_list)
		m.Item_list = append(m.Item_list, newItem)
	}

	//dodajemo item tamo gdje pripada po tagu
	for i := range newItem.Tag {
		a_tag := newItem.Tag[i]
		//ukoliko tag ne postoji kreiramo tag
		_, ok := m.Baza[a_tag]
		if !ok {
			m.CreateTag(a_tag)
		}

		//ako tag postoji
		var tag_koji_postoji = m.Baza[a_tag]
		//ukoliko ne postoji ni jedno slobodno mjesto dodajemo na kraj
		if len(tag_koji_postoji.Free_space) == 0 {
			x := m.Baza[a_tag]
			x.Values = append(m.Baza[a_tag].Values, newItem.Index)
			m.Baza[a_tag] = x
		} else {
			//ukoliko postoji slobodno mjesto
			y := m.Baza[a_tag]
			y.Values[y.Free_space[len(y.Free_space)-1]] = newItem.Index

			if len(y.Free_space) > 0 {
				y.Free_space = y.Free_space[:len(y.Free_space)-1]
			}
			m.Baza[a_tag] = y
		}
	}
}
