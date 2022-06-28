// Copyright 2022 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTOGENERATED. DO NOT EDIT.

// Package assets is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains all [*.css *.html *.js *.tmpl *.yaml] files found in the package as byte arrays.
package assets

// GetAsset returns an asset by its name. Returns nil if no such asset exists.
func GetAsset(name string) []byte {
	return []byte(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset exists.
func GetAssetString(name string) string {
	return files[name]
}

// GetAssetSHA256 returns the asset checksum. Returns nil if no such asset
// exists.
func GetAssetSHA256(name string) []byte {
	data := fileSha256s[name]
	if data == nil {
		return nil
	}
	return append([]byte(nil), data...)
}

// Assets returns a map of all assets.
func Assets() map[string]string {
	cpy := make(map[string]string, len(files))
	for k, v := range files {
		cpy[k] = v
	}
	return cpy
}

var files = map[string]string{
	"avro_schema.yaml": string([]byte{35, 32,
		67, 111, 112, 121, 114, 105, 103, 104, 116, 32, 50, 48, 50, 50,
		32, 84, 104, 101, 32, 67, 104, 114, 111, 109, 105, 117, 109, 32,
		65, 117, 116, 104, 111, 114, 115, 46, 32, 65, 108, 108, 32, 114,
		105, 103, 104, 116, 115, 32, 114, 101, 115, 101, 114, 118, 101, 100,
		46, 10, 35, 32, 85, 115, 101, 32, 111, 102, 32, 116, 104, 105,
		115, 32, 115, 111, 117, 114, 99, 101, 32, 99, 111, 100, 101, 32,
		105, 115, 32, 103, 111, 118, 101, 114, 110, 101, 100, 32, 98, 121,
		32, 97, 32, 66, 83, 68, 45, 115, 116, 121, 108, 101, 32, 108,
		105, 99, 101, 110, 115, 101, 32, 116, 104, 97, 116, 32, 99, 97,
		110, 32, 98, 101, 10, 35, 32, 102, 111, 117, 110, 100, 32, 105,
		110, 32, 116, 104, 101, 32, 76, 73, 67, 69, 78, 83, 69, 32,
		102, 105, 108, 101, 46, 10, 10, 35, 32, 84, 104, 105, 115, 32,
		105, 115, 32, 121, 97, 109, 108, 32, 102, 111, 114, 109, 97, 116,
		32, 111, 102, 32, 97, 118, 114, 111, 32, 115, 99, 104, 101, 109,
		97, 32, 117, 115, 101, 100, 32, 116, 111, 32, 115, 116, 111, 114,
		101, 32, 110, 105, 110, 106, 97, 108, 111, 103, 32, 116, 111, 32,
		66, 105, 103, 81, 117, 101, 114, 121, 46, 10, 35, 32, 83, 101,
		101, 32, 104, 116, 116, 112, 115, 58, 47, 47, 97, 118, 114, 111,
		46, 97, 112, 97, 99, 104, 101, 46, 111, 114, 103, 47, 100, 111,
		99, 115, 47, 99, 117, 114, 114, 101, 110, 116, 47, 115, 112, 101,
		99, 46, 104, 116, 109, 108, 32, 97, 98, 111, 117, 116, 32, 116,
		104, 101, 32, 65, 86, 82, 79, 32, 115, 99, 104, 101, 109, 97,
		46, 10, 10, 35, 32, 84, 104, 105, 115, 32, 115, 104, 111, 117,
		108, 100, 32, 98, 101, 32, 115, 121, 110, 99, 32, 119, 105, 116,
		104, 32, 98, 113, 83, 99, 104, 101, 109, 97, 32, 105, 110, 32,
		98, 105, 103, 113, 117, 101, 114, 121, 46, 103, 111, 46, 10, 10,
		116, 121, 112, 101, 58, 32, 114, 101, 99, 111, 114, 100, 10, 110,
		97, 109, 101, 58, 32, 110, 105, 110, 106, 97, 108, 111, 103, 10,
		102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 45, 32, 110, 97,
		109, 101, 58, 32, 98, 117, 105, 108, 100, 95, 105, 100, 10, 32,
		32, 32, 32, 116, 121, 112, 101, 58, 32, 108, 111, 110, 103, 10,
		32, 32, 45, 32, 110, 97, 109, 101, 58, 32, 116, 97, 114, 103,
		101, 116, 115, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 10,
		32, 32, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 97, 114,
		114, 97, 121, 10, 32, 32, 32, 32, 32, 32, 105, 116, 101, 109,
		115, 58, 32, 115, 116, 114, 105, 110, 103, 10, 32, 32, 45, 32,
		110, 97, 109, 101, 58, 32, 115, 116, 101, 112, 95, 110, 97, 109,
		101, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 115, 116,
		114, 105, 110, 103, 10, 32, 32, 32, 32, 100, 101, 102, 97, 117,
		108, 116, 58, 32, 39, 39, 10, 32, 32, 45, 32, 110, 97, 109,
		101, 58, 32, 106, 111, 98, 115, 10, 32, 32, 32, 32, 116, 121,
		112, 101, 58, 32, 105, 110, 116, 10, 32, 32, 45, 32, 110, 97,
		109, 101, 58, 32, 111, 115, 10, 32, 32, 32, 32, 116, 121, 112,
		101, 58, 32, 115, 116, 114, 105, 110, 103, 10, 32, 32, 45, 32,
		110, 97, 109, 101, 58, 32, 99, 112, 117, 95, 99, 111, 114, 101,
		10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 105, 110, 116,
		10, 32, 32, 45, 32, 110, 97, 109, 101, 58, 32, 98, 117, 105,
		108, 100, 95, 99, 111, 110, 102, 105, 103, 115, 10, 32, 32, 32,
		32, 116, 121, 112, 101, 58, 10, 32, 32, 32, 32, 32, 32, 116,
		121, 112, 101, 58, 32, 97, 114, 114, 97, 121, 10, 32, 32, 32,
		32, 32, 32, 105, 116, 101, 109, 115, 58, 10, 32, 32, 32, 32,
		32, 32, 32, 32, 110, 97, 109, 101, 58, 32, 98, 117, 105, 108,
		100, 95, 99, 111, 110, 102, 105, 103, 10, 32, 32, 32, 32, 32,
		32, 32, 32, 116, 121, 112, 101, 58, 32, 114, 101, 99, 111, 114,
		100, 10, 32, 32, 32, 32, 32, 32, 32, 32, 102, 105, 101, 108,
		100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		45, 32, 110, 97, 109, 101, 58, 32, 107, 101, 121, 10, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 116, 121, 112, 101,
		58, 32, 115, 116, 114, 105, 110, 103, 10, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 45, 32, 110, 97, 109, 101, 58, 32, 118,
		97, 108, 117, 101, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 116, 121, 112, 101, 58, 32, 115, 116, 114, 105, 110,
		103, 10, 32, 32, 45, 32, 110, 97, 109, 101, 58, 32, 108, 111,
		103, 95, 101, 110, 116, 114, 105, 101, 115, 10, 32, 32, 32, 32,
		116, 121, 112, 101, 58, 10, 32, 32, 32, 32, 32, 32, 116, 121,
		112, 101, 58, 32, 97, 114, 114, 97, 121, 10, 32, 32, 32, 32,
		32, 32, 105, 116, 101, 109, 115, 58, 10, 32, 32, 32, 32, 32,
		32, 32, 32, 110, 97, 109, 101, 58, 32, 108, 111, 103, 95, 101,
		110, 116, 114, 121, 10, 32, 32, 32, 32, 32, 32, 32, 32, 116,
		121, 112, 101, 58, 32, 114, 101, 99, 111, 114, 100, 10, 32, 32,
		32, 32, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 110, 97,
		109, 101, 58, 32, 111, 117, 116, 112, 117, 116, 115, 10, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 116, 121, 112, 101,
		58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 116, 121, 112, 101, 58, 32, 97, 114, 114, 97, 121, 10,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
		105, 116, 101, 109, 115, 58, 32, 115, 116, 114, 105, 110, 103, 10,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 110, 97,
		109, 101, 58, 32, 115, 116, 97, 114, 116, 95, 100, 117, 114, 97,
		116, 105, 111, 110, 95, 115, 101, 99, 10, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 100,
		111, 117, 98, 108, 101, 10, 32, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 45, 32, 110, 97, 109, 101, 58, 32, 101, 110, 100, 95,
		100, 117, 114, 97, 116, 105, 111, 110, 95, 115, 101, 99, 10, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 116, 121, 112,
		101, 58, 32, 100, 111, 117, 98, 108, 101, 10, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 45, 32, 110, 97, 109, 101, 58, 32,
		119, 101, 105, 103, 104, 116, 101, 100, 95, 100, 117, 114, 97, 116,
		105, 111, 110, 95, 115, 101, 99, 10, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 100, 111,
		117, 98, 108, 101, 10, 32, 32, 35, 32, 84, 104, 105, 115, 32,
		105, 115, 32, 97, 108, 115, 111, 32, 117, 115, 101, 100, 32, 102,
		111, 114, 32, 116, 105, 109, 101, 32, 112, 97, 114, 116, 105, 116,
		105, 111, 110, 105, 110, 103, 32, 105, 110, 32, 66, 81, 32, 116,
		97, 98, 108, 101, 46, 10, 32, 32, 45, 32, 110, 97, 109, 101,
		58, 32, 99, 114, 101, 97, 116, 101, 100, 95, 97, 116, 10, 32,
		32, 32, 32, 116, 121, 112, 101, 58, 10, 32, 32, 32, 32, 32,
		32, 116, 121, 112, 101, 58, 32, 108, 111, 110, 103, 10, 32, 32,
		32, 32, 32, 32, 108, 111, 103, 105, 99, 97, 108, 84, 121, 112,
		101, 58, 32, 116, 105, 109, 101, 115, 116, 97, 109, 112, 45, 109,
		105, 99, 114, 111, 115, 10}),
}

var fileSha256s = map[string][]byte{
	"avro_schema.yaml": {129, 70,
		198, 138, 245, 134, 110, 8, 26, 199, 50, 231, 249, 12, 75, 99,
		169, 216, 197, 201, 185, 6, 205, 122, 217, 162, 131, 44, 105, 35,
		173, 76},
}
