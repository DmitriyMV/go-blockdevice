// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "cstruct -pkg gpt -struct Header -input header.h -endianness LittleEndian"; DO NOT EDIT.

package gpt

import "encoding/binary"

var _ = binary.LittleEndian

// Header is a byte slice representing the header.h C header.
type Header []byte

// Get_signature returns "EFI PART".
func (s Header) Get_signature() uint64 {
	return binary.LittleEndian.Uint64(s[0:8])
}

// Put_signature sets "EFI PART".
func (s Header) Put_signature(v uint64) {
	binary.LittleEndian.PutUint64(s[0:8], v)
}

// Get_revision returns revision.
func (s Header) Get_revision() uint32 {
	return binary.LittleEndian.Uint32(s[8:12])
}

// Put_revision sets revision.
func (s Header) Put_revision(v uint32) {
	binary.LittleEndian.PutUint32(s[8:12], v)
}

// Get_header_size returns usually 92 bytes.
func (s Header) Get_header_size() uint32 {
	return binary.LittleEndian.Uint32(s[12:16])
}

// Put_header_size sets usually 92 bytes.
func (s Header) Put_header_size(v uint32) {
	binary.LittleEndian.PutUint32(s[12:16], v)
}

// Get_header_crc32 returns header_crc32.
func (s Header) Get_header_crc32() uint32 {
	return binary.LittleEndian.Uint32(s[16:20])
}

// Put_header_crc32 sets header_crc32.
func (s Header) Put_header_crc32(v uint32) {
	binary.LittleEndian.PutUint32(s[16:20], v)
}

// Get_reserved1 returns reserved1.
func (s Header) Get_reserved1() uint32 {
	return binary.LittleEndian.Uint32(s[20:24])
}

// Put_reserved1 sets reserved1.
func (s Header) Put_reserved1(v uint32) {
	binary.LittleEndian.PutUint32(s[20:24], v)
}

// Get_my_lba returns location of this header copy.
func (s Header) Get_my_lba() uint64 {
	return binary.LittleEndian.Uint64(s[24:32])
}

// Put_my_lba sets location of this header copy.
func (s Header) Put_my_lba(v uint64) {
	binary.LittleEndian.PutUint64(s[24:32], v)
}

// Get_alternate_lba returns location of the other header copy.
func (s Header) Get_alternate_lba() uint64 {
	return binary.LittleEndian.Uint64(s[32:40])
}

// Put_alternate_lba sets location of the other header copy.
func (s Header) Put_alternate_lba(v uint64) {
	binary.LittleEndian.PutUint64(s[32:40], v)
}

// Get_first_usable_lba returns first usable LBA for partitions.
func (s Header) Get_first_usable_lba() uint64 {
	return binary.LittleEndian.Uint64(s[40:48])
}

// Put_first_usable_lba sets first usable LBA for partitions.
func (s Header) Put_first_usable_lba(v uint64) {
	binary.LittleEndian.PutUint64(s[40:48], v)
}

// Get_last_usable_lba returns last usable LBA for partitions.
func (s Header) Get_last_usable_lba() uint64 {
	return binary.LittleEndian.Uint64(s[48:56])
}

// Put_last_usable_lba sets last usable LBA for partitions.
func (s Header) Put_last_usable_lba(v uint64) {
	binary.LittleEndian.PutUint64(s[48:56], v)
}

// Get_disk_guid returns disk UUID.
func (s Header) Get_disk_guid() []byte {
	return s[56:72]
}

// Put_disk_guid sets disk UUID.
func (s Header) Put_disk_guid(v []byte) {
	copy(s[56:72], v)
}

// Get_partition_entries_lba returns always 2 in primary header copy.
func (s Header) Get_partition_entries_lba() uint64 {
	return binary.LittleEndian.Uint64(s[72:80])
}

// Put_partition_entries_lba sets always 2 in primary header copy.
func (s Header) Put_partition_entries_lba(v uint64) {
	binary.LittleEndian.PutUint64(s[72:80], v)
}

// Get_num_partition_entries returns num_partition_entries.
func (s Header) Get_num_partition_entries() uint32 {
	return binary.LittleEndian.Uint32(s[80:84])
}

// Put_num_partition_entries sets num_partition_entries.
func (s Header) Put_num_partition_entries(v uint32) {
	binary.LittleEndian.PutUint32(s[80:84], v)
}

// Get_sizeof_partition_entry returns sizeof_partition_entry.
func (s Header) Get_sizeof_partition_entry() uint32 {
	return binary.LittleEndian.Uint32(s[84:88])
}

// Put_sizeof_partition_entry sets sizeof_partition_entry.
func (s Header) Put_sizeof_partition_entry(v uint32) {
	binary.LittleEndian.PutUint32(s[84:88], v)
}

// Get_partition_entry_array_crc32 returns partition_entry_array_crc32.
func (s Header) Get_partition_entry_array_crc32() uint32 {
	return binary.LittleEndian.Uint32(s[88:92])
}

// Put_partition_entry_array_crc32 sets partition_entry_array_crc32.
func (s Header) Put_partition_entry_array_crc32(v uint32) {
	binary.LittleEndian.PutUint32(s[88:92], v)
}

// HEADER_SIZE is the size of the Header struct.
const HEADER_SIZE = 92
