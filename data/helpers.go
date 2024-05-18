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

func AuthServiceLocationUUIDs(r *http.Request) []string {
	serviceLocations := AuthServiceLocations(r)

	uuids := make([]string, 0)
	for _, serviceLocation := range serviceLocations {
		uuids = append(uuids, serviceLocation.UUID)
	}

	return uuids
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

func AuthRoleTo(r *http.Request, names ...string) bool {
	if len(names) == 0 {
		return false
	}

	roles := AuthRoles(r)
	if roles == nil {
		return false
	}

	for _, name := range names {
		if access, ok := (*roles)[name]; ok {
			if !access.(bool) {
				return false
			}
		}
	}

	return true
}

func AuthPermissionTo(r *http.Request, names ...string) bool {
	if len(names) == 0 {
		return false
	}

	permissions := AuthPermissions(r)
	if permissions == nil {
		return false
	}

	for _, name := range names {
		if access, ok := (*permissions)[name]; ok {
			if !access.(bool) {
				return false
			}
		}
	}

	return true
}

func AuthAccessTo(accesses *map[string]interface{}, names ...string) bool {
	if accesses == nil {
		return false
	}

	if len(names) == 0 {
		return false
	}

	for _, name := range names {
		if access, ok := (*accesses)[name]; ok {
			if !access.(bool) {
				return false
			}
		}
	}

	return true
}

func AuthPermission(r *http.Request) bool {
	employee := AuthEmployee(r)
	if employee.ID == "" {
		return false
	}

	return employee.Superadmin
}
