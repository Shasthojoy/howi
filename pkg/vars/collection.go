// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

// Collection holds collection of variables
type Collection map[string]Value

// Getvar retrieves the value of the variable named by the key.
// It returns the value, which will be empty string if the variable is not set
// or value was empty.
func (c Collection) Getvar(k string) (v Value) {
	if len(k) == 0 {
		return ""
	}
	v, _ = c[k]
	return
}

// GetvarOrDefaultTo is same as Getvar but returns default value if
// value of variable [key] is empty or does not exist.
// It only returns this case default it neither sets or exports that default
func (c Collection) GetvarOrDefaultTo(k string, defVal string) (v Value) {
	v = c.Getvar(k)
	if v == "" {
		v = ValueFromString(defVal)
	}
	return
}

// GetvarsWithPrefix return all variables with prefix if any as map[]
func (c Collection) GetvarsWithPrefix(prfx string) (vars Collection) {
	vars = make(Collection)
	for k, v := range c {
		if len(k) >= len(prfx) && k[0:len(prfx)] == prfx {
			vars[k] = v
		}
	}
	return
}
