package user

// User all the fields that user needs.
// This could be a structure defined in protocol buffers and retreived using gRPC
type User struct {
	Username string
	// Additional user info
}
