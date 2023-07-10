package nw

import (
	"bufio"
	"net"
	"time"

	"github.com/gogo/protobuf/proto"
)

// Conn interface for a connection, either a client connection or server accepted connection
type Conn interface {
	net.Conn
	GetSession() Session    // get the bound session
	GetConnTime() time.Time // get connection init time
	Wait()                  // wait for close
	SetCloseReason(reason string)
}

// Server the interface for wsserver and tcpserver
type Server interface {
	Start(addr string) error
	Stop()
	Broadcast(sessionIds []uint32, data []byte) // broadcast data to all connected sessions
	SetAcceptConnect(acceptConnect bool)
	//GetActiveConnNum() int                      // get current count of connections
}

// Context context for create a dialer or a listener
type Context struct {
	SessionCreator        func(conn Conn) Session
	Splitter              bufio.SplitFunc      // packet splitter
	IPChecker             func(ip string) bool // check if an accepted connection is allowed
	IdleTimeAfterOpen     time.Duration        // idle time when open, conn will be closed if not activated after this time
	ReadBufferSize        int                  // buffer size for reading
	WriteBufferSize       int                  // buffer size for writing
	UseNoneBlockingChan   bool                 // use none blocking chan
	ChanSize              int                  // chan size for bufferring
	MaxMessageSize        int                  // max message size for a single packet
	MergedWriteBufferSize int                  // buffer size for merged write
	DisableMergedWrite    bool                 // disable merge multiple message to a single net.Write
	EnableStatistics      bool                 // enable statistics of packets send and recv
	Extra                 interface{}          // used for special cases when custom data is needed
}

// ProtoMessage interface for gogo protobuf generated code
type ProtoMessage interface {
	proto.Message
	Marshal() ([]byte, error)
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	Size() int
}
