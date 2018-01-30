// Copyright 2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package emailaddr

import "fmt"

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
