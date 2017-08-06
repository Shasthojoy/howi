// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package mail

import "net/mail"

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
