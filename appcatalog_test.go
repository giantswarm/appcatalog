package appcatalog

import "testing"

// Test_matchesAppVersion covers both chart version formats architect has published.
//
// The orb 6.x format is "<version>-<full 40 character SHA>"; the orb 9.x format is the
// gitsemver development version, whose trailing SHA is abbreviated to seven characters.
// Callers pass either a full SHA (apptest App.SHA, from CIRCLE_SHA1) or a plain chart
// version (apptest App.Version).
func Test_matchesAppVersion(t *testing.T) {
	const (
		fullSHA  = "a7818251ce584fd3279f8457c90fa2e02dbb3ba4"
		otherSHA = "2df89eee6b9775863be1dcf6bdb375ef108026cd"

		// Real entries taken from the published control-plane-test-catalog index.
		legacyEntry = "2.0.1-2df89eee6b9775863be1dcf6bdb375ef108026cd"
		devEntry    = "2.0.2-dev.teams-alignment-branch.2026-08-19.13-58-39.ha781825"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := matchesAppVersion(tc.entryVersion, tc.appVersion)
			if actual != tc.expected {
				t.Fatalf("matchesAppVersion(%#q, %#q) == %v, want %v", tc.entryVersion, tc.appVersion, actual, tc.expected)
			}
		})
	}
}
