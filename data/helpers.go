package data

import (
	"net/http"
)

func AuthEmployee(r *http.Request) EmployeeIdentifierData {
	return r.Context().Value("IDENTIFIER").(EmployeeIdentifierData)
}

func AuthEmployeeCompanyOffice(r *http.Request) companyOffice {
	employee := AuthEmployee(r)
	return employee.CompanyOffice
}

func AuthEmployeeCompanyOfficeIds(r *http.Request) []int {
	employee := AuthEmployee(r)
	return employee.CompanyOfficeIds
}

func AuthAccess(r *http.Request) AccessIdentifierData {
	return r.Context().Value("ACCESS").(AccessIdentifierData)
}

func AuthCustomer(r *http.Request) MyGXIdentifierData {
	return r.Context().Value("IDENTIFIER").(MyGXIdentifierData)
}

func AuthServiceLocations(r *http.Request) []mygxServiceLocation {
	customer := AuthCustomer(r)

	return customer.ServiceLocations
}

func AuthServiceLocation(r *http.Request, uuid string) *mygxServiceLocation {
	serviceLocations := AuthServiceLocations(r)
	for _, serviceLocation := range serviceLocations {
		if serviceLocation.UUID == uuid {
			return &serviceLocation
		}
	}

	return nil
}

func AuthCustomerCompanyOffice(r *http.Request) mygxCompanyOffice {
	customer := AuthCustomer(r)
	return customer.CompanyOffice
}

func AuthRoles(r *http.Request) *map[string]interface{} {
	access := AuthAccess(r)
	return access.Roles
}

func AuthPermissions(r *http.Request) *map[string]interface{} {
	access := AuthAccess(r)
	return access.Permissions
}

func AuthRoleTo(r *http.Request, name string) bool {
	roles := AuthRoles(r)
	if roles == nil {
		return false
	}

	if role, ok := (*roles)[name]; ok {
		return role.(bool)
	} else {
		return false
	}
}

func AuthPermissionTo(r *http.Request, name string) bool {
	permissions := AuthPermissions(r)
	if permissions == nil {
		return false
	}

	if permission, ok := (*permissions)[name]; ok {
		return permission.(bool)
	} else {
		return false
	}
}
