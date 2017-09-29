package main

import (
	_ "crawler/lib"

	"github.com/henrylee2cn/pholcus/exec"
)

func main() {
	exec.DefaultRun("web")
}
