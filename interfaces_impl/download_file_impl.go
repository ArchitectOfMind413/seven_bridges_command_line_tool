package interfaces_impl

import (
	"../model"
	"net/http"
	"os"
	"io"
	"encoding/json"
	"../constants"
	f "../formating"
	"../errors_check"
	"fmt"
)

func (P) DownloadFile(l ...string) {

	urlJson, statusCode := GetRequestResponse("GET", constants.GET_FILE_DOWNLOAD_ADDRESS_URL+l[0]+"/download_info","", l[1])

	if statusCode ==200{
		url := new(model.UrlDownload)
		json.Unmarshal(urlJson,&url)
		errors_check.DirExist(l[2][:f.LastIndexOfChar(l[2],"/")+1])
		DownloadFileFromUrl(l[2],url.Address)
	}else{
		f.PrintResponseUrl(urlJson)
	}

}

func DownloadFileFromUrl(filepath string, url string){

	resp, err := http.Get(url)
	errors_check.Fatal("Can't get resource from URL!", err)
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	errors_check.Fatal("Can't create file path!", err)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	errors_check.Fatal("Can write []byte from resource to specified file path!", err)

	fmt.Printf("File has been downloaded successfully into path: <%v>\n", filepath)
}
