package server

type IdenMsg struct {
	Addr   []byte
	PubKey []byte
}

type Server struct {
	Port    string
	IdenSet map[string]IdenMsg
}

const (
	SuperNode = ":9999"
)

func NewServer(port string) (*Server, error) {
	return &Server{
		Port:    port,
		IdenSet: make(map[string]IdenMsg),
	}, nil
}
