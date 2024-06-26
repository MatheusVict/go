// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import _ "unsafe"

// used by os
//go:linkname closedir
//go:linkname readdir_r

// used by internal/poll
//go:linkname fdopendir

// used by internal/syscall/unix
//go:linkname unlinkat
//go:linkname openat
//go:linkname fstatat

// used by cmd/link
//go:linkname msync
//go:linkname fcntl

// mmap should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - modernc.org/memory
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname mmap
