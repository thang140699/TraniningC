package main

import (
	download "crawl/utilities"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

var (
	Time         = time.Now()
	Regex        = `<a.+?\s*href\s*=\s*["\']?([^"\'\s>]+)["\']?`
	NumberOfWork = 100
	mode         string
	configPrefix string
)

func getString(URL string) (string, error) {
	resHTMl, err := download.RequestURL(URL)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resHTMl.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	// for date
	const layout = "2006-01-02 "
	FirstDay, _ := time.Parse(layout, "2023-02-15")
	fmt.Println(FirstDay)

	Year, Month, Day := download.GetTime()
	date := Year + "-" + Month + "-" + Day

	Yesterday, _ := time.Parse(layout, date)
	fmt.Println(Yesterday)
	for i := FirstDay; i != Yesterday; {
		fmt.Println(i.Format(layout))

		// get urrl
		URL := "https://malshare.com/daily/" + i.Format(layout) + "/"
		StringHTML, err := getString(URL)
		if err != nil {
			log.Fatalln(err)
			return
		}

		link := download.HandleGetLinkTxT(StringHTML)

		stringHashCode, err := getString(URL + link)

		if err != nil {
			log.Fatalln(err)
			return
		}

		col1s, col2s, col3s, col4s := download.HashCode(stringHashCode)

		monthInt := int(i.Month())
		stringMonth := strconv.Itoa(monthInt)

		if monthInt < 10 {
			stringMonth = "0" + stringMonth
		}

		year, month, day := strconv.Itoa(i.Year()), stringMonth, strconv.Itoa(i.Day()-1)
		pathCol1, pathCol2, pathCol3, pathCol4 := download.HandleCreateFile(year, month, day)


		err = download.HandleWriteFile(col1s, pathCol1)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(col2s, pathCol2)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(col3s, pathCol3)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(col4s, pathCol4)
		if err != nil {
			log.Fatal(err)
		}
		i = i.Add(time.Hour * 24)
	}
	// var config Config
	// err := download.LoadEnvFromFile(&config, configPrefix, config)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}


// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	flag.StringVar(&configPrefix, "config", "configPrefix ")
// }
