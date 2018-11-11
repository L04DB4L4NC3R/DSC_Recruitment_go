package controller

var (
	u UserType
	m ManagementType
)

func Startup() {
	u.RegisterRoute()
	m.registerRoute()
}
