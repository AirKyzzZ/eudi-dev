// Copyright 2026 Dominik Schlosser
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	"strconv"
	"strings"
)

// Version is a release version reduced to what a compatibility check needs.
// Pre-release and build metadata are dropped, so the CLI treats 1.19.0-rc.1
// like 1.19.0.
type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion reads a semantic version from /api/version or the CLI build. It
// reports false without a numeric major and minor, such as for the "dev" build
// or an empty version.
func ParseVersion(raw string) (Version, bool) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "v")
	if i := strings.IndexAny(raw, "-+"); i >= 0 {
		raw = raw[:i]
	}
	if raw == "" {
		return Version{}, false
	}
	parts := strings.Split(raw, ".")
	if len(parts) < 2 {
		return Version{}, false
	}
	numbers := make([]int, 3)
	for i := 0; i < len(parts) && i < 3; i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil || n < 0 {
			return Version{}, false
		}
		numbers[i] = n
	}
	return Version{Major: numbers[0], Minor: numbers[1], Patch: numbers[2]}, true
}

func (v Version) String() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch)
}

type Compatibility int

const (
	// CompatibilityUnknown means one side reports no comparable version, such as
	// a development build or an instance that reports no version.
	CompatibilityUnknown Compatibility = iota
	Compatible
	Incompatible
)

// CheckCompatibility compares the CLI's release with an instance's release.
// Under semantic versioning the same major release is compatible.
func CheckCompatibility(cli, instance string) Compatibility {
	cliVersion, cliOK := ParseVersion(cli)
	instanceVersion, instanceOK := ParseVersion(instance)
	if !cliOK || !instanceOK {
		return CompatibilityUnknown
	}
	if cliVersion.Major != instanceVersion.Major {
		return Incompatible
	}
	return Compatible
}
