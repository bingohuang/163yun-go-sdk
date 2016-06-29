package cloudcomb

import (
	"fmt"
	"io"
	URL "net/url"
)

const (
	Version = "0.0.1"
)

const (
	// defaultEndPoint: access url, support https
	defaultEndPoint = "open.c.163.com"

	// defaultConnectTimeout: connection timeout when connect to cloudcomb endpoint
	defaultConnectTimeout = 60

	// Default(Min/Max)ChunkSize: set the buffer size when doing copy operation
	defaultChunkSize = 32 * 1024
)

// chunkSize: chunk size when copy
var (
	chunkSize = defaultChunkSize
)

// User Agent
func makeUserAgent() string {
	return fmt.Sprintf("CloudComb Go SDK %s", Version)
}

// URI escape
func escapeURI(uri string) string {
	Uri := URL.URL{}
	Uri.Path = uri
	return Uri.String()
}

// Because of io.Copy use a 32Kb buffer, and, it is hard coded
// user can specify a chunksize with upyun.SetChunkSize
func chunkedCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	buf := make([]byte, chunkSize)

	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])

			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return
}
