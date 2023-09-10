package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	tonnam := person{
		firstName: "Tonnam1",
		lastName:  "Warin",
		contactInfo: contactInfo{
			email:   "tonnamwarin@gmail.com",
			zipCode: 12365,
		},
	}

	tonnam.updateName("cute")
	tonnam.print()
}
// receiver โดยในตอนนี้ค่าไม่เปลี่ยนเพราะ go มีการ copy ค่า struct ของ tonnam ไว้แล้วเก็บไปอีก ฟกกพำหห 
//ซึ่ง address ที่ทำการเปลี่ยนค่าเป็น address ที่ copy มา พอ print ออกมาเลยยังมีค่าเดิม
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
