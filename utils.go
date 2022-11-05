package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

func runScript(name string, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	pkg, err := findPackageJSON(cwd)
	if err != nil {
		return err
	}

	if pkg.Scripts == nil {
		pkg.Scripts = &map[string]string{}
	}

	binDirs := findBinDirs(cwd)

	script, ok := (*pkg.Scripts)[name]
	if !ok {
		script, err = resolveBinary(name, binDirs)
		if err != nil {
			return fmt.Errorf("No script or binary found for %q", name)
		}
	}

	env := os.Environ()
	env = append(env, fmt.Sprintf("PATH=%s:%s", strings.Join(binDirs, ":"), os.Getenv("PATH")))

	if ok {
		env = append(env, fmt.Sprintf("npm_lifecycle_event=%s", name))
		env = append(env, fmt.Sprintf("npm_lifecycle_script=%s", name))
	}

	commandArgs := strings.Join(append([]string{script}, args...), " ")

	fmt.Printf("%s %s\n", green.Render("$"), gray.Render(name))
	fmt.Printf("%s %s\n", green.Render("$"), gray.Render(commandArgs))

	command := exec.Command("sh", "-c", commandArgs)
	command.Env = env
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func listScripts() []string {
	empty := []string{}
	cwd, err := os.Getwd()
	if err != nil {
		return empty
	}
	pkg, err := findPackageJSON(cwd)
	if err != nil {
		return empty
	}
	if pkg.Scripts == nil {
		pkg.Scripts = &map[string]string{}
	}

	// Add all executibles to the script list
	binDirs := findBinDirs(cwd)
	for _, dir := range binDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if _, ok := (*pkg.Scripts)[file.Name()]; ok {
				continue
			}
			info, err := os.Stat(filepath.Join(dir, file.Name()))
			if err != nil {
				continue
			}
			if info.Mode()&0111 != 0 {
				(*pkg.Scripts)[file.Name()] = ""
			}
		}
	}

	scripts := maps.Keys(*pkg.Scripts)
	sort.Strings(scripts)
	return scripts
}

func resolveBinary(name string, binDirs []string) (string, error) {
	for _, binDir := range binDirs {
		path := filepath.Join(binDir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("No binary found for %q", name)
}

type PackageJSON struct {
	Scripts *map[string]string `json:"scripts"`
}

func findPackageJSON(current string) (*PackageJSON, error) {
	filename, err := findUp(current, "package.json")
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var pkg PackageJSON
	if err := json.NewDecoder(file).Decode(&pkg); err != nil {
		return nil, err
	}
	return &pkg, nil
}

func findBinDirs(current string) []string {
	moduleDirs := findAllUp(current, "node_modules")
	binDirs := make([]string, 0, len(moduleDirs))
	for _, moduleDir := range moduleDirs {
		binDir := filepath.Join(moduleDir, ".bin")
		binDirs = append(binDirs, binDir)
	}
	return binDirs
}

func findUp(current, name string) (string, error) {
	for {
		path := filepath.Join(current, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
		next := filepath.Dir(current)
		if next == current {
			break
		}
		current = next
	}
	return "", fmt.Errorf("No %s found", name)
}

func findAllUp(current, name string) []string {
	filenames := []string{}
	for {
		path := filepath.Join(current, name)
		if _, err := os.Stat(path); err == nil {
			filenames = append(filenames, path)
		}
		next := filepath.Dir(current)
		if next == current {
			break
		}
		current = next
	}
	return filenames
}
