// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package contributors

import (
	"sync"

	"github.com/howi-ce/howi/std/net/mail"
)

// NewContributor returns empty Project Contributor instance
func NewContributor() *Contributor {
	return &Contributor{}
}

// Contributor holds information a person or thing that is owner of curent pid
type Contributor struct {
	sync.Mutex
	homeDir string
	emails  map[string]mail.Address
	name    string
	email   mail.Address // primary email
}

// AddEmail for user
func (c *Contributor) AddEmail(address string, primary bool) error {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return err
	}
	if c.emails == nil {
		c.emails = make(map[string]mail.Address)
	}
	if _, ok := c.emails[addr.Email]; !ok {
		if primary {
			for a, obj := range c.emails {
				c.Lock()
				c.emails[a] = obj
				c.Unlock()
			}
			// Set contributors name to anme got from email if name is empty
			if c.name == "" && addr.Name != "" {
				c.SetName(addr.Name)
			}
			c.email = *addr
		}
		c.emails[addr.Email] = *addr
	}
	return nil
}

// SetName sets contributors name
func (c *Contributor) SetName(name string) {
	c.name = name
}

// Name returns the name of the contributor
func (c *Contributor) Name() string {
	return c.name
}

// Email returns contributors primary email
func (c *Contributor) Email() string {
	return c.email.Email
}
