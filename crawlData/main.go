package main

import (
	download "crawl/utilities"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var (
	Time         = time.Now()
	Regex        = `<a.+?\s*href\s*=\s*["\']?([^"\'\s>]+)["\']?`
	NumberOfWork = 100
	mode         string
	container    *Container

	configPrefix string
	configSource string
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

		md5s, sha1s, sha256s, ds := download.HashCode(stringHashCode)

		monthInt := int(i.Month())
		stringMonth := strconv.Itoa(monthInt)

		if monthInt < 10 {
			stringMonth = "0" + stringMonth
		}

		year, month, day := strconv.Itoa(i.Year()), stringMonth, strconv.Itoa(i.Day()-1)
		pathMD5, pathSHA1, pathSHA256, pathD := download.HandleCreateFile(year, month, day)

		err = download.HandleWriteFile(md5s, pathMD5)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(sha1s, pathSHA1)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(sha256s, pathSHA256)
		if err != nil {
			log.Fatal(err)
		}

		err = download.HandleWriteFile(ds, pathD)
		if err != nil {
			log.Fatal(err)
		}
		i = i.Add(time.Hour * 24)
	}
	var config Config
	err := download.LoadEnvFromFile(&config, configPrefix, configSource)
	if err != nil {
		log.Fatalln(err)
	}
	container, err = NewContainer(config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server is running at : " + config.Binding)

	http.ListenAndServe(config.Binding, NewAPIv1(container))

}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&configPrefix, "config", "configPrefix", "crawl")
	flag.StringVar(&configSource, "configSource", ".env", "config source")
}
