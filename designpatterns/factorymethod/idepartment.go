package factorymethod

type iDepartment interface {
	setName(name string)
	setEmployees(employees int)
	getName() string
	getEmployees() int
}