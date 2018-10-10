package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	flag "github.com/spf13/pflag"
)

type args_type struct {
	start_page, end_page, page_len int
	page_type                      bool
	filename, dest                 string
}

var (
	pname string
)

func process_input(args args_type) {
	//open the file
	var infile *os.File
	var err error
	if args.filename == "" {
		infile = os.Stdin
	} else {
		infile, err = os.Open(args.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open the file: %s\n", args.filename)
			os.Exit(1)
		}
	}
	//execute the command
	var cmd *exec.Cmd
	var outfile *os.File
	if args.dest != "" {
		cmd = exec.Command("/usr/bin/lp", fmt.Sprintf("-d%s", args.dest))
		r, w, err := os.Pipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open a pipe for %s\n", args.dest)
		}
		cmd.Stdin = r
		outfile = w
	} else {
		outfile = os.Stdout
	}

	//read the file
	reader := bufio.NewReader(infile)
	writer := bufio.NewWriter(outfile)
	if args.page_type {
		for i := 1; i <= args.end_page; i++ {
			line, err := reader.ReadString('\f')
			line = strings.TrimSpace(line)
			if err != nil {
				if err == io.EOF && i == args.end_page {
					break
				}
				fmt.Fprintf(os.Stderr, "The range of page is illegal!\n")
				os.Exit(2)
			}
			if i >= args.start_page {
				writer.WriteString(line + "\n")
				writer.Flush()
			}
		}
	} else {
		for i := 1; i <= args.end_page; i++ {
			for j := 0; j < args.page_len; j++ {
				line, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF && i == args.end_page {
						break
					}
					fmt.Fprintf(os.Stderr, "The range of page is illegal!\n")
					os.Exit(3)
				}
				if i >= args.start_page {
					writer.WriteString(line)
					writer.Flush()
				}
			}
		}
	}
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if cmd != nil {
		err := cmd.Run()
		if err != nil {
			fmt.Println(string(w.Bytes()))
		}
	}
}

func main() {
	pname = os.Args[0]

	args := args_type{}

	flag.IntVar(&args.start_page, "s", -1, "Start page")
	flag.IntVar(&args.end_page, "e", -1, "End page")
	flag.IntVar(&args.page_len, "l", 72, "Page length")
	flag.BoolVar(&args.page_type, "f", false, "Page type")
	flag.StringVar(&args.dest, "d", "", "The destination of pipe")

	flag.Parse()
	if len(flag.Args()) != 0 {
		args.filename = flag.Args()[0]
	}

	if args.start_page < 1 {
		fmt.Fprintf(os.Stderr, "The number of start page %d should be positive\n", args.start_page)
		os.Exit(4)
	}
	if args.end_page < 1 {
		fmt.Fprintf(os.Stderr, "The number of end page %d should be positive\n", args.end_page)
		os.Exit(5)
	}
	if args.start_page > args.end_page {
		fmt.Fprintf(os.Stderr, "The number of start page should be greater than or equal to the end page's!\n")
		os.Exit(6)
	}
	if args.page_len < 1 {
		fmt.Fprintf(os.Stderr, "The number of line in a page %d should be positive\n", args.page_len)
		os.Exit(7)
	}

	process_input(args)
}
