package runner

type Role uint

const (
	RoleServer Role = iota
	RoleClient
)

func (r Role) String() string {
	switch r {
	case RoleServer:
		return "server"
	case RoleClient:
		return "client"
	}
	return "unknown"
}
