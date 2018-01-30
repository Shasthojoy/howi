// Copyright 2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package emailaddr

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

// ParseAddressList parses the given string as a list of comma-separated addresses
// of the form "John Doe <john.doe@example.com>" or "john.doe@example.com".
func ParseAddressList(list string) ([]*Address, error) {
	maddrList, err := mail.ParseAddressList(list)
	if err != nil {
		return nil, err
	}
	var listSlice []*Address
	for _, maddr := range maddrList {
		addr := &Address{
			Name:  maddr.Name,
			Email: maddr.Address,
			Addr:  maddr.String(),
		}
		listSlice = append(listSlice, addr)
	}
	return listSlice, nil
}
