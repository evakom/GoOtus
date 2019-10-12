/*
 * HomeWork-7: envdir utility like envdir
 * Created on 12.10.2019 12:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// short name of the test program to run with params
const EXECNAME = "envdir"

var testCases = []struct {
	envDir      string
	envVars     []string
	description string
}{
	{
		"TestDir",
		[]string{"QQQ=AAA", "WWW=SSS"},
		"no inherit 1",
	},
	{
		"TestDir1",
		[]string{"EEE=DDD", "RRR=FFF", "ttt=ggg"},
		"no inherit 2",
	},
}

func TestEnvDirExec(t *testing.T) {

	execFile := getExecFile()

	for _, test := range testCases {

		//cleanEnvDir(test.envDir)
		//generateEnvDir(test.envDir, test.envVars)

		// you may test exec directly by uncomment this line instead bellows
		//out, err := exec.Command(execFile, "-env", test.envDir, "-exec", execFile).Output()

		out := new(strings.Builder)

		err := EnvDirExec(out, execFile, test.envDir, false)
		if err != nil {
			t.Errorf("FAIL '%s' - TestEnvDirExec() returns error\n %s, expected no error.",
				test.description, err)
			continue
		}

		fmt.Printf("%s", out)

		t.Logf("PASS TestEnvDirExec - %s", test.description)

		// end clean if not need results
		// cleanEnvDir(test.envDir)
	}
}

func getExecFile() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Can't get current test directory!", err)
	}
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	return filepath.Join(dir, EXECNAME+ext)
}

func generateEnvDir(envDir string, envVars []string) {
	err := os.Mkdir(envDir, 0777)
	if err != nil {
		log.Fatalln("Can't create test directory!", err)
	}
	for _, ev := range envVars {
		fileName := strings.SplitN(ev, "=", 2)
		file, err := os.Create(path.Join(envDir, fileName[0]))
		if err != nil {
			log.Fatalln("Can't create test file!", err)
		}
		_, err = file.Write([]byte(fileName[1]))
		if err != nil {
			log.Fatalln("Can't write test data to file!", err)
		}
		err = file.Close()
		if err != nil {
			log.Fatalln("Can't close test file!", err)
		}
	}
}

func cleanEnvDir(envDir string) {
	err := os.RemoveAll(envDir)
	if err != nil {
		log.Fatalln("Can't delete test directory!", err)
	}
}
