package error

// NewResourceNotFound creates an new instance of Error
func NewResourceNotFound(resourceName, resourceValue string) ResourceNotFound {
	return ResourceNotFound{"Key_ResourceNotFound", resourceName, resourceValue}
}

// ResourceNotFound represents HTTP 404 error
type ResourceNotFound struct {
	ErrorKey      string `json:"errorKey"`
	ResourceName  string `json:"resourceName"`
	ResourceValue string `json:"resourceValue"`
}

// Error returns the error string
func (e ResourceNotFound) Error() string {
	return e.ErrorKey
}
