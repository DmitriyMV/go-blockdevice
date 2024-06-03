// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package block provides support for operations on blockdevices.
package block

import "os"

// Device wraps blockdevice operations.
type Device struct {
	f         *os.File
	ownedFile bool

	devNo uint64
}

// NewFromFile returns a new Device from the specified file.
func NewFromFile(f *os.File) *Device {
	return &Device{f: f}
}

// Close the device.
//
// No-op if the device was created from a file.
func (d *Device) Close() error {
	if d.ownedFile {
		return d.f.Close()
	}

	return nil
}

// DefaultBlockSize is the default block size in bytes.
const DefaultBlockSize = 512
