package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/korylprince/macserial"
)

func getUser() (uid uint32, username, fullName string, err error) {
	fi, err := os.Stat("/dev/console")
	if err != nil {
		return 0, "", "", err
	}
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, "", "", fmt.Errorf("unknown stat type %T", fi.Sys())
	}
	u, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
	if err != nil {
		return 0, "", "", fmt.Errorf("could not get user: %w", err)
	}
	uidint, err := strconv.Atoi(u.Uid)
	if err != nil {
		return 0, "", "", fmt.Errorf("could not parse uid: %w", err)
	}
	return uint32(uidint), u.Username, u.Name, nil
}

// Collect retrieves the uid, user, name, serial, and host part of an Entry
func Collect() *Entry {
	uid, user, name, err := getUser()
	if err != nil {
		uid = 0
		user = ""
		name = ""
		log.Println("Error getting user information:", err)
	}
	serial, err := macserial.Get()
	if err != nil {
		serial = ""
		log.Println("Error getting serial information:", err)
	}
	host, err := os.Hostname()
	if err != nil {
		host = ""
		log.Println("Error getting host information:", err)
	}
	clientIdentifier, err := GetClientIdentifier()
	if err != nil {
		clientIdentifier = ""
		log.Println("Error getting ClientIdentifier information:", err)
	}
	return &Entry{
		UID:              uid,
		Username:         user,
		FullName:         name,
		Serial:           serial,
		ClientIdentifier: clientIdentifier,
		Hostname:         host,
	}
}
