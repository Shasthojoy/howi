// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/okramlabs/howi/pkg/errors"
)

// Really basic example lik std error
func ExampleNew() {
	err := errors.New("your error msg")
	fmt.Println(err)
	fmt.Println(err.Error())
	// Output:
	// your error msg
	// your error msg
}

func ExampleNew_printf() {
	err := errors.New("your error")
	fmt.Printf("%+v", err)

	// output:
	// your error
}

// You can use also arguments for .New which are handled in manner of fmt.Sprint
func ExampleNew_args() {
	err := errors.New("your", "error", "msg", "2")
	fmt.Println(err)
	fmt.Println(err.Error())
	// Output:
	// your error msg 2
	// your error msg 2
}

// You can format errors with .Newf args are handled in manner of fmt.Sprintf
func ExampleNewf() {
	err := errors.Newf("your error %s 3", "msg")
	fmt.Println(err)
	fmt.Println(err.Error())
	// Output:
	// your error msg 3
	// your error msg 3
}

// Example how to use context with errors
func ExampleNewWithContext() {
	ctx := context.WithValue(context.Background(), "err-ctx", "your-ctx-key")

	// trigger this error only if ctx contains "your-ctx-key"
	err := errors.NewWithContext(ctx, "your-ctx-key", "error msg")
	// this error will be nil since ctx has no "err-ctx" key with value  "another-ctx-key",
	err2 := errors.NewWithContext(ctx, "another-ctx-key", "error msg")
	// this error will be nil since ctx has no "err-ctx" key,
	err3 := errors.NewWithContext(context.Background(), "your-ctx-key", "error msg")
	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(err3)
	// Output:
	// ctx: your-ctx-key msg: error msg
	// <nil>
	// <nil>
}

// Example how to get error type
func ExampleGetTypeOf() {
	err_std := errors.New("your error msg")
	err_deprecated := errors.NewDeprecated()
	err_not_implemented := errors.NewNotImplemented()

	fmt.Println(errors.GetTypeOf(nil))
	fmt.Println(errors.GetTypeOf(&err_std))
	fmt.Println(errors.GetTypeOf(&err_deprecated))
	fmt.Println(errors.GetTypeOf(&err_not_implemented))
	// Output:
	// <nil>
	// error
	// DeprecatedErr
	// NotImplementedErr
}

// Example shows you how to validate error types of any other package
func ExampleGetTypeOf_otherPKG() {
	fmt.Println(errors.GetTypeOf(http.ErrShortBody))
	// Output:
	// http.ProtocolError
}

// Example shows you how to create error with stack trace and later retrieve that stack trace
func ExampleWithStackTrace() {
	err := errors.WithStackTrace("your errror")

	fmt.Println(err)
	fmt.Println(err.Error())

	st := err.GetStackTrace()
	fmt.Println(st[0].File(), st[0].Package(), st[0].Func())

	// Output:
	// your errror
	// your errror
	// github.com/okramlabs/howi/pkg/errors/example_test.go errors_test ExampleWithStackTrace
}

func ExampleMultiError() {
	merr := errors.NewMultiError()
	// prints first true
	fmt.Println(merr.Nil())

	merr.Append("first error")
	merr.Appendf("%dnd error", 2)
	merr.Add(errors.New("third error"))
	merr.Append("last error")

	fmt.Println(merr.Error())
	fmt.Println(merr.Len())
	fmt.Println(merr.AsError())
	fmt.Println(merr.Nil())

	fmt.Println("errors:")

	for _, err := range merr {
		fmt.Println(err.Error())
	}
	// Output:
	// true
	// last error (total errors: 4)
	// 4
	// last error (total errors: 4)
	// false
	// errors:
	// first error
	// 2nd error
	// third error
	// last error
}
