package factory_method

type department struct {
	name  string
	employees int
}


func (g *department) setName(name string) {
	g.name = name
}

func (g *department) getName() string {
	return g.name
}

func (g *department) setEmployees(employees int) {
	g.employees = employees
}

func (g *department) getEmployees() int {
	return g.employees
}