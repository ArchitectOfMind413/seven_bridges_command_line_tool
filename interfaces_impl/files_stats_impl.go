package interfaces_impl

import (
	"../constants"
	f "../formating"
)

func (P) GetFilesStats(l ...string)  {
	urlJson, _ := GetRequestResponse("GET", constants.FILES_STATS_URL+l[0],"", l[1])
	f.PrintResponseUrl(urlJson)
}
