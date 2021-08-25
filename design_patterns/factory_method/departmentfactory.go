package factory_method

import "fmt"

func getDepartment(departmentType string) (iDepartment, error) {
	if departmentType == "accounting" {
		return newAccounting(), nil
	}
	if departmentType == "finance" {
		return newFinance(), nil
	}
	return nil, fmt.Errorf("wrong department type passed")
}
