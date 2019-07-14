package interfaces_impl

import (
	"../constants"
	f "../formating"
)

func (P) UpdateFileMetadata(l ...string)  {
	urlJson, _ := GetRequestResponse("PATCH", constants.UPDATE_FILE_METADATA_URL+l[0],"{\"metadata\": {\""+l[1]+"\": \""+l[2]+"\"}}", l[3])
	f.PrintResponseUrl(urlJson)
}
