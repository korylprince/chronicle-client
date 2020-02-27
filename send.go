package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const apiPrefix = "/api/v1.1"

//Submit generates an Entry and sends it to a Chronicler
func Submit() error {
	entry := Collect()

	//a custom dialer is required to get local IP
	dial := func(context context.Context, network, addr string) (net.Conn, error) {
		conn, err := (&net.Dialer{}).DialContext(context, network, addr)
		if err != nil {
			return conn, err
		}

		entry.IP = conn.LocalAddr().(*net.TCPAddr).IP.String()

		return conn, err
	}

	t := http.Transport{DialContext: dial}
	client := http.Client{Transport: &t, Timeout: time.Second * 5}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s%s/submit", config.proto, config.Chronicler, apiPrefix), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = entry

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Request not completed successfully")
	}
	log.Printf("Submitted: %#v\n", entry)
	return nil
}
