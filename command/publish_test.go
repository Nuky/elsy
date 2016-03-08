package command

import "testing"

type publishTestData struct {
	Input       string
	TagName     string
	ErrExpected bool
}

type extractor func(string) (string, error)

func TestExtractTagFromBranch(t *testing.T) {
	data := []publishTestData{
		{"origin/master", "latest", false},
		{"origin/release", "snapshot.release", false},
		{"origin/release/1.0", "1.0", false},
		{"origin/foo", "snapshot.foo", false},
		{"origin/foo/bar", "snapshot.foo.bar", false},
		{"origin", "", true},
		{"release", "", true},
		{"foo", "", true},
	}
	doTest(t, &data, extractTagFromBranch)
}

func TestExtractTagFromTag(t *testing.T) {
	data := []publishTestData{
		{"v0.0.0", "v0.0.0", false},
		{"v9.9.9", "v9.9.9", false},
		{"v0.0.0-rc1", "v0.0.0-rc1", false},
		{"v0.0.0-r/c1", "v0.0.0-r.c1", false},
		{"v0.0.0rc1", "snapshot.v0.0.0rc1", false},
		{"9.9.9", "snapshot.9.9.9", false},
		{"9", "snapshot.9", false},
		{"foo-test", "snapshot.foo-test", false},
		{"testing-1.2.3", "snapshot.testing-1.2.3", false},
		{"x/[z", "", true},
		{"", "", true},
	}
	doTest(t, &data, extractTagFromTag)
}

func doTest(t *testing.T, data *[]publishTestData, f extractor) {
	for _, d := range *data {
		if tagName, err := f(d.Input); d.TagName != tagName {
			t.Errorf("expected input: %q to produce tag: %q but got %q instead", d.Input, d.TagName, tagName)
		} else if d.ErrExpected && err == nil {
			t.Errorf("expected input: %q to produce an error but did not receive an error", d.Input)
		} else if !d.ErrExpected && err != nil {
			t.Errorf("expected input: %q to not produce an error but received error: %v", d.Input, err)
		}
	}
}
