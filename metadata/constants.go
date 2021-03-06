/*
 * constants.go - Some metadata constants used throughout fscrypt
 *
 * Copyright 2017 Google Inc.
 * Author: Joe Richey (joerichey@google.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package metadata

import (
	"crypto/sha256"

	"golang.org/x/sys/unix"
)

// Lengths for our keys, buffers, and strings used in fscrypt.
const (
	// DescriptorLen is the length of all Protector and Policy descriptors.
	DescriptorLen = 2 * unix.FS_KEY_DESCRIPTOR_SIZE
	// We always use 256-bit keys internally (compared to 512-bit policy keys).
	InternalKeyLen = 32
	IVLen          = 16
	SaltLen        = 16
	// We use SHA256 for the HMAC, and len(HMAC) == len(hash size).
	HMACLen = sha256.Size
	// PolicyKeyLen is the length of all keys passed directly to the Keyring
	PolicyKeyLen = unix.FS_MAX_KEY_SIZE
)

var (
	// DefaultOptions use the supported encryption modes and max padding.
	DefaultOptions = &EncryptionOptions{
		Padding:   32,
		Contents:  EncryptionOptions_AES_256_XTS,
		Filenames: EncryptionOptions_AES_256_CTS,
	}
	// DefaultSource is the source we use if none is specified.
	DefaultSource = SourceType_custom_passphrase
)
