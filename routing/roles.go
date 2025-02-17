package routing

type roles struct {
	userRole      string
	moderatorRole string
	adminRole     string
}

func initRoles() *roles {
	return &roles{
		userRole:      "user",
		moderatorRole: "moderator",
		adminRole:     "admin",
	}
}
