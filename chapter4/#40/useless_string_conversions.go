package _40

import (
	"bytes"
	"io"
	"strings"
)

func getBytesBad(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []byte(sanitizeBad(string(b))), nil
}

func sanitizeBad(s string) string {
	// Hint: Is this necessary?
	return strings.TrimSpace(s)
}

func getBytesGood(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitizeGood(b), nil
}

func sanitizeGood(b []byte) []byte {
	return bytes.TrimSpace(b)
}

//   Trimming operations:
//  - bytes.TrimSpace(b) - removes leading/trailing whitespace
//  - bytes.Trim(b, cutset) - removes leading/trailing bytes in cutset
//  - bytes.TrimLeft(b, cutset) / bytes.TrimRight(b, cutset)
//  - bytes.TrimPrefix(b, prefix) / bytes.TrimSuffix(b, suffix)
//
//  Searching operations:
//  - bytes.Contains(b, subslice) - check if b contains subslice
//  - bytes.Index(b, sep) / bytes.LastIndex(b, sep)
//  - bytes.HasPrefix(b, prefix) / bytes.HasSuffix(b, suffix)
//  - bytes.Count(b, sep) - count non-overlapping instances
//
//  Comparison operations:
//  - bytes.Equal(a, b) - compare two byte slices
//  - bytes.Compare(a, b) - returns -1, 0, or 1
//  - bytes.EqualFold(a, b) - case-insensitive comparison
//
//  Modification operations:
//  - bytes.ToUpper(b) / bytes.ToLower(b) / bytes.Title(b)
//  - bytes.Replace(b, old, new, n) / bytes.ReplaceAll(b, old, new)
//  - bytes.Split(b, sep) - returns [][]byte
//  - bytes.Join(slices, sep) - joins [][]byte
