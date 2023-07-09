package main

import (
	"bytes"
	"os"
	"os/exec"
)

func main() {
	/* What to do
			* 1. Define base directory
			* 2. Search base level to know how many go routines to launch
			* 3. Launch go routines search max depth (with find)
			* 4. Routines that have a hit for maxdepth 1 will return
			* 4a. Routines that do not have a hit will die (close)
		     * 5. Routines that have a hit will count
			    * 5a. The `x` of hits will launch `x` routines (make this ia
			    *recursive function until hit or EOF exit graceful with note)
	Possible traps:
			  Not terminating all channels immediately there is a desired hit
	*/
	word := os.Args
	search(word[:])
}

//menu for search -- if locally compiled change this to your notifier of choice
var menu = exec.Command("dmenu")

/*Need to make a pipe (channel per number output of count)
send channel (blocks)
receive channel
*/

func search(w []string) string {
	//find search word
	find := exec.Command("find", -w[], "type", "f", "-maxdepth", "-1")
	var buf bytes.Buffer
	find.Stdout = &buf
	return find
}
