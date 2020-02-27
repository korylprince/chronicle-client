package main

import (
	"bytes"
	"encoding/json"
)

//Entry represents information about a computer
type Entry struct {
	UID              uint32 `json:"uid"`
	Username         string `json:"username"`
	FullName         string `json:"full_name"`
	Serial           string `json:"serial"`
	ClientIdentifier string `json:"client_identifier"`
	Hostname         string `json:"hostname"`
	IP               string `json:"ip"`
}

// Close implements io.Closer
func (e *Entry) Close() error {
	return nil
}

// Close implements io.Reader
func (e *Entry) Read(p []byte) (int, error) {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	err := enc.Encode(e)
	if err != nil {
		return 0, err
	}

	return buf.Read(p)
}
