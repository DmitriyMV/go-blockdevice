// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "cstruct -pkg squashfs -struct SuperBlock -input superblock.h -endianness LittleEndian"; DO NOT EDIT.

package squashfs

import "encoding/binary"

var _ = binary.LittleEndian

// SuperBlock is a byte slice representing the superblock.h C header.
type SuperBlock []byte

// Get_magic returns magic.
func (s SuperBlock) Get_magic() uint32 {
	return binary.LittleEndian.Uint32(s[0:4])
}

// Put_magic sets magic.
func (s SuperBlock) Put_magic(v uint32) {
	binary.LittleEndian.PutUint32(s[0:4], v)
}

// Get_inode_count returns inode_count.
func (s SuperBlock) Get_inode_count() uint32 {
	return binary.LittleEndian.Uint32(s[4:8])
}

// Put_inode_count sets inode_count.
func (s SuperBlock) Put_inode_count(v uint32) {
	binary.LittleEndian.PutUint32(s[4:8], v)
}

// Get_mod_time returns mod_time.
func (s SuperBlock) Get_mod_time() uint32 {
	return binary.LittleEndian.Uint32(s[8:12])
}

// Put_mod_time sets mod_time.
func (s SuperBlock) Put_mod_time(v uint32) {
	binary.LittleEndian.PutUint32(s[8:12], v)
}

// Get_block_size returns block_size.
func (s SuperBlock) Get_block_size() uint32 {
	return binary.LittleEndian.Uint32(s[12:16])
}

// Put_block_size sets block_size.
func (s SuperBlock) Put_block_size(v uint32) {
	binary.LittleEndian.PutUint32(s[12:16], v)
}

// Get_frag_count returns frag_count.
func (s SuperBlock) Get_frag_count() uint32 {
	return binary.LittleEndian.Uint32(s[16:20])
}

// Put_frag_count sets frag_count.
func (s SuperBlock) Put_frag_count(v uint32) {
	binary.LittleEndian.PutUint32(s[16:20], v)
}

// Get_compressor returns compressor.
func (s SuperBlock) Get_compressor() uint16 {
	return binary.LittleEndian.Uint16(s[20:22])
}

// Put_compressor sets compressor.
func (s SuperBlock) Put_compressor(v uint16) {
	binary.LittleEndian.PutUint16(s[20:22], v)
}

// Get_block_log returns block_log.
func (s SuperBlock) Get_block_log() uint16 {
	return binary.LittleEndian.Uint16(s[22:24])
}

// Put_block_log sets block_log.
func (s SuperBlock) Put_block_log(v uint16) {
	binary.LittleEndian.PutUint16(s[22:24], v)
}

// Get_flags returns flags.
func (s SuperBlock) Get_flags() uint16 {
	return binary.LittleEndian.Uint16(s[24:26])
}

// Put_flags sets flags.
func (s SuperBlock) Put_flags(v uint16) {
	binary.LittleEndian.PutUint16(s[24:26], v)
}

// Get_id_count returns id_count.
func (s SuperBlock) Get_id_count() uint16 {
	return binary.LittleEndian.Uint16(s[26:28])
}

// Put_id_count sets id_count.
func (s SuperBlock) Put_id_count(v uint16) {
	binary.LittleEndian.PutUint16(s[26:28], v)
}

// Get_version_major returns version_major.
func (s SuperBlock) Get_version_major() uint16 {
	return binary.LittleEndian.Uint16(s[28:30])
}

// Put_version_major sets version_major.
func (s SuperBlock) Put_version_major(v uint16) {
	binary.LittleEndian.PutUint16(s[28:30], v)
}

// Get_version_minor returns version_minor.
func (s SuperBlock) Get_version_minor() uint16 {
	return binary.LittleEndian.Uint16(s[30:32])
}

// Put_version_minor sets version_minor.
func (s SuperBlock) Put_version_minor(v uint16) {
	binary.LittleEndian.PutUint16(s[30:32], v)
}

// Get_root_inode returns root_inode.
func (s SuperBlock) Get_root_inode() uint64 {
	return binary.LittleEndian.Uint64(s[32:40])
}

// Put_root_inode sets root_inode.
func (s SuperBlock) Put_root_inode(v uint64) {
	binary.LittleEndian.PutUint64(s[32:40], v)
}

// Get_bytes_used returns bytes_used.
func (s SuperBlock) Get_bytes_used() uint64 {
	return binary.LittleEndian.Uint64(s[40:48])
}

// Put_bytes_used sets bytes_used.
func (s SuperBlock) Put_bytes_used(v uint64) {
	binary.LittleEndian.PutUint64(s[40:48], v)
}

// Get_id_table returns id_table.
func (s SuperBlock) Get_id_table() uint64 {
	return binary.LittleEndian.Uint64(s[48:56])
}

// Put_id_table sets id_table.
func (s SuperBlock) Put_id_table(v uint64) {
	binary.LittleEndian.PutUint64(s[48:56], v)
}

// Get_xattr_table returns xattr_table.
func (s SuperBlock) Get_xattr_table() uint64 {
	return binary.LittleEndian.Uint64(s[56:64])
}

// Put_xattr_table sets xattr_table.
func (s SuperBlock) Put_xattr_table(v uint64) {
	binary.LittleEndian.PutUint64(s[56:64], v)
}

// Get_inode_table returns inode_table.
func (s SuperBlock) Get_inode_table() uint64 {
	return binary.LittleEndian.Uint64(s[64:72])
}

// Put_inode_table sets inode_table.
func (s SuperBlock) Put_inode_table(v uint64) {
	binary.LittleEndian.PutUint64(s[64:72], v)
}

// Get_dir_table returns dir_table.
func (s SuperBlock) Get_dir_table() uint64 {
	return binary.LittleEndian.Uint64(s[72:80])
}

// Put_dir_table sets dir_table.
func (s SuperBlock) Put_dir_table(v uint64) {
	binary.LittleEndian.PutUint64(s[72:80], v)
}

// Get_frag_table returns frag_table.
func (s SuperBlock) Get_frag_table() uint64 {
	return binary.LittleEndian.Uint64(s[80:88])
}

// Put_frag_table sets frag_table.
func (s SuperBlock) Put_frag_table(v uint64) {
	binary.LittleEndian.PutUint64(s[80:88], v)
}

// Get_export_table returns export_table.
func (s SuperBlock) Get_export_table() uint64 {
	return binary.LittleEndian.Uint64(s[88:96])
}

// Put_export_table sets export_table.
func (s SuperBlock) Put_export_table(v uint64) {
	binary.LittleEndian.PutUint64(s[88:96], v)
}

// SUPERBLOCK_SIZE is the size of the SuperBlock struct.
const SUPERBLOCK_SIZE = 96
