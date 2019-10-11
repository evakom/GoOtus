/*
 * HomeWork-7: envdir utility like envdir
 * Created on 11.10.2019 21:51
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// EnvDirExec runs program with env from given directory
func EnvDirExec(pathProgram, pathEnvDir string) error {

	cmd := exec.Command(pathProgram)

	// replace system env with files env
	cmd.Env = replaceSystemEnvOnFilesEnv(os.Environ(), getEnvFromFiles(pathEnvDir))

	// run and print env
	out, err := cmd.Output()
	if err != nil {
		return err
	} else {
		fmt.Printf("Get output:\n%s", out)
	}

	return nil
}

func getEnvFromFiles(envDir string) []string {

	return []string{"QQQ=qqq1", "VVV=vvv1", "PATH=/EEE/rrr"}
}

func replaceSystemEnvOnFilesEnv(sysEnv, filesEnv []string) []string {
	if !inheritEnv {
		return filesEnv
	}

	hash := make(map[string]string)
	inter := make([]string, 0)

	// hash sys env
	for _, se := range sysEnv {
		name := strings.SplitN(se, "=", 2)
		fmt.Println(name)
		hash[name[0]] = name[1]
	}

	// hash dir env
	for _, fe := range filesEnv {
		name := strings.SplitN(fe, "=", 2)
		hash[name[0]] = name[1]
	}

	// hash -> slice
	for key, val := range hash {
		env := fmt.Sprintf("%s=%s", key, val)
		inter = append(inter, env)
	}

	return inter
}
