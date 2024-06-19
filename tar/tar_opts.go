/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tar

// TarOption represents options to be applied to Tar.
type TarOption func(*tarOpts)

// WithMaxUntarSize sets the limit size for archives being decompressed by Untar.
// When max is equal or less than 0 disables size checks.
func WithMaxUntarSize(max int) TarOption {
	return func(t *tarOpts) {
		t.maxUntarSize = max
	}
}

// WithSkipSymlinks allows for symlinks to be present in the tarball and skips them when decompressing.
func WithSkipSymlinks() TarOption {
	return func(t *tarOpts) {
		t.skipSymlinks = true
	}
}

// WithSkipGzip allows for un-taring plain tar files too, that aren't gzipped.
func WithSkipGzip() TarOption {
	return func(t *tarOpts) {
		t.skipGzip = true
	}
}

func (t *tarOpts) applyOpts(tarOpts ...TarOption) {
	for _, clientOpt := range tarOpts {
		clientOpt(t)
	}
}
