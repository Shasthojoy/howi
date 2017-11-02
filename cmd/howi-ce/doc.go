// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

/*
HOWI-CE Collection of extended Go standard libraries, replacements, helpers and
additional packages to transform HOWI API from it's other language bindings into Go.

Usage:
 howi-ce command
 howi-ce command [command-flags] [arguments]
 howi-ce [global-flags] command [command-flags] [arguments]
 howi-ce [global-flags] command ...subcommand [command-flags] [arguments]

 The commands are:

 HOWI-CE

  devel               Subcommands for development and contributing to HOWI CE.

 INTERNAL

  about-cli           Display information about this CLI app


 The global flags are:
  --debug                  enable debug log level. when debug flag is after the command then debugging will be enabled only for that command
  --verbose                enable verbose log level
   -v

  --help                   display help or help for the command. [...command --help]
   -h

*/
package main
