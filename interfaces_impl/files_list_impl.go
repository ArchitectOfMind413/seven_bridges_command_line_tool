package interfaces_impl

import (
	f "../formating"
	"../constants"
)

func (P) GetFilesList(l ...string) {
	urlJson, _ := GetRequestResponse("GET", constants.LIST_OF_FILES_IN_PROJECT_URL+"?project="+l[0],"", l[1])
	f.PrintResponseUrl(urlJson)
}
