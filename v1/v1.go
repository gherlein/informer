// Copyright 2019 Cisco Systems, Inc.
//
// This work incorporates works covered by the following notices:
// ---
// Copyright 2016 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ---
// The MIT License (MIT)
//
// Copyright (c) 2017 Middlemost Systems, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package v1

import "fmt"

// InfoService represents the v1 implementation of InfoService.
type InfoService struct {
	// Metadata.
	Program string
	License string
	URL     string
	// Build information.
	BuildUser string
	BuildDate string
	GoVersion string
	// Version information.
	Version  string
	Revision string
	Branch   string
}

// NewInfoService returns a new InfoService.
func NewInfoService(program, license, url, buildUser, buildDate, goVersion,
	version, revision, branch string) *InfoService {
	return &InfoService{
		Program:   program,
		License:   license,
		URL:       url,
		BuildUser: buildUser,
		BuildDate: buildDate,
		GoVersion: goVersion,
		Version:   version,
		Revision:  revision,
		Branch:    branch,
	}
}

// BuildInfo returns build information as a string.
func (s *InfoService) BuildInfo() string {
	return fmt.Sprintf("(go=%s, user=%s, date=%s)", s.GoVersion, s.BuildUser, s.BuildDate)
}

// Metadata returns metadata as a string.
func (s *InfoService) Metadata() string {
	return fmt.Sprintf("(program=%s, license=%s, url=%s)", s.Program, s.License, s.URL)
}

// String returns metadata, build and version information as a string.
func (s *InfoService) String() string {
	return fmt.Sprintf("(metadata=%s, versionInfo=%s, buildInfo=%s)", s.Metadata(), s.VersionInfo(), s.BuildInfo())
}

// VersionInfo returns version information as a string.
func (s *InfoService) VersionInfo() string {
	return fmt.Sprintf("(version=%s, branch=%s, revision=%s)", s.Version, s.Branch, s.Revision)
}
