package pkg_parent

import "fmt"

type Data struct {
	Id   int
	Name string
}

type Bak interface {
	Save()
}

func (r Data) Save() {
	fmt.Println("Data save. id = ", r.Id, "; name = ", r.Name)
}
