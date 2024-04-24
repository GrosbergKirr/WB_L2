package num3

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func Num3(args []string) int {
	var str Data
	err := str.ReadArgs(args)
	if err != nil {
		return 2
	}
	if err = str.Sort(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type Data struct {
	IsNumeric       bool
	IsReverse       bool
	DeleteDuplicate bool
	Column          int
	Strings         io.ReadCloser
}

func (d *Data) ReadArgs(args []string) error {
	fl := flag.NewFlagSet("Sort", flag.ContinueOnError)
	fl.IntVar(&d.Column, "k", 1, "sort via a column")
	fl.BoolVar(&d.IsNumeric, "n", false, "compare according to string numerical value")
	fl.BoolVar(&d.IsReverse, "r", false, "reverse the result of comparisons")
	fl.BoolVar(&d.DeleteDuplicate, "u", false, "delete duplicate strings")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	stat, _ := os.Stdin.Stat()
	if stat.Mode() == 0 {
		d.Strings = os.Stdin
		return nil
	}

	file, err := os.Open(fl.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", fl.Arg(0), err)
		return err
	}
	d.Strings = file

	return nil
}

func (d *Data) Sort() error {
	defer d.Strings.Close()
	data := make([]string, 0)

	scanner := bufio.NewScanner(d.Strings)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if d.Column == 1 && !d.IsNumeric {
		data = d.sort(data)
		writeToOutput(data)
		return nil
	}

	data = d.sortColumns(data)
	writeToOutput(data)

	return nil
}

func (d *Data) sort(data []string) []string {
	if d.IsReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	} else {
		sort.Strings(data)
	}

	if d.DeleteDuplicate {
		data = delDuplicate(data)
	}

	return data
}

func (d *Data) sortColumns(data []string) []string {
	t := stringTable{
		data:      make([][]string, 0, len(data)),
		column:    d.Column - 1,
		isNumeric: d.IsNumeric,
	}

	for _, v := range data {
		t.data = append(t.data, strings.Fields(v))
	}

	if d.IsReverse {
		sort.Sort(sort.Reverse(t))
	} else {
		sort.Sort(t)
	}

	for i, v := range t.data {
		data[i] = strings.Join(v, " ")
	}

	if d.DeleteDuplicate {
		data = delDuplicate(data)
	}

	return data
}
