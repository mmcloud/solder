package constants

// Service represents a servicekind name.
type ServiceKind string

const (
	ServeKind ServiceKind = "serve"
)

func (l ServiceKind) String() string {
	return string(l)
}
