package interfaces_impl

import (
	"net/http"
	"io/ioutil"
	e "../errors_check"
	"bytes"
)

//method - GET, url - https://cgc-api.sbgenomics.com/v2/projects,
//bytes.NewBufferString("{\"name\": \"human_g1k_v413_decoy.dict\"}")

func GetRequestResponse(method, url, body, token string) ([]byte, int) {

	req,err := http.NewRequest(method,url,bytes.NewBufferString(body))
	req.Header.Set("X-SBG-Auth-Token",token)
	req.Header.Set("Content-type","application/json")
	req.Host = "cgc-api.sbgenomics.com"
	req.Proto = "HTTP/1.1"

	e.DefaultErrorCheck(err)

	client := &http.Client{}
	resp,err2 := client.Do(req)
	e.DefaultErrorCheck(err2)

	res_body, err := ioutil.ReadAll(resp.Body)
	e.DefaultErrorCheck(err)

	return res_body, resp.StatusCode
}

