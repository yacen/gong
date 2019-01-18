package server

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

type TLSNextProto map[string]func(*http.Server, *tls.Conn, http.Handler)

type ConnState func(net.Conn, http.ConnState)

type ServerBuilder struct {
	_Addr              string
	_Handler           http.Handler // handler to invoke, http.DefaultServeMux if nil
	_TLSConfig         *tls.Config
	_ReadTimeout       time.Duration
	_ReadHeaderTimeout time.Duration
	_WriteTimeout      time.Duration
	_IdleTimeout       time.Duration
	_MaxHeaderBytes    int
	_TLSNextProto      TLSNextProto
	_ConnState         ConnState
	_ErrorLog          *log.Logger
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

// TCP address to listen on, ":http" if empty
// If addr is blank, ":http" is used.
func (b *ServerBuilder) Addr(addr string) *ServerBuilder {
	b._Addr = addr
	return b
}

// ReadTimeout is the maximum duration for reading the entire
// request, including the body.
//
// Because ReadTimeout does not let Handlers make per-request
// decisions on each request body's acceptable deadline or
// upload rate, most users will prefer to use
// ReadHeaderTimeout. It is valid to use them both.
func (b *ServerBuilder) ReadTimeout(readTimeout time.Duration) *ServerBuilder {
	b._ReadTimeout = readTimeout
	return b
}

// ReadHeaderTimeout is the amount of time allowed to read
// request headers. The connection's read deadline is reset
// after reading the headers and the Handler can decide what
// is considered too slow for the body.
func (b *ServerBuilder) ReadHeaderTimeout(readHeaderTimeout time.Duration) *ServerBuilder {
	b._ReadHeaderTimeout = readHeaderTimeout
	return b
}

// WriteTimeout is the maximum duration before timing out
// writes of the response. It is reset whenever a new
// request's header is read. Like ReadTimeout, it does not
// let Handlers make decisions on a per-request basis.
func (b *ServerBuilder) WriteTimeout(writeTimeout time.Duration) *ServerBuilder {
	b._WriteTimeout = writeTimeout
	return b
}

// IdleTimeout is the maximum amount of time to wait for the
// next request when keep-alives are enabled. If IdleTimeout
// is zero, the value of ReadTimeout is used. If both are
// zero, ReadHeaderTimeout is used.
func (b *ServerBuilder) IdleTimeout(idleTimeout time.Duration) *ServerBuilder {
	b._IdleTimeout = idleTimeout
	return b
}

// MaxHeaderBytes controls the maximum number of bytes the
// server will read parsing the request header's keys and
// values, including the request line. It does not limit the
// size of the request body.
// If zero, DefaultMaxHeaderBytes is used.
func (b *ServerBuilder) MaxHeaderBytes(maxHeaderBytes int) *ServerBuilder {
	b._MaxHeaderBytes = maxHeaderBytes
	return b
}

// TLSNextProto optionally specifies a function to take over
// ownership of the provided TLS connection when an NPN/ALPN
// protocol upgrade has occurred. The map key is the protocol
// name negotiated. The Handler argument should be used to
// handle HTTP requests and will initialize the Request's TLS
// and RemoteAddr if not already set. The connection is
// automatically closed when the function returns.
// If TLSNextProto is not nil, HTTP/2 support is not enabled
// automatically.
func (b *ServerBuilder) TLSNextProto(proto TLSNextProto) *ServerBuilder {
	b._TLSNextProto = proto
	return b
}

// ConnState specifies an optional callback function that is
// called when a client connection changes state. See the
// ConnState type and associated constants for details.
func (b *ServerBuilder) ConnState(connState ConnState) *ServerBuilder {
	b._ConnState = connState
	return b
}

// ErrorLog specifies an optional logger for errors accepting
// connections, unexpected behavior from handlers, and
// underlying FileSystem errors.
// If nil, logging is done via the log package's standard logger.
func (b *ServerBuilder) ErrorLog(errorLog *log.Logger) *ServerBuilder {
	b._ErrorLog = errorLog
	return b
}

// TLSConfig optionally provides a TLS configuration for use
// by ServeTLS and ListenAndServeTLS. Note that this value is
// cloned by ServeTLS and ListenAndServeTLS, so it's not
// possible to modify the configuration with methods like
// tls.Config.SetSessionTicketKeys. To use
// SetSessionTicketKeys, use Server.Serve with a TLS Listener
// instead.
func (b *ServerBuilder) TLSConfig(config *tls.Config) *ServerBuilder {
	b._TLSConfig = config
	return b
}

func (b *ServerBuilder) Build() *Server {
	server := &Server{
		server: &http.Server{
			Addr:              b._Addr,
			TLSConfig:         b._TLSConfig,
			ReadTimeout:       b._ReadTimeout,
			ReadHeaderTimeout: b._ReadHeaderTimeout,
			WriteTimeout:      b._WriteTimeout,
			IdleTimeout:       b._IdleTimeout,
			MaxHeaderBytes:    b._MaxHeaderBytes,
			TLSNextProto:      b._TLSNextProto,
			ConnState:         b._ConnState,
			ErrorLog:          b._ErrorLog,
		},
	}
	return server
}
