package main

import "flag"
import "fmt"

var (
	sources = flag.String("sources", "", "path to the directory containing examples in markdown")
)

func main() {
	flag.Parse()
	parseFolder(*sources)
}

func parseFolder(path string) {
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
	fmt.Println("HEREEEEEEEEEEEEEEEEEEEEEEEEE")
}
