package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/paulvollmer/htmltable2csv/scraper"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func usage() {
	fmt.Println("Usage: htmltable2csv [flags]")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  htmltable2csv -url 'https://www.w3schools.com/html/html_tables.asp' -selector '#customers > tbody > tr' -csv data.csv")
	fmt.Println("\nAuthor     : Paul Vollmer")
	fmt.Println("Repository : https://github.com/paulvollmer/htmltable2csv")
	fmt.Println("")
}

func main() {
	flagVersion := flag.Bool("v", false, "Print the version and exit")
	flagSource := flag.String("source", "", "The filepath or website url")
	flagSelector := flag.String("selector", "", "The table css selector")
	flagStart := flag.Int("start", 0, "The row to start collecting data")
	flagStop := flag.Int("stop", 0, "The row to stop collecting data")
	flagTrim := flag.Bool("trim", false, "Trim the whitespace for each table column")
	flagCSV := flag.String("csv", "", "The csv filename. if empty, print csv to stdout")
	flag.Usage = usage
	flag.Parse()

	if *flagVersion {
		fmt.Printf("%v, commit %v, built at %v", version, commit, date)
		os.Exit(0)
	}

	if *flagSource == "" {
		fmt.Println("Flag -source cannot be empty")
		os.Exit(1)
	}

	if *flagSelector == "" {
		fmt.Println("Flag -selector cannot be empty")
		os.Exit(1)
	}

	var err error
	scraper := htmltable2csv.Scraper{}
	scraper.Source = *flagSource
	scraper.Selector = *flagSelector
	scraper.Trim = *flagTrim
	scraper.Start = *flagStart
	scraper.Stop = *flagStop
	_, err = scraper.Scrape()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *flagCSV == "" {
		err = scraper.CSV(os.Stdout)
	} else {
		err = scraper.WriteCSV(*flagCSV)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
