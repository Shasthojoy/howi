// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package mail

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// ParseAddress parses a single RFC 5322 address, e.g. "John Doe <john.doe@example.com>"
func ParseAddress(address string) (*Address, error) {
	maddr, err := mail.ParseAddress(address)
	if err != nil {
		return nil, err
	}
	addr := &Address{
		Name:  maddr.Name,
		Email: maddr.Address,
		Addr:  maddr.String(),
	}
	return addr, nil
}

// ParseAddressFromPublicKey reads an openpgp.entity identities from the
// given armored block key and composes Address from first valid identity.
// It returns error if no identity could be used for ParseAddress.
func ParseAddressFromPublicKey(armoredBlock string) (*Address, error) {
	br := strings.NewReader(armoredBlock)
	block, err := armor.Decode(br)
	if err != nil {
		return nil, err
	}

	pr := packet.NewReader(block.Body)
	entity, err := openpgp.ReadEntity(pr)
	if err != nil {
		return nil, err
	}

	var addr *Address
	err = errors.New("failed to parse email address")
	for _, id := range entity.Identities {
		str := fmt.Sprintf("%q <%s>", id.UserId.Name, id.UserId.Email)
		addr, err = ParseAddress(str)
		if err == nil && addr != nil {
			break
		}
	}
	return addr, err
}

// Address represents a single email address.
// An address such as "John Doe <john.doe@example.com>" is represented as
//
//  Address{
//    Addr: "John Doe <john.doe@example.com>"
//    Name: "John Doe",
//    Address: "john.doe@example.com"}.
//  }
//
type Address struct {
	Addr  string `json:"addr"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// String formats the address as a valid RFC 5322 address.
// If the address's name contains non-ASCII characters
// the name will be rendered according to RFC 2047.
// this should be matching with net/mail.Address.String()
func (a *Address) String() string {
	// normal
	if a.Addr != "" {
		return a.Addr
	}
	// struct literal
	if a.Name != "" || a.Email != "" {
		str := fmt.Sprintf("%q <%s>", a.Name, a.Email)
		addr, _ := ParseAddress(str)
		if addr != nil {
			a.Addr = addr.Addr
		}
	}
	return a.Addr
}
