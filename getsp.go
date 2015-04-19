package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/PuerkitoBio/goquery"
	"encoding/csv"
	"fmt"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "getsp"
	app.Version = Version
	app.Usage = ""
	app.Author = "Souichi"
	app.Email = "sk.cf.msc@gmail.com"
	app.Action = doMain
	app.Run(os.Args)
}

const (
  TARGET = "http://stocks.finance.yahoo.co.jp/stocks/history/?code="
	TAIL = ".T"
)

func doMain(c *cli.Context) {

	URL := TARGET + c.Args()[0] + TAIL

	records := [][]string {}

	doc, _ := goquery.NewDocument(URL)

	doc.Find(".boardFin").Each(func(_ int, s *goquery.Selection) {
		i := []string{}
		s.Find("th").Each(func(_ int, s *goquery.Selection) {
			i = append(i,s.Text())
			})
		records = append(records, i)

		s.Find("tr").Each(func(_ int, s *goquery.Selection) {
			j := []string{}
			s.Find("td").Each(func(_ int, s *goquery.Selection) {
				j = append(j,s.Text())
				})
			records = append(records, j)
		})

		s.Find(".ymuiPagingBottom").Each(func(_ int, s*goquery.Selection) {
				a := s.Text()
				fmt.Println(a)
		})
	})


	if err := os.Mkdir("data", 0777); err != nil {
    fmt.Println(err)
  }

	now := time.Now().Format("2006-01-02T15:04:05Z07:00")
	today := strings.Replace(now,"-","",-1)
	FILENAME := "data/" + c.Args()[0] + "_" + today[0:8] + ".csv"

	csvfile, err := os.Create(FILENAME)
		if err != nil {
						fmt.Println("Error:", err)
						return
		}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	writer.Flush()

}
