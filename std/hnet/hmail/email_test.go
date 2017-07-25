// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hmail

import "testing"

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
