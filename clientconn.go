package yarpc

import (
	"net"

	"context"

	"github.com/gogo/protobuf/codec"
	"github.com/hashicorp/yamux"
	"github.com/uber-go/zap"
)

type dialOptions struct {
	codec Codec
	log   zap.Logger
}

type DialOption func(*dialOptions)

func WithCodec(c Codec) DialOption {
	return func(o *dialOptions) {
		o.codec = c
	}
}

func WithLogger(l zap.Logger) DialOption {
	return func(o *dialOptions) {
		o.log = l
	}
}

func Dial(addr string, opt ...DialOption) (*ClientConn, error) {
	return DialContext(context.Background(), addr, opt...)
}

func DialContext(ctx context.Context, addr string, opts ...DialOption) (*ClientConn, error) {
	cn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	s, err := yamux.Client(cn, nil)
	if err != nil {
		return nil, err
	}
	cc := &ClientConn{s: s}
	cc.ctx, cc.cancel = context.WithCancel(ctx)

	for _, opt := range opts {
		opt(&cc.dopts)
	}

	if cc.dopts.codec == nil {
		// TODO(sgc): codec is not safe for concurrent access
		cc.dopts.codec = codec.New(1024)
	}

	if cc.dopts.log == nil {
		cc.dopts.log = zap.New(zap.NullEncoder())
	}

	return cc, nil
}

type ClientConn struct {
	ctx    context.Context
	cancel context.CancelFunc
	s      *yamux.Session
	dopts  dialOptions
	log    zap.Logger
}

func (cc *ClientConn) NewStream() (*yamux.Stream, error) {
	return cc.s.OpenStream()
}

func (cc *ClientConn) Close() error {
	cc.cancel()
	return cc.s.Close()
}
