package main

import "os/exec"

/*Port the following script into go:
    rsync --remove-source-files -a pi@mypi:~/camera/ input/
*/

func runCmd(cmd string, args ...interface{}) {
	cmd := exec.Command(cmd, args...)
	err := cmd.Run()
	must(err)
}
runCmd("rsync", "--remove-source-files", "-a", "pi@mypi:~/camera/", "input/")
