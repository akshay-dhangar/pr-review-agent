package core

import "strings"

type FileChange struct {
	Path           string
	Added, Deleted int
}

// Summarize counts textual unified-diff additions/deletions. It does not review correctness.
func Summarize(diff string) []FileChange {
	out := []FileChange{}
	current := -1
	inHunk := false
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "diff --git ") {
			current = -1
			inHunk = false
		}
		if strings.HasPrefix(line, "+++ ") {
			name := strings.TrimPrefix(line, "+++ ")
			name = strings.TrimPrefix(name, "b/")
			out = append(out, FileChange{Path: name})
			current = len(out) - 1
			inHunk = false
			continue
		}
		if strings.HasPrefix(line, "@@ ") {
			inHunk = true
			continue
		}
		if current < 0 || !inHunk {
			continue
		}
		if strings.HasPrefix(line, "+") {
			out[current].Added++
		} else if strings.HasPrefix(line, "-") {
			out[current].Deleted++
		}
	}
	return out
}
