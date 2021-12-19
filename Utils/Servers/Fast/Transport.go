package Fast

import (
	"net"
	"net/url"
)

func (s *Server) Endpoint() (*url.URL, error) {
	s.once.Do(func() {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return
		}
		s.lis = lis
	})
	if s.err != nil {
		return nil, s.err
	}
	return nil, nil
}
