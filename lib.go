package main

import (
	"fmt"
	"strings"
)

func lcs_multiple_strings(s1 []string, s2 []string) []string {
	M := len(s1)
	N := len(s2)
	arr := make([][][]string, M+1)
	for index := range arr {
		arr[index] = make([][]string, N+1)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if s1[i] == s2[j] {
				arr[i+1][j+1] = append(arr[i][j], s1[i])
			} else {
				str := arr[i][j+1]
				if len(arr[i+1][j]) > len(str) {
					str = arr[i+1][j]
				}
				arr[i+1][j+1] = str
			}
		}
	}

	return arr[M][N]
}

type DiffResult struct {
	Operation int
	OldValues []string
	NewValues []string
}

const (
	OpDelete  = -1
	OpInsert  = 1
	OpReplace = 2
)

func (e *DiffResult) String() string {
	var result strings.Builder
	for _, s := range e.OldValues {
		result.WriteString(fmt.Sprintf("- %s\n", s))
	}

	for _, s := range e.NewValues {
		result.WriteString(fmt.Sprintf("+ %s\n", s))
	}

	return result.String()
}

type Diffs []*DiffResult

func (d Diffs) String() string {
	var result strings.Builder
	for _, e := range d {
		result.WriteString(e.String())
		result.WriteString("-------------------\n")
	}
	return result.String()
}

func diff(s1 []string, s2 []string) Diffs {
	common := lcs_multiple_strings(s1, s2)

	p1 := 0
	p2 := 0

	result := []*DiffResult{}
	appendDiff := func(s1Diff, s2Diff []string) {
		if len(s1Diff) > 0 && len(s2Diff) > 0 {
			result = append(result, &DiffResult{
				Operation: OpReplace,
				OldValues: s1Diff,
				NewValues: s2Diff,
			})
		} else if len(s1Diff) > 0 {
			result = append(result, &DiffResult{
				Operation: OpDelete,
				OldValues: s1Diff,
			})
		} else if len(s2Diff) > 0 {
			result = append(result, &DiffResult{
				Operation: OpInsert,
				NewValues: s2Diff,
			})
		}
	}

	for commonIndex := 0; commonIndex < len(common); commonIndex++ {
		s1Diff := []string{}
		for s1[p1] != common[commonIndex] {
			s1Diff = append(s1Diff, s1[p1])
			p1++
		}

		s2Diff := []string{}
		for s2[p2] != common[commonIndex] {
			s2Diff = append(s2Diff, s2[p2])
			p2++
		}
		appendDiff(s1Diff, s2Diff)

		p1++
		p2++
	}

	// check remaining leftover
	appendDiff(s1[p1:], s2[p2:])

	return result
}
