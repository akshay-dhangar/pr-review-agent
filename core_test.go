package core

import "testing"

func TestDiff(t *testing.T) {
	r := Summarize("diff --git a/x.go b/x.go\n--- a/x.go\n+++ b/x.go\n@@ -1,1 +1,2 @@\n-old\n+new\n+more\n context\n")
	if len(r) != 1 || r[0].Path != "x.go" || r[0].Added != 2 || r[0].Deleted != 1 {
		t.Fatal(r)
	}
	if len(Summarize("")) != 0 {
		t.Fatal("empty diff")
	}
}
