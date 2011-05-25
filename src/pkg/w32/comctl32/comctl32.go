// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdi32

import (
    . "w32"
)

var (
    lib uintptr
)

func init() {
    lib = LoadLib("comctl32.dll")
}
