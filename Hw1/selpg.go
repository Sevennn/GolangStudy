package main

import (
    "flag"
		"fmt"
		"os"
		"bufio"
		"os/exec"
)

type selpg  struct {
		s int
		e int 
		page_len int
		page_type int
		in_filename string
		print_dest string
}
func main() {
		sa := selpg {
			s:-1,
			e:-1,
			page_len:72,
			page_type:1,
			in_filename: "",
			print_dest:"",
		}
		flag.IntVar(&sa.s,"s", -1, "specify start page(>=1)")
		flag.IntVar(&sa.e,"e", -1, "specify end page(>=s)")
		flag.IntVar(&sa.page_len, "l", 72, "specify length of one page")
		page_type := flag.Bool("f", false, "-f means type 2 page and you can't set -f and page length at the same time")
		print_dest := flag.String("d", "", "specify print dest.")
		flag.Usage = usage
    flag.Parse()

		if sa.s == -1 || sa.e == -1 || sa.s > sa.e || sa.s < 1 || sa.e < 1{
			flag.Usage()
			return
		}
		
		if sa.page_len != 72 && *page_type == true {
			flag.Usage()
		}

		if *page_type == true {
			sa.page_type = 2
		}

		if *print_dest  != "" {
			sa.print_dest = *print_dest
		}

		if len(flag.Args()) > 1 {
			flag.Usage()
		}

		if len(flag.Args()) == 1 {
			sa.in_filename = flag.Args()[0]
		}

		if sa.page_type == 1 {
			type_1_handler(sa, sa.in_filename != "", sa.print_dest != "");
		} else {
			type_2_handler(sa, sa.in_filename != "", sa.print_dest != "");
		}
}

func type_1_handler(sa selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err:= cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cur_page := 1
	cur_lines := 0
	if file {
		file_in, err := os.OpenFile(sa.in_filename,os.O_RDWR,os.ModeType)
		defer file_in.Close()
		if err != nil {
			panic(err)
			return
		}
    line := bufio.NewScanner(file_in)
    for line.Scan() {
			if cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte(line.Text()+"\n"))
				stdin.Write([]byte(line.Text()+"\n"))
			}
			cur_lines++;
			if cur_lines %= sa.page_len; cur_lines == 0 {
				cur_page++;
			}
    }
	} else {
		tmp_s := bufio.NewScanner(os.Stdin)
		for tmp_s.Scan() {
			if cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte(tmp_s.Text()+"\n"))
				stdin.Write([]byte(tmp_s.Text()+"\n"))
			}
			cur_lines++;
			if cur_lines %= sa.page_len; cur_lines == 0 {
				cur_page++;
			}
		}
	}
	if cur_page < sa.e {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {
		stdin.Close()
		cmd.Stdout = os.Stdout;
		cmd.Start()
	}

}

func type_2_handler(sa selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err:= cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cur_page := 1
	if file {
		file_in, err := os.OpenFile(sa.in_filename,os.O_RDWR,os.ModeType)
		defer file_in.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(file_in)
    for line.Scan() {
			flag := false
			for _,c := range line.Text() {
				if c == '\f' {
					if cur_page >= sa.s && cur_page <= sa.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					cur_page++;
				} else {
					if cur_page >= sa.s && cur_page <= sa.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
    }
	} else {
		tmp_s := bufio.NewScanner(os.Stdin)
		for tmp_s.Scan() {
			flag := false
			for _,c := range tmp_s.Text() {
				if c == '\f' {
					if cur_page >= sa.s && cur_page <= sa.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					cur_page++;
				} else {
					if cur_page >= sa.s && cur_page <= sa.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	}
	if cur_page < sa.e {
		fmt.Fprintf(os.Stderr, "This file is too short to reach end page\n")
	}
	if pipe {

		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`usage: [-s start page(>=1)] [e end page(>=s)] [-l length of page(default 72)] [-f type of file(default 1)] [-d dest] [filename specify input file]
`)
}