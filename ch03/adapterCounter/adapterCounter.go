package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}

func main() {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := &Counter{
		Writer: pipeWriter,
	}
	tee := io.TeeReader(pipeReader, f)
	go func() {
		io.Copy(os.Stdout, tee)
	}()
	c.Count(5)

}
