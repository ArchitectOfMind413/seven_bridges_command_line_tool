package interfaces_impl

import (
	"../constants"
	f "../formating"
)

func (P) UpdateFileName(l ...string)  {
	urlJson, _ := GetRequestResponse("PATCH", constants.UPDATE_FILE_NAME_URL+l[0],"{\"name\": \""+l[1]+"\"}", l[2])
	f.PrintResponseUrl(urlJson)
}

