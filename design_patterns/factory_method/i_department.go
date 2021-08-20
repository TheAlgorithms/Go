package factory_method

type iDepartment interface {
	setName(name string)
	setEmployees(employees int)
	getName() string
	getEmployees() int
}