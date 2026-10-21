package appcatalog

import "testing"

// Test_MatchesVersion covers the chart version formats architect and App Build Suite have
// published.
//
// The orb 6.x format is "<version>-<full 40 character SHA>". Since then charts carry a
// gitsemver development version: the v2 schema "<version>-dev.<branch>.<date>.<time>.h<short
// SHA>" (architect orb 9.x) or the v3 schema "<version>-r<crc32>t<YYYYMMDDHHMMSS>h<short
// SHA>" (App Build Suite). Both abbreviate the trailing SHA to seven characters. Callers
// pass either a full SHA (apptest App.SHA, from CIRCLE_SHA1) or a plain chart version
// (apptest App.Version).
func Test_MatchesVersion(t *testing.T) {
	const (
		fullSHA  = "a7818251ce584fd3279f8457c90fa2e02dbb3ba4"
		otherSHA = "2df89eee6b9775863be1dcf6bdb375ef108026cd"
		// Synthetic: only the first seven characters come from v3Entry.
		v3SHA = "150308c4e1f2a3b4c5d6e7f8091a2b3c4d5e6f70"

		// Real entries taken from the published control-plane-test-catalog and
		// default-test-catalog indexes.
		legacyEntry = "2.0.1-2df89eee6b9775863be1dcf6bdb375ef108026cd"
		devEntry    = "2.0.2-dev.teams-alignment-branch.2026-08-19.13-58-39.ha781825"
		v3Entry     = "4.2.2-r0fb6a962t20261021065246h150308c"
	)

	testCases := []struct {
		name         string
		entryVersion string
		appVersion   string
		expected     bool
	}{
		{
			name:         "case 0: orb 6.x entry matches its full SHA",
			entryVersion: legacyEntry,
			appVersion:   otherSHA,
			expected:     true,
		},
		{
			name:         "case 1: orb 6.x entry does not match a different full SHA",
			entryVersion: legacyEntry,
			appVersion:   fullSHA,
			expected:     false,
		},
		{
			name:         "case 2: orb 9.x dev entry matches the full SHA it abbreviates",
			entryVersion: devEntry,
			appVersion:   fullSHA,
			expected:     true,
		},
		{
			name:         "case 3: orb 9.x dev entry does not match a different full SHA",
			entryVersion: devEntry,
			appVersion:   otherSHA,
			expected:     false,
		},
		{
			name:         "case 4: orb 9.x dev entry matches the abbreviated SHA itself",
			entryVersion: devEntry,
			appVersion:   "a781825",
			expected:     true,
		},
		{
			name:         "case 5: plain chart version still matches by suffix",
			entryVersion: "3.8.1",
			appVersion:   "3.8.1",
			expected:     true,
		},
		{
			name:         "case 6: plain chart version does not match a different version",
			entryVersion: "3.8.1",
			appVersion:   "3.8.2",
			expected:     false,
		},
		{
			name:         "case 7: a plain version is never compared against the dev SHA",
			entryVersion: devEntry,
			appVersion:   "2.0.2",
			expected:     false,
		},
		{
			name:         "case 8: an SHA shorter than the entry abbreviation does not match",
			entryVersion: devEntry,
			appVersion:   "a78182",
			expected:     false,
		},
		{
			name:         "case 9: a non-hex string is not treated as an SHA",
			entryVersion: "1.0.0-dev.branch.2026-08-19.13-58-39.hzzzzzzz",
			appVersion:   fullSHA,
			expected:     false,
		},
		{
			name:         "case 10: a dev version without the h-prefixed SHA does not match",
			entryVersion: "1.0.0-dev.branch.2026-08-19.13-58-39",
			appVersion:   fullSHA,
			expected:     false,
		},
		{
			name:         "case 11: gitsemver v3 dev entry matches the full SHA it abbreviates",
			entryVersion: v3Entry,
			appVersion:   v3SHA,
			expected:     true,
		},
		{
			name:         "case 12: gitsemver v3 dev entry matches a longer abbreviation of its SHA",
			entryVersion: v3Entry,
			appVersion:   "150308c4",
			expected:     true,
		},
		{
			name:         "case 13: gitsemver v3 dev entry does not match a different full SHA",
			entryVersion: v3Entry,
			appVersion:   fullSHA,
			expected:     false,
		},
		{
			name:         "case 14: an SHA that only contains the gitsemver v3 abbreviation does not match",
			entryVersion: v3Entry,
			appVersion:   "0150308c4e1f2a3b4c5d6e7f8091a2b3c4d5e6f7",
			expected:     false,
		},
		{
			name:         "case 15: a plain version is never compared against the gitsemver v3 SHA",
			entryVersion: v3Entry,
			appVersion:   "4.2.2",
			expected:     false,
		},
		{
			name:         "case 16: an h-prefixed SHA without the r and t parts does not match",
			entryVersion: "4.2.2-h150308c",
			appVersion:   v3SHA,
			expected:     false,
		},
		{
			name:         "case 17: a malformed timestamp does not match",
			entryVersion: "4.2.2-r0fb6a962t2026102106524h150308c",
			appVersion:   v3SHA,
			expected:     false,
		},
		{
			name:         "case 18: a non-hex crc32 does not match",
			entryVersion: "4.2.2-r0fb6a96zt20261021065246h150308c",
			appVersion:   v3SHA,
			expected:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MatchesVersion(tc.entryVersion, tc.appVersion)
			if actual != tc.expected {
				t.Fatalf("MatchesVersion(%#q, %#q) == %v, want %v", tc.entryVersion, tc.appVersion, actual, tc.expected)
			}
		})
	}
}
