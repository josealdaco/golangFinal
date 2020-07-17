package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery" // Used for Scraping Web
	// "encoding/json"
	 "html/template"
	 // "github.com/jung-kurt/gofpdf" /// Used to create pdf File
	 "io"
	 "strings"
	// "golang.org/x/net/html"
	"io/ioutil"
	  "encoding/json"
	      "bytes"
   "log"
 //  "path/filepath"
   "os"
   "net/http"
   "strconv"
     "time"
//     "path"
)

type ScrapePage struct { // Used to save The data from Scraping the web
	Page string
}


func handler(w http.ResponseWriter, r *http.Request) {
// This will handle our home and  search links
fmt.Println(r.URL.String())
	if r.URL.Path == "/" {
		fmt.Println("HOME")

	webScrape := ScrapePage{""}
	fp := "home.html"//path.Join("templates", "home.html")
	tmpl, err := template.ParseFiles(fp) // Parsing our home html which we will use to render data
   if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
   }
   if err := tmpl.Execute(w, webScrape); err != nil { // parsing data from our webScrape to template
           http.Error(w, err.Error(), http.StatusInternalServerError)
       }

}
if r.URL.Path == "/search" {


	fmt.Println("Inside search")
	toString:= ""
	 checker := r.URL.Query()["SearchValue"][0]
	 fmt.Println("Check the URL", checker)
	doc, err := goquery.NewDocument(checker)
	    if err != nil {
	        //log.Fatal(err)
	    } else{

			doc.Find("div").Each(func(i int, s *goquery.Selection) {
				//for now we are simply parsing through every div in the site and extracting the values from p
		        toString = toString +  strings.TrimSuffix(s.Find("p").Text(), "\n")//s.Find("p").Text()  // Here we are ensuring we don't add any \n

		    })
		}


	webScrape,err := json.Marshal(ScrapePage{toString}) // Serializing to json object
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
   w.Header().Set("Content-Type", "application/json")
     w.Write(webScrape)



}




}

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit download")
	fmt.Println(r.Method)
	if r.Method == "GET" {
		r.ParseForm()


		toString :=  r.Form["data"]//r.URL.Query()["pValue"][0]

		actualString := ""
		for _,value := range toString{
			actualString = actualString + value//strings.TrimSuffix(value, "\n")
		}

			err5 := GeneratePdf("response.pdf", actualString) // Here we want to generate THE PDF File that will be sent to the browser for download
    	if err5 != nil {
        	panic(err5)
    	}

		 		downloadBytes, err := ioutil.ReadFile("response.pdf")

				if err != nil {
						fmt.Println(err)
				}

				// set the default MIME type to send
				mime := http.DetectContentType(downloadBytes)

				fileSize := len(string(downloadBytes))

				// Generate the server headers
				w.Header().Set("Content-Type", mime)
				w.Header().Set("Content-Disposition", "attachment; filename="+"responseFile"+"")
				w.Header().Set("Expires", "0")
				w.Header().Set("Content-Transfer-Encoding", "binary")
				w.Header().Set("Content-Length", strconv.Itoa(fileSize))
				w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

				b := bytes.NewBuffer(downloadBytes)
				if _, err := b.WriteTo(w); err != nil {
				             fmt.Fprintf(w, "%s", err)
				     }

				// force it down the client's.....
				http.ServeContent(w, r,"response.pdf", time.Now(), bytes.NewReader(downloadBytes))



	}



}
func main() {
// Connect all the handlers and links here
	http.HandleFunc("/", handler)
	http.HandleFunc("/search", handler)
	http.HandleFunc("/download", handlerDownload)

	  log.Fatal(http.ListenAndServe( ":80", nil))//+ os.Getenv("PORT"), nil))//":8080",nil)) //+ os.Getenv("PORT"), nil))


}

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename) ///
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}
func GeneratePdf(filename string, values string) error {

	const (
	    fontPtSize = 18.0
	    wd         = 100.0
	)
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf.SetFont("Times", "", fontPtSize)
	_, lineHt := pdf.GetFontSize()
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)
	lines := pdf.SplitLines([]byte(values), wd)

	for _, line := range lines {
	    pdf.CellFormat(190.0, lineHt, string(line), "", 1, "C", false, 0, "")
	}

    return pdf.OutputFileAndClose(filename)
}
