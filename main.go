package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const arraySize = 300 * 1024 * 1024

func showMemStats() {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)

	fmt.Println(mem.Alloc / 1024)

	showResMemStats()
}

func showResMemStats() {
	pid := fmt.Sprintf("%d", os.Getpid())

	cmd := exec.Command("ps", "-p", pid, "-o", "rss")
	out := bytes.Buffer{}
	cmd.Stdout = &out
	cmd.Run()

	fmt.Println(out.String())
}

func allocateAndRelease() {
	a := [arraySize]int{}
	for i := range a {
		a[i] = i
	}
}

func main() {
	showMemStats()

	allocateAndRelease()

	showMemStats()
	runtime.GC()
	showMemStats()

	fmt.Print("Press enter to quit")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
