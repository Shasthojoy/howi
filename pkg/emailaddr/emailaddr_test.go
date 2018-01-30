// Copyright 2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package emailaddr

import (
	"testing"

	"golang.org/x/crypto/openpgp/errors"
)

func TestParseAddressList(t *testing.T) {
	tests := []struct {
		addrsStr string
		exp      []*Address
	}{
		// Bare address
		{
			`jdoe@machine.example`,
			[]*Address{{
				Email: "jdoe@machine.example",
			}},
		},
		// RFC 5322, Appendix A.1.1
		{
			`John Doe <jdoe@machine.example>`,
			[]*Address{{
				Name:  "John Doe",
				Email: "jdoe@machine.example",
			}},
		},
		// RFC 5322, Appendix A.1.2
		{
			`"Joe Q. Public" <john.q.public@example.com>`,
			[]*Address{{
				Name:  "Joe Q. Public",
				Email: "john.q.public@example.com",
			}},
		},
		{
			`Mary Smith <mary@x.test>, jdoe@example.org, Who? <one@y.test>`,
			[]*Address{
				{
					Name:  "Mary Smith",
					Email: "mary@x.test",
				},
				{
					Email: "jdoe@example.org",
				},
				{
					Name:  "Who?",
					Email: "one@y.test",
				},
			},
		},
		{
			`<boss@nil.test>, "Giant; \"Big\" Box" <sysservices@example.net>`,
			[]*Address{
				{
					Email: "boss@nil.test",
				},
				{
					Name:  `Giant; "Big" Box`,
					Email: "sysservices@example.net",
				},
			},
		},
		// RFC 5322, Appendix A.1.3
		// TODO(dsymonds): Group addresses.

		// RFC 2047 "Q"-encoded ISO-8859-1 address.
		{
			`=?iso-8859-1?q?J=F6rg_Doe?= <joerg@example.com>`,
			[]*Address{
				{
					Name:  `Jörg Doe`,
					Email: "joerg@example.com",
				},
			},
		},
		// RFC 2047 "Q"-encoded US-ASCII address. Dumb but legal.
		{
			`=?us-ascii?q?J=6Frg_Doe?= <joerg@example.com>`,
			[]*Address{
				{
					Name:  `Jorg Doe`,
					Email: "joerg@example.com",
				},
			},
		},
		// RFC 2047 "Q"-encoded UTF-8 address.
		{
			`=?utf-8?q?J=C3=B6rg_Doe?= <joerg@example.com>`,
			[]*Address{
				{
					Name:  `Jörg Doe`,
					Email: "joerg@example.com",
				},
			},
		},
		// RFC 2047, Section 8.
		{
			`=?ISO-8859-1?Q?Andr=E9?= Pirard <PIRARD@vm1.ulg.ac.be>`,
			[]*Address{
				{
					Name:  `André Pirard`,
					Email: "PIRARD@vm1.ulg.ac.be",
				},
			},
		},
		// Custom example of RFC 2047 "B"-encoded ISO-8859-1 address.
		{
			`=?ISO-8859-1?B?SvZyZw==?= <joerg@example.com>`,
			[]*Address{
				{
					Name:  `Jörg`,
					Email: "joerg@example.com",
				},
			},
		},
		// Custom example of RFC 2047 "B"-encoded UTF-8 address.
		{
			`=?UTF-8?B?SsO2cmc=?= <joerg@example.com>`,
			[]*Address{
				{
					Name:  `Jörg`,
					Email: "joerg@example.com",
				},
			},
		},
		// Custom example with "." in name. For issue 4938
		{
			`Asem H. <noreply@example.com>`,
			[]*Address{
				{
					Name:  `Asem H.`,
					Email: "noreply@example.com",
				},
			},
		},
		// RFC 6532 3.2.3, qtext /= UTF8-non-ascii
		{
			`"Gø Pher" <gopher@example.com>`,
			[]*Address{
				{
					Name:  `Gø Pher`,
					Email: "gopher@example.com",
				},
			},
		},
		// RFC 6532 3.2, atext /= UTF8-non-ascii
		{
			`µ <micro@example.com>`,
			[]*Address{
				{
					Name:  `µ`,
					Email: "micro@example.com",
				},
			},
		},
		// RFC 6532 3.2.2, local address parts allow UTF-8
		{
			`Micro <µ@example.com>`,
			[]*Address{
				{
					Name:  `Micro`,
					Email: "µ@example.com",
				},
			},
		},
		// RFC 6532 3.2.4, domains parts allow UTF-8
		{
			`Micro <micro@µ.example.com>`,
			[]*Address{
				{
					Name:  `Micro`,
					Email: "micro@µ.example.com",
				},
			},
		},
		// Issue 14866
		{
			`"" <emptystring@example.com>`,
			[]*Address{
				{
					Name:  "",
					Email: "emptystring@example.com",
				},
			},
		},
		{
			"",
			nil,
		},
	}
	for _, test := range tests {
		addrs, err := ParseAddressList(test.addrsStr)
		if test.exp == nil {
			if err == nil {
				t.Errorf("Parse (lis) of %q: got %+v, want %+v", test.addrsStr, addrs, test.exp)
			}
			continue
		}
		if err != nil || len(addrs) != len(test.exp) {
			t.Errorf("Failed parsing (list) %q: %v", test.addrsStr, err)
			continue
		}
		for i := 0; i < len(addrs); i++ {
			if addrs[i].String() != test.exp[i].String() {
				t.Errorf("Parse (list) of %q: got %+v, want %+v", test.addrsStr, addrs, test.exp)
			}
		}
	}
}

// Check if all valid addresses can be parsed, formatted and parsed again
func TestAddressParsingAndFormatting(t *testing.T) {

	// Should pass
	tests := []string{
		`<Bob@example.com>`,
		`<bob.bob@example.com>`,
		`<".bob"@example.com>`,
		`<" "@example.com>`,
		`<some.mail-with-dash@example.com>`,
		`<"dot.and space"@example.com>`,
		`<"very.unusual.@.unusual.com"@example.com>`,
		`<admin@mailserver1>`,
		`<postmaster@localhost>`,
		"<#!$%&'*+-/=?^_`{}|~@example.org>",
		`<"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual"@strange.example.com>`, // escaped quotes
		`<"()<>[]:,;@\\\"!#$%&'*+-/=?^_{}| ~.a"@example.org>`,                      // escaped backslashes
		`<"Abc\\@def"@example.com>`,
		`<"Joe\\Blow"@example.com>`,
		`<test1/test2=test3@example.com>`,
		`<def!xyz%abc@example.com>`,
		`<_somename@example.com>`,
		`<joe@uk>`,
		`<~@example.com>`,
		`<"..."@test.com>`,
		`<"john..doe"@example.com>`,
		`<"john.doe."@example.com>`,
		`<".john.doe"@example.com>`,
		`<"."@example.com>`,
		`<".."@example.com>`,
		`<"0:"@0>`,
	}

	for _, test := range tests {
		addr, err := ParseAddress(test)
		if err != nil {
			t.Errorf("Couldn't parse address %s: %s", test, err.Error())
			continue
		}
		str := addr.String()
		addr, err = ParseAddress(str)
		if err != nil {
			t.Errorf("ParseAddr(%q) error: %v", test, err)
			continue
		}

		if addr.String() != test {
			t.Errorf("String() round-trip = %q; want %q", addr, test)
			continue
		}
	}

	// Should fail
	badTests := []string{
		`<Abc.example.com>`,
		`<A@b@c@example.com>`,
		`<a"b(c)d,e:f;g<h>i[j\k]l@example.com>`,
		`<just"not"right@example.com>`,
		`<this is"not\allowed@example.com>`,
		`<this\ still\"not\\allowed@example.com>`,
		`<john..doe@example.com>`,
		`<john.doe@example..com>`,
		`<john.doe@example..com>`,
		`<john.doe.@example.com>`,
		`<john.doe.@.example.com>`,
		`<.john.doe@example.com>`,
		`<@example.com>`,
		`<.@example.com>`,
		`<test@.>`,
		`< @example.com>`,
		`<""test""blah""@example.com>`,
		`<""@0>`,
	}

	for _, test := range badTests {
		_, err := ParseAddress(test)
		if err == nil {
			t.Errorf("Should have failed to parse address: %s", test)
			continue
		}
	}
}

func TestAddressFormattingAndParsing(t *testing.T) {
	tests := []*Address{
		// {Name: "@lïce", Email: "alice@example.com"},
		{Name: "Böb O'Connor", Email: "bob@example.com"},
		{Name: "???", Email: "bob@example.com"},
		{Name: "Böb ???", Email: "bob@example.com"},
		{Name: "Böb (Jacöb)", Email: "bob@example.com"},
		{Name: "à#$%&'(),.:;<>@[]^`{|}~'", Email: "bob@example.com"},
		// // https://golang.org/issue/12782
		{Name: "naé, mée", Email: "test.mail@gmail.com"},
	}

	for i, test := range tests {
		parsed, err := ParseAddress(test.String())
		if err != nil {
			t.Errorf("test #%d: ParseAddr(%q) error: %v", i, test.String(), err)
			continue
		}
		if parsed.Name != test.Name {
			t.Errorf("test #%d: Parsed name = %q; want %q", i, parsed.Name, test.Name)
		}
		if parsed.Email != test.Email {
			t.Errorf("test #%d: Parsed address = %q; want %q", i, parsed.Email, test.Email)
		}
	}
}

var validArmoredBlockStr = `
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
0g1UqCORr+W3up5/j0RPiP15fg==
=EqtL
-----END PGP PUBLIC KEY BLOCK-----
	`

var malformedArmoredBlockStr = `
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  mQINBFlfks8BEADMMJCzG3pEYvqosVVromEmABguORrfwYelhQmd7zPO624CDZHJ
  BsKsC5hblNKA5vTH+u05mMJ/5pmnxSpCCFCL9FAgxNSq6v5qNGrRclyuFRpiVLGr
  0g1UqCORr+W3up5/j0RPiP15fg==
  =EqtL
  -----END PGP PUBLIC KEY BLOCK-----
  	`

func TestParseAddressFromValidPublicKey(t *testing.T) {
	addr, err := ParseAddressFromPublicKey(validArmoredBlockStr)
	if err != nil {
		t.Fatalf("Failed to parse armored block string, %s", err)
	}
	if addr.Email != "john.doe@example.com" {
		t.Fatalf("ParseAddressFromPublicKey().Email = %q want %q", addr.Email, "john.doe@example.com")
	}
	if addr.Name != "John Doe" {
		t.Fatalf("ParseAddressFromPublicKey().Name = %q want %q", addr.Name, "John Doe")
	}
}

func TestParseAddressFromMalformedPublicKey(t *testing.T) {
	_, err := ParseAddressFromPublicKey(malformedArmoredBlockStr)
	if err == nil {
		t.Fatalf("TestParseAddressFromMalformedPublicKey: expected errors.StructuralError got nil")
	}

	if _, ok := err.(errors.StructuralError); ok {
		t.Errorf("TestParseAddressFromMalformedPublicKey: expected StructuralError, got:%s", err)
	}
}

func TestParseAddressFromEmptyPublicKey(t *testing.T) {
	_, err := ParseAddressFromPublicKey("")
	if err == nil {
		t.Fatalf("TestParseAddressFromMalformedPublicKey: expected errors.StructuralError got nil")
	}

	if _, ok := err.(errors.StructuralError); ok {
		t.Errorf("TestParseAddressFromMalformedPublicKey: expected StructuralError, got:%s", err)
	}
}
