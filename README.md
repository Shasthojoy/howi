| |
| :---: |
| <h2> HOWI Community Edition </h2> |
| [![GoDoc][godoc-img]][godoc-url] [![GitHub license][license-img]][license-url] [![Build Status][travis-ci-img]][travis-ci-url] [![Coverage Status][coverage-img]][coverage-url] [![Go Report Card][go-report-card-img]][go-report-card-link] [![codebeat badge][codebeat-img]][codebeat-link] |

HOWI Community Edition is open source project governed by a [The MIT License][license-url].  
HOWI has been a proprietary [SDK][sdk-link], framework and package collection developed  
by Marko Kungla since 2005. It has modular `addon/plugin/lib` design style  
which as collection enables rapid development very wide range of software with  
clean and stable [API's][api-link]. Majority of HOWI was written in `C` while independent  
Addons, plugins and libraries cover [150+ programming languages][lang-stats-link] which could be  
used for very specific use cases within developed software as source or shared library.  

## Goals of 5th series of

- transform multitude of libraries from other language bindings to Go Lang.
- decide which Addons, Plugins and Libraries can be open sourced.
- decide which Addons, Plugins and Libraries will be deprecated.
- implement pre-release designs in Go for Addons, Plugins and Libraries which will be open sourced.

## Redesign with Go
Following describes HOWI-CE design for transforming it to Go Lang while keeping
HOWI's design style and philosophy.

- **HOWI Addons ./addon/<addon-name>**  
> Addons provide higher level API to include full feature set provided by that Addon and it's  
> Plugins to your software. Addons are safest to use since their high level API introduces rarely
> braking changes even when API's of it's plugins or used libraries change.

  - **Addons may import**
    - Go source
    - own plugins
    - any HOWI STD Library
    - any HOWI Library
    - external libraries
  - **Addons may never import**  
    - other Addons
    - Plugins from other Addons

- **HOWI Plugins ./addon/<addon-name>/plugin/<plugin-name>**  
> Plugins are always sub packages of Addon. While plugins are the core which Addons are made of they  
> are always usable as fully independent packages. Plugins are also safe to use and many cases more  
> reasonable to use when you only need features in your software which are provided that Plugin which  
> case importing/using entire Addon would bring unnecessary overhead. As noted earlier Plugins may
> introduce breaking changes earlier than Addons would so therefore that trade of should be taken  
> into consideration when planning your software.   

  - **Plugins may import**
    - Go source
    - any HOWI STD Library
    - any HOWI Library
  - **Plugins may never import**
    - other Addons

- **HOWI Libraries ./lib/<library-name>**  
> Libraries are low level and packages within HOWI. Purpose of these libraries is
> to provide low level bleeding edge features for HOWI Addons and their Plugins.
> HOWI Libraries mostly are extending or providing stable API for external libraries.
> (libraries can be nested)

  - **Libraries may import**
    - Go source
    - any HOWI STD Library
    - external libraries
    - any HOWI Library
  - **Libraries may never import**
    - any Addons
    - any Plugins

- **HOWI Standard Libraries ./std/<library-name>**  
> Standard Libraries are lowest level and most unstable packages within HOWI. Purpose of these
> libraries to provide low level bleeding edge features for HOWI Addons and their Plugins.
> HOWI Standard Libraries are often extending or replacing language features.
> (libraries can be nested)

  - **Libraries may import**
    - Go source
    - external libraries
  - **Libraries may never import**
    - any Addons
    - any Plugins
    - any HOWI Library

<!-- ASSETS and LINKS -->
[sdk-link]: https://en.wikipedia.org/wiki/Software_development_kit
[api-link]: https://en.wikipedia.org/wiki/Application_programming_interface
[lang-stats-link]: https://github.com/howi-ce/howi/blob/master/docs/orig-langs.md
<!-- travis-ci -->
[travis-ci-img]: https://travis-ci.org/howi-ce/howi.svg?branch=master
[travis-ci-url]: https://travis-ci.org/howi-ce/howi

<!-- License -->
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square
[license-url]: https://raw.githubusercontent.com/howi-ce/howi/master/LICENSE

<!-- godoc -->
[godoc-img]: https://godoc.org/github.com/howi-ce/howi?status.png
[godoc-url]: https://godoc.org/github.com/howi-ce/howi

<!-- coverage -->
[coverage-img]: https://coveralls.io/repos/github/howi-ce/howi/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/howi-ce/howi?branch=master

<!-- Go Report Card -->
[go-report-card-img]: https://goreportcard.com/badge/github.com/howi-ce/howi
[go-report-card-link]: https://goreportcard.com/report/github.com/howi-ce/howi

<!-- codebeat -->
[codebeat-img]: https://codebeat.co/badges/7ef35523-4eec-416e-8bc9-a5ffc6be7457
[codebeat-link]: https://codebeat.co/projects/github-com-howi-ce-howi-master
