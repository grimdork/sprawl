package main

// Import is the site containing users, sites etc.
type Import struct {
	// Sites to create.
	Sites []Site `json:"sites"`
	// Users to create.
	Users []User `json:"users"`
}

// Site to create.
type Site struct {
	// Name of site (usually a domain).
	Name string `json:"name"`
	// Groups to create for the site.
	Groups []Group `json:"groups"`
}

// Group to create for a site.
type Group struct {
	// Name of the group, unique to the site it's in.
	Name string `json:"name"`
	// Permission keywords will be created if missing.
	Permissions []string `json:"permissions"`
}

// User is a loadable file with user, group and site membership settings.
type User struct {
	// Name is required
	Name string `json:"name"`
	// Password may be empty, in which case it's generated.
	Password string `json:"password"`
	// Groups is a list of groups to add this user to.
	Groups []string `json:"groups"`
	// Sites is a list of sites to add this user to.
	Sites []string `json:"sites"`
	// Admin is a list of sites to make this admin of.
	Admin []string `json:"admin"`
}
