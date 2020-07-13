package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery" // Used for Scraping Web
	// "encoding/json"
	 "html/template"
	  "github.com/jung-kurt/gofpdf" /// Used to create pdf File
	 "io"
	 "strings"
	// "io/ioutil"
	  "encoding/json"
	  //     "bytes"
   "log"
 //  "path/filepath"
   "os"
   "net/http"
   "strconv"
     //"time"
     "path"
)

// type webData struct { //
// 	Body string
// }
type ScrapePage struct { // Used to save The data from Scraping the web
	Page string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
fmt.Println(r.URL.String())
	if r.URL.Path == "/" {
		fmt.Println("HOME")

	webScrape := ScrapePage{""}
	fp := path.Join("templates", "home.html")
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
		       // toString = toString + s.Find("h1").Text()
		        toString = toString +  strings.TrimSuffix(s.Find("p").Text(), "\n")//s.Find("p").Text()  // Here we are ensuring we don't add any \n

		       // fmt.Printf("Review %d: %s - %s\n", i, band, title)
		    })
		}


	webScrape,err := json.Marshal(ScrapePage{toString}) // Serializing to json object
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
	//fp := path.Join("templates", "home.html")
	//tmpl, err := template.ParseFiles(fp)
   // if err != nil {
	//    http.Error(w, err.Error(), http.StatusInternalServerError)
	//    return
   // }
   // if err := tmpl.Execute(w, webScrape); err != nil {
   //         http.Error(w, err.Error(), http.StatusInternalServerError)
   //     }
   w.Header().Set("Content-Type", "application/json")
     w.Write(webScrape)



}




}

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit download")
	//var file = "response.pdf"
	fmt.Println(r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		// fmt.Println("WE ARE IN GET", r.Body)
	 //
		// decoder := json.NewDecoder(r.Body)
		// u:= Text{ }
		// err := decoder.Decode(&u)
	 //  if err != nil {
	 //    log.Fatalln(err)
	 // }

		toString :=  r.Form["data"]//r.URL.Query()["pValue"][0]

		actualString := ""
		for _,value := range toString{
			actualString = actualString + value//strings.TrimSuffix(value, "\n")
		}
		fmt.Println(toString)
		// err2 := WriteToFile("response.pdf", "Hello")
		//    if err2 != nil {
		// 		log.Fatal(err2)
		// }

			err5 := GeneratePdf("response.pdf", actualString)
    	if err5 != nil {
        	panic(err5)
    	}



			w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("lkdalda.pdf"))
			w.Header().Set("Content-Type", "application/octet-stream")
					http.ServeFile(w, r, "/Users/josearellanes/makeUtility/response.pdf")

			 const TmpDir = "/Users/josearellanes/makeUtility/response.pdf";
// 		downloadBytes, err := ioutil.ReadFile(file)
//
// 		if err != nil {
// 				fmt.Println(err)
// 		}
//
// 		// set the default MIME type to send
// 		mime := http.DetectContentType(downloadBytes)
//
// 		fileSize := len(string(downloadBytes))
//
// 		// Generate the server headers
// 		w.Header().Set("Content-Type", mime)
// 		w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
// 		w.Header().Set("Expires", "0")
// 		w.Header().Set("Content-Transfer-Encoding", "binary")
// 		w.Header().Set("Content-Length", strconv.Itoa(fileSize))
// 		w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
//
// 		//b := bytes.NewBuffer(downloadBytes)
// 		//if _, err := b.WriteTo(w); err != nil {
// 		//              fmt.Fprintf(w, "%s", err)
// 		//      }
//
// 		// force it down the client's.....
// 		http.ServeContent(w, r,file, time.Now(), bytes.NewReader(downloadBytes))
// fmt.Println("should be completed")

			// http.ServeFile(w, r, filepath.Join( TmpDir, "/response.pdf" ))
			//
			// Openfile, err := os.Open("response.pdf")
			// 	defer Openfile.Close() //Close after function return
			// 	if err != nil {
			// 		//File not found, send 404
			// 		http.Error(w, "File not found.", 404)
			// 		return
			// 	}
			//
			// 	//File is found, create and send the correct headers
			//
			// 	//Get the Content-Type of the file
			// 	//Create a buffer to store the header of the file in
			// 	FileHeader := make([]byte, 512)
			// 	//Copy the headers into the FileHeader buffer
			// 	Openfile.Read(FileHeader)
			// 	//Get content type of file
			// 	FileContentType := http.DetectContentType(FileHeader)
			//
			// 	//Get the file size
			// 	FileStat, _ := Openfile.Stat()                     //Get info from file
			// 	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
			//
			// 	//Send the headers
			// 	Filename:= "TheEnd"
			// 	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
			// 	w.Header().Set("Content-Type", FileContentType)
			// 	w.Header().Set("Content-Length", FileSize)
			//
			// 	//Send the file
			// 	//We read 512 bytes from the file already, so we reset the offset back to 0
			// 	Openfile.Seek(0, 0)
			// 	io.Copy(w, Openfile) //'Copy' the file to the client
			// 	return




	}


	// doc, err := goquery.NewDocument("") // Webscrape our own
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	// 	doc.Find("#Data").Each(func(i int, s *goquery.Selection) {
	// 		//fmt.Println(s.Find("p").Text())
	// 		toString = toString + s.Find("p").Text()
	//        // fmt.Printf("Review %d: %s - %s\n", i, band, title)
	//     })


}
func main() {


	http.HandleFunc("/", handler)
	http.HandleFunc("/search", handler)
	http.HandleFunc("/download", handlerDownload)

	    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))







	// Instantiate default collector
	// c := colly.NewCollector()
	//
	// // On every a element which has href attribute call callback
	// c.OnHTML("body", func(e *colly.HTMLElement) {
    //             //bodyData := e.Attr("body")
	// 			fmt.Println(e.Text)
	// 			jsonStruct := &webData{Body:e.Text}
	// 	// Print link
    //             //fmt.Printf("Text found: %q -> %s\n", e.Text, link)
	// 		result, err := json.Marshal(jsonStruct)
	// 		   if err != nil {
	// 			   fmt.Println(err)
	// 			   return
	// 		   }
	//
	// 		  // fmt.Println(result)
	// 		   err2 := WriteToFile("result.txt", string(result))
	// 	   if err2 != nil {
	//    		log.Fatal(err2)
   	// 	}
	// })
	//
	//
	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })
	//
	// // Start scraping on https://hackerspaces.org
	// c.Visit("https://news.ycombinator.com/")


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

//     pdf := gofpdf.New("P", "mm", "A4", "")
//     pdf.SetFont("Arial", "B", 3)
// pdf.AddPage()
// pdf.SetMargins(10, 10, 10)
// lines := pdf.SplitLines([]byte(values), 100.0)
    // CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
    //pdf.CellFormat(190, 7, "", "0", 0, "CM", false, 0, "")

    // ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
    // pdf.ImageOptions(
    //     "avatar.jpg",
    //     80, 20,
    //     0, 0,
    //     false,
    //     gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
    //     0,
    //     "",
    // )
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
	//ht := float64(len(lines)) * lineHt
	// y := (297.0 - ht) / 2.0
	// pdf.SetDrawColor(128, 128, 128)
	// pdf.SetFillColor(255, 255, 210)
	// x := (210.0 - (wd + 40.0)) / 2.0
	// pdf.Rect(x, y-20.0, wd+40.0, ht+40.0, "FD")
	// pdf.SetY(y)
	for _, line := range lines {
	    pdf.CellFormat(190.0, lineHt, string(line), "", 1, "C", false, 0, "")
	}
	// fileStr := example.Filename("Fpdf_Splitlines")
	// err := pdf.OutputFileAndClose(fileStr)
	// example.Summary(err, fileStr)
    return pdf.OutputFileAndClose(filename)
}
