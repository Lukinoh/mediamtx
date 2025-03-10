package core

import (
	"time"

	"github.com/google/uuid"

	"github.com/bluenviron/mediamtx/internal/conf"
)

type apiError struct {
	Error string `json:"error"`
}

type apiPathConfList struct {
	ItemCount int                  `json:"itemCount"`
	PageCount int                  `json:"pageCount"`
	Items     []*conf.OptionalPath `json:"items"`
}

type apiPathSourceOrReader struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type apiPath struct {
	Name          string                  `json:"name"`
	ConfName      string                  `json:"confName"`
	Source        *apiPathSourceOrReader  `json:"source"`
	Ready         bool                    `json:"ready"`
	ReadyTime     *time.Time              `json:"readyTime"`
	Tracks        []string                `json:"tracks"`
	BytesReceived uint64                  `json:"bytesReceived"`
	Readers       []apiPathSourceOrReader `json:"readers"`
}

type apiPathList struct {
	ItemCount int        `json:"itemCount"`
	PageCount int        `json:"pageCount"`
	Items     []*apiPath `json:"items"`
}

type apiHLSMuxer struct {
	Path        string    `json:"path"`
	Created     time.Time `json:"created"`
	LastRequest time.Time `json:"lastRequest"`
	BytesSent   uint64    `json:"bytesSent"`
}

type apiHLSMuxerList struct {
	ItemCount int            `json:"itemCount"`
	PageCount int            `json:"pageCount"`
	Items     []*apiHLSMuxer `json:"items"`
}

type apiRTSPConn struct {
	ID            uuid.UUID `json:"id"`
	Created       time.Time `json:"created"`
	RemoteAddr    string    `json:"remoteAddr"`
	BytesReceived uint64    `json:"bytesReceived"`
	BytesSent     uint64    `json:"bytesSent"`
}

type apiRTSPConnsList struct {
	ItemCount int            `json:"itemCount"`
	PageCount int            `json:"pageCount"`
	Items     []*apiRTSPConn `json:"items"`
}

type apiRTMPConnState string

const (
	apiRTMPConnStateIdle    apiRTMPConnState = "idle"
	apiRTMPConnStateRead    apiRTMPConnState = "read"
	apiRTMPConnStatePublish apiRTMPConnState = "publish"
)

type apiRTMPConn struct {
	ID            uuid.UUID        `json:"id"`
	Created       time.Time        `json:"created"`
	RemoteAddr    string           `json:"remoteAddr"`
	State         apiRTMPConnState `json:"state"`
	Path          string           `json:"path"`
	BytesReceived uint64           `json:"bytesReceived"`
	BytesSent     uint64           `json:"bytesSent"`
}

type apiRTMPConnList struct {
	ItemCount int            `json:"itemCount"`
	PageCount int            `json:"pageCount"`
	Items     []*apiRTMPConn `json:"items"`
}

type apiRTSPSessionState string

const (
	apiRTSPSessionStateIdle    apiRTSPSessionState = "idle"
	apiRTSPSessionStateRead    apiRTSPSessionState = "read"
	apiRTSPSessionStatePublish apiRTSPSessionState = "publish"
)

type apiRTSPSession struct {
	ID            uuid.UUID           `json:"id"`
	Created       time.Time           `json:"created"`
	RemoteAddr    string              `json:"remoteAddr"`
	State         apiRTSPSessionState `json:"state"`
	Path          string              `json:"path"`
	Transport     *string             `json:"transport"`
	BytesReceived uint64              `json:"bytesReceived"`
	BytesSent     uint64              `json:"bytesSent"`
}

type apiRTSPSessionList struct {
	ItemCount int               `json:"itemCount"`
	PageCount int               `json:"pageCount"`
	Items     []*apiRTSPSession `json:"items"`
}

type apiSRTConnState string

const (
	apiSRTConnStateIdle    apiSRTConnState = "idle"
	apiSRTConnStateRead    apiSRTConnState = "read"
	apiSRTConnStatePublish apiSRTConnState = "publish"
)

type apiSRTConn struct {
	ID            uuid.UUID       `json:"id"`
	Created       time.Time       `json:"created"`
	RemoteAddr    string          `json:"remoteAddr"`
	State         apiSRTConnState `json:"state"`
	Path          string          `json:"path"`
	BytesReceived uint64          `json:"bytesReceived"`
	BytesSent     uint64          `json:"bytesSent"`
}

type apiSRTConnList struct {
	ItemCount int           `json:"itemCount"`
	PageCount int           `json:"pageCount"`
	Items     []*apiSRTConn `json:"items"`
}

type apiWebRTCSessionState string

const (
	apiWebRTCSessionStateRead    apiWebRTCSessionState = "read"
	apiWebRTCSessionStatePublish apiWebRTCSessionState = "publish"
)

type apiWebRTCSession struct {
	ID                        uuid.UUID             `json:"id"`
	Created                   time.Time             `json:"created"`
	RemoteAddr                string                `json:"remoteAddr"`
	PeerConnectionEstablished bool                  `json:"peerConnectionEstablished"`
	LocalCandidate            string                `json:"localCandidate"`
	RemoteCandidate           string                `json:"remoteCandidate"`
	State                     apiWebRTCSessionState `json:"state"`
	Path                      string                `json:"path"`
	BytesReceived             uint64                `json:"bytesReceived"`
	BytesSent                 uint64                `json:"bytesSent"`
}

type apiWebRTCSessionList struct {
	ItemCount int                 `json:"itemCount"`
	PageCount int                 `json:"pageCount"`
	Items     []*apiWebRTCSession `json:"items"`
}
