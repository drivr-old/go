package place

// Repository - interface for place repository
type Repository interface {
	Query(query *Query) ([]Place, error)
}
