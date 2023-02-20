package download

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Time         = time.Now()
	Regex        = `<a.+?\s*href\s*=\s*["\']?([^"\'\s>]+)["\']?`
	NumberOfWork = 100
	url          = os.Getenv("URL")
)

func RequestURL(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 15,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fail: prepareing request get to  %s\n", url))
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fail: making get to %\n", url))
	}
	return response, nil
}
func GetTime() (string, string, string) {
	Year, Month, Day := Time.Date()
	monthInt := int(Month)
	stringMonth := strconv.Itoa(monthInt)
	if monthInt < 10 {
		stringMonth = "0" + stringMonth
	}
	return strconv.Itoa(Year), stringMonth, strconv.Itoa(Day - 1)
}
func HandleGetLinkTxT(htmlstring string) (result string) {
	makeRegex := regexp.MustCompile(Regex)
	tag := makeRegex.FindAllStringSubmatch(htmlstring, -1)
	tagA := findTagB(tag, "all")
	arrTagSplit := strings.Split(tagA, `"`)
	result = arrTagSplit[len(arrTagSplit)-2]
	return
}
func HashCode(stringHashCode string) (col1s, col2s, col3s, col4s []string) {
	reader := strings.NewReader(stringHashCode)
	var col1, col2, col3, col4 string

	for {
		_, err := fmt.Fscan(reader, &col1, &col2, &col3, &col4)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		col1s = append(col1s, col1)
		col2s = append(col2s, col2)
		col3s = append(col3s, col3)
		col4s = append(col4s, col4)
	}
	return
}
func findTagB(tags [][]string, substr string) (result string) {
	for _, v := range tags {
		if strings.Contains(v[0], substr) {
			result = v[0]
			break
		}
	}
	return
}
func HandleCreateFile(Year, Month, Day string) (string, string, string, string) {
	pathFolder := Year + "/" + Month + "/" + Day
	err := os.MkdirAll(pathFolder, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	pathCol1 := pathFolder + "/col1.txt"
	pathCol2 := pathFolder + "/col2.txt"
	pathCol3 := pathFolder + "/col3.txt"
	pathCol4 := pathFolder + "/col4.txt"
	return pathCol1, pathCol2, pathCol3, pathCol4
}
func HandleWriteFile(array []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	queue := makeQueue(array)
	for i := 1; i < NumberOfWork; i++ {
		go saveToFile(queue, file)
	}
	return nil
}
func makeQueue(array []string) <-chan string {
	queue := make(chan string, 100)
	go func() {
		for _, v := range array {
			queue <- v
		}
		close(queue)
	}()

	return queue
}

func saveToFile(queue <-chan string, file *os.File) {
	for v := range queue {
		_, err := file.WriteString(v + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
