package structural

import "fmt"

type User struct {
}

func (u *User) FullCleanInHouse(cln Cleaner) {
	fmt.Println("User cleans the whole house")
	cln.CleanInHouse()
}
