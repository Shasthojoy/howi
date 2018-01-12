// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/okramlabs/howicli/pkg/errors"
)

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

func ExampleNew_args() {
	err := errors.New("your", "error", "msg", "2")
	fmt.Println(err)
	fmt.Println(err.Error())
	// Output:
	// your error msg 2
	// your error msg 2
}

func ExampleNewf() {
	err := errors.Newf("your error %s 3", "msg")
	fmt.Println(err)
	fmt.Println(err.Error())
	// Output:
	// your error msg 3
	// your error msg 3
}

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

// ExampleGetTypeOf_extended shows you how to validate error types of any other package
func ExampleGetTypeOf_otherPKG() {
	fmt.Println(errors.GetTypeOf(http.ErrShortBody))
	// Output:
	// http.ProtocolError
}

func ExampleWithStackTrace() {
	err := errors.WithStackTrace("your errror")

	fmt.Println(err)
	fmt.Println(err.Error())

	for i, f := range err.GetStackTrace() {
		fmt.Println(i, f.File(), f.Package(), f.Func())
	}
	// Output:
	// your errror
	// your errror
	// 0 github.com/okramlabs/howicli/pkg/errors/example_test.go errors_test ExampleWithStackTrace
	// 1 testing/example.go testing runExample
	// 2 testing/example.go testing runExamples
	// 3 testing/testing.go testing (*M).Run
	// 4 github.com/okramlabs/howicli/pkg/errors/_test/_testmain.go main main
	// 5 runtime/proc.go runtime main
	// 6 runtime/asm_amd64.s runtime goexit
}
