package modules

// Project struct of of gitlab's project
type Project struct {
	ID   int    `json:"id"`   // project的ID
	Name string `json:"name"` // project的Name
}
