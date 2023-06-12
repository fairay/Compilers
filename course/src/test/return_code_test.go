package main_test

import (
	"course/compilers/compiler"
	"course/compilers/utils"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCompileAndCompare(t *testing.T) {
	dir := "return_code"

	files, err := filepath.Glob(filepath.Join(dir, "*.c"))
	if err != nil {
		t.Fatalf("Failed to read directory: %s", err)
	}

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			testCompileAndCompare(t, file, dir)
		})
	}
}

func testCompileAndCompare(t *testing.T, file string, dir string) {
	src := file
	buildDir := filepath.Join(dir, "build")
	config := utils.NewCompileConfig(src, buildDir)
	config.ExecutablePath = filepath.Join(buildDir, filepath.Base(file)+".exe")

	err := myCompile(config)
	if err != nil {
		t.Errorf("Failed to compile %s: %s", src, err)
		return
	}

	executable := filepath.Join(buildDir, "a.exe")
	err = gccCompile(src, executable)
	if err != nil {
		t.Errorf("Failed to compile %s with gcc: %s", src, err)
		return
	}

	myExitCode, err := runExecutable(config.ExecutablePath)
	if err != nil {
		t.Errorf("Failed to run %s: %s", config.ExecutablePath, err)
		return
	}

	gccExitCode, err := runExecutable(executable)
	if err != nil {
		t.Errorf("Failed to run %s: %s", executable, err)
		return
	}

	if myExitCode != gccExitCode {
		t.Errorf("Exit codes do not match:\nmyCompile: %d\ngcc: %d", myExitCode, gccExitCode)
	} else {
		t.Logf("Exit codes match: %d", myExitCode)
	}
}

func myCompile(config *utils.CompileConfigT) error {
	defer func() error {
		r := recover()
		if r == nil {
			return nil
		}

		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("myCompile panic cast error, recover %v", r)
		}
		return err
	}()

	compiler.Compile(config)
	return nil
}

func gccCompile(srcFile, executableFile string) error {
	cmd := exec.Command("gcc", srcFile, "-o", executableFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func runExecutable(executable string) (int, error) {
	cmd := exec.Command(executable)
	err := cmd.Run()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), nil
		}
		return 0, err
	}

	return 0, nil
}
