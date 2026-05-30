package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

func GenDiff(file1, file2 string, format string) (string, error) {
	absPath1, absPath2, err := absPathToFiles(file1, file2)
	if err != nil {
		return "", err
	}
	cont1, cont2, err := readFiles(absPath1, absPath2)
	if err != nil {
		return "", err
	}
	parsed1, parsed2, err := parseFiles(cont1, cont2, absPath1, absPath2)
	if err != nil {
		return "", err
	}

	diff, err := getDiff(parsed1, parsed2)
	if err != nil {
		return "", err
	}
	return diff, nil
}

func getDiff(parsed1 map[string]any, parsed2 map[string]any) (string, error) {
	keysMap := map[string]bool{}

	for key := range parsed1 {
		keysMap[key] = true
	}

	for key := range parsed2 {
		keysMap[key] = true
	}

	keys := make([]string, 0, len(keysMap))

	for key := range keysMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	strDiff := "{"
	for _, key := range keys {
		val1, exists1 := parsed1[key]
		val2, exists2 := parsed2[key]
		if exists1 && exists2 && val1 == val2 {
			strDiff += fmt.Sprintf("\n    %v: %v", key, val1)
		} else if exists1 && exists2 && val1 != val2 {
			strDiff += fmt.Sprintf("\n  - %v: %v\n  + %v: %v", key, val1, key, val2)
		} else if exists1 && !exists2 {
			strDiff += fmt.Sprintf("\n  - %v: %v", key, val1)
		} else if exists2 && !exists1 {
			strDiff += fmt.Sprintf("\n  + %v: %v", key, val2)
		}
	}
	strDiff += "\n}"
	return strDiff, nil
}

func parseFiles(
	cont1 []byte,
	cont2 []byte,
	path1 string,
	path2 string) (map[string]any, map[string]any, error) {
	ext1 := filepath.Ext(path1)
	ext2 := filepath.Ext(path2)
	parsed1, err := parse(cont1, ext1)
	if err != nil {
		return nil, nil, err
	}
	parsed2, err := parse(cont2, ext2)
	if err != nil {
		return nil, nil, err
	}
	return parsed1, parsed2, nil
}

func absPathToFiles(file1, file2 string) (string, string, error) {
	absPath1, err := filepath.Abs(file1)
	if err != nil {
		return "", "", err
	}
	absPath2, err := filepath.Abs(file2)
	if err != nil {
		return "", "", err
	}
	return absPath1, absPath2, nil
}

func readFiles(path1, path2 string) ([]byte, []byte, error) {
	cont1, err := os.ReadFile(path1)
	if err != nil {
		return nil, nil, err
	}
	cont2, err := os.ReadFile(path2)
	if err != nil {
		return nil, nil, err
	}
	return cont1, cont2, nil
}

func parse(content []byte, ext string) (map[string]any, error) {
	result := make(map[string]any)
	switch ext {
	case ".json":
		err := json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
	case ".yaml", ".yml":
		err := yaml.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
	return result, nil
}
