package lib

import (
	"jtthinkStudy/casbinTest/models"
)

type RoleStruct struct {
	RoleName  string
	PRoleName string
}

func initRoles() {
	result := make([]*RoleStruct, 0)
	getRoles(0, "", &result)
	for _, role := range result {
		E.AddRoleForUser(role.PRoleName, role.RoleName)
	}
}

func getRoles(PID int, PRole string, result *[]*RoleStruct) {
	roles := make([]*models.RoleModel, 0)
	DB.Where("role_pid = ?", PID).Find(&roles)
	for _, role := range roles {
		if PRole != "" {
			*result = append(*result, &RoleStruct{
				RoleName:  role.Name,
				PRoleName: PRole,
			})
		}
		getRoles(role.ID, role.Name, result)
	}
}

func initUserRoles() {
	var userRoles []struct {
		UserName string `gorm:"column:user_name"`
		RoleName string `gorm:"column:role_name"`
	}
	DB.Raw("select user_name,role_name from users,user_roles,roles where users.user_id = user_roles.user_id and user_roles.role_id = roles.role_id").Find(&userRoles)
	for _, userRole := range userRoles {
		E.AddRoleForUser(userRole.UserName, userRole.RoleName)
	}
}

func initRouteRolesPolicy() {
	var routeRoles []struct {
		RoleName    string `gorm:"column:role_name"`
		RouteURI    string `gorm:"column:route_uri"`
		RouteMethod string `gorm:"column:route_method"`
	}
	DB.Raw("select roles.role_name,routes.route_uri,routes.route_method from route_roles,roles,routes WHERE route_roles.role_id = roles.role_id and route_roles.route_id = routes.route_id").Find(&routeRoles)
	for _, routeRole := range routeRoles {
		E.AddPolicy(routeRole.RoleName, routeRole.RouteURI, routeRole.RouteMethod)
	}
}
