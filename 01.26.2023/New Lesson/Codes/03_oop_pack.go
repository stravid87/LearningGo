//This code is not working for 03_oop.go because you have to create your own package with this code.

package payroll

type Employee struct {
	Id    int
	name  string
	phone string
	email string
}

func (e *Employee) SetEmployee(name, phone, email string) {
	e.SetName(name)
	e.SetPhone(phone)
	e.SetEmail(email)
}

func (e *Employee) GetId() int {
	return e.Id
}
func (e *Employee) GetName() string {
	return e.name
}
func (e *Employee) GetPhone() string {
	return e.phone
}
func (e *Employee) GetEmail() string {
	return e.email
}
func (e *Employee) SetId(id int) {
	e.Id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}
func (e *Employee) SetPhone(phone string) {
	e.phone = phone
}
func (e *Employee) SetEmail(email string) {
	e.email = email
}