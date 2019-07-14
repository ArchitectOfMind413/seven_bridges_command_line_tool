package interfaces_impl

import (
	"fmt"
	"../model"
	"../constants"
	f "../formating"
	"encoding/json"
)
//"https://cgc-api.sbgenomics.com/v2/projects
//func GetRequestResponse(method, url, body, token string) string {

func (P) GetProjectsList(l ...string)  {

	urlJson, statusCode := GetRequestResponse("GET", constants.LIST_OF_PROJECTS_URL+"?offset="+l[0]+"&limit="+l[1]+"&fields=id,name","", l[2])

	if statusCode ==200{
		p := model.ProjectList{}
		json.Unmarshal(urlJson,&p)

		fmt.Println(fmt.Sprintf("%-40s %-40s \n","ID","NAME"))

		for _,v := range p.ItemsList {
			fmt.Println(fmt.Sprintf("%-40s %-40s",v.Id,v.Name))
		}
	}else{
		f.PrintResponseUrl(urlJson)
	}
}

/*
func GetRequestResponse(method, url, body, token string) string {

	req,err := http.NewRequest(method,"https://cgc-api.sbgenomics.com/v2/projects",nil)
	req.Header.Set("X-SBG-Auth-Token","0a4ac2b8ef004ac2a8efde9a85640bb9")
	req.Header.Set("Content-type","application/json")
	req.Host = "cgc-api.sbgenomics.com"
	req.Proto = "HTTP/1.1"

	e.DefaultErrorCheck(err)

	client := &http.Client{}
	resp,err2 := client.Do(req)
	fmt.Println(resp)						//&{200 OK 200 HTTP/2.0 2 0 map[Accept:[application/json].......
	e.DefaultErrorCheck(err2)

	res_body,err := ioutil.ReadAll(resp.Body)
	log.Println(string(res_body))				//2019/07/08 15:19:56 {"href":"https://cgc-api.sbgenomics.com/v2/projects?offset=0&limit=1".......

	return string(res_body)
}
 */

