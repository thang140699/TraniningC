package main

import (
	download "crawl/utilities"
	"io"
	"log"
	"time"
)

var (
	Time         = time.Now()
	Regex        = `<a.+?\s*href\s*=\s*["\']?([^"\'\s>]+)["\']?`
	NumberOfWork = 100
	mode         string
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
	Year, Month, Day := download.GetTime()

	date := Year + "-" + Month + "-" + Day
	URL := "https://malshare.com/daily/" + date + "/"
	StringHTML, err := getString(URL)
	if err != nil {
		log.Fatalln(err)
		return
	}
	link := download.HandleGetLinkTxT(StringHTML)
	// fmt.Println(link)

	stringHashCode, err := getString(URL + link)

	if err != nil {
		log.Fatalln(err)
		return
	}

	col1s, col2s, col3s, col4s := download.HashCode(stringHashCode)
	pathCol1, pathCol2, pathCol3, pathCol4 := download.HandleCreateFile(Year, Month, Day)

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

}
