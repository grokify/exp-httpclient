// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httpclient is an experimental design for a new Go HTTP client.
//
// It does not work. Do not use it. It exists for brainstorming,
// prototyping, and reviewing godoc only.
package httpclient // import "inet.af/httpclient"

// MethodGet references https://tools.ietf.org/html/rfc7231#section-4.3.1
const MethodGet = "GET"

// MethodHead references https://tools.ietf.org/html/rfc7231#section-4.3.2
const MethodHead = "HEAD"

// MethodPost references https://tools.ietf.org/html/rfc7231#section-4.3.3
const MethodPost = "POST"

// MethodPut references https://tools.ietf.org/html/rfc7231#section-4.3.4
const MethodPut = "PUT"

// MethodDelete references https://tools.ietf.org/html/rfc7231#section-4.3.5
const MethodDelete = "DELETE"

// MethodConnect references https://tools.ietf.org/html/rfc7231#section-4.3.6
const MethodConnect = "CONNECT"

// MethodOptions references https://tools.ietf.org/html/rfc7231#section-4.3.7
const MethodOptions = "OPTIONS"

// MethodTrace references https://tools.ietf.org/html/rfc7231#section-4.3.8
const MethodTrace = "TRACE"

// MethodPatch references https://tools.ietf.org/html/rfc5789
const MethodPatch = "PATCH"
