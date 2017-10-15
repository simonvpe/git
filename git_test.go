package git

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func createTestDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "gitcommand-test-dir")
	if err != nil {
		t.Error("Failed to create temporary directory")
	}
	return dir
}

func gitFail(t *testing.T, dir string, args ...string) {
	arglist := append(
		[]string {},
		args...,
	)
	if _, err := Run(dir, args...); err == nil {
		t.Errorf("Expected git command to fail (git %#v)", arglist)
	}
}

func gitSuccess(t *testing.T, dir string, args ...string) []byte {
	arglist := append(
		[]string {},
		args...,
	)

	output, err := Run(dir, args...)
	if err != nil {
		t.Errorf("Failed to run git command (git %#v)", arglist)
	}
	return output
}

func TestGitInit(t *testing.T) {
	initialized := createTestDir(t)
	defer os.RemoveAll(initialized)
	
	uninitialized := createTestDir(t)
	defer os.RemoveAll(uninitialized)

	{
		output := gitSuccess(t, initialized, "init")
		match, err := regexp.Match("Initialized empty Git repository in .*", output)
		if match != true || err != nil {
			t.Errorf("Failed to match regex (%q)", string(output))
		}
	}

	gitSuccess(t, initialized, "status")
	gitFail(t, uninitialized, "status")
}

func TestGitAdd(t *testing.T) {
	dir := createTestDir(t)
	defer os.RemoveAll(dir)

	testfile := dir + "/testfile"
	{
		output := gitSuccess(t, dir, "init")
		ioutil.WriteFile(testfile, output, 0777)
	}

	gitSuccess(t, dir, "add", testfile)

	{
		output := gitSuccess(t, dir, "status")
		match, err := regexp.Match("new file: *testfile", output)
		if match != true || err != nil {
			t.Errorf("Failed to match regex (%q)", string(output))
		}
	}

}
