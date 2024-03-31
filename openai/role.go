package openai

// Role is a type that represents the role of a message
type Role struct {
	name string
}

// SystemRole returns a Role with the name "system"
func SystemRole() Role {
	return Role{name: "system"}
}

// UserRole returns a Role with the name "user"
func UserRole() Role {
	return Role{name: "user"}
}

// AssistantRole returns a Role with the name "assistant"
func AssistantRole() Role {
	return Role{name: "assistant"}
}
