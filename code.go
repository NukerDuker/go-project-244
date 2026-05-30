package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
	diff := make(map[string]any)

	for key1, value1 := range parsed1 {
		value2, exists := parsed2[key1]
		if exists && parsed1[key1] == parsed2[key1] {
			diff["  "+key1] = value1
		}
		if exists && parsed1[key1] != parsed2[key1] {
			diff["- "+key1] = value1
			diff["+ "+key1] = value2
		}
		if !exists {
			diff["- "+key1] = value1
		}
	}
	for key2, value2 := range parsed2 {
		_, exists := diff[key2]
		if exists {
			continue
		}
		_, exists = parsed1[key2]
		if !exists {
			diff["+ "+key2] = value2
		}
	}
	bytes, err := json.Marshal(diff)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
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
