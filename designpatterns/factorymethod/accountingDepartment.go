package factorymethod

type accounting struct {
	department
}

func newAccounting() iDepartment {
	return &accounting{
		department: department{
			name:  "Accounting Department",
			employees: 1,
		},
	}
}