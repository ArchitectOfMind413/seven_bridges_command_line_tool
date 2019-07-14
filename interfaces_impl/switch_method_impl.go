package interfaces_impl

import (
	"log"
	"fmt"
)

var p P

var Token, Method, Offset, Limit, ProjectID, FileID, Meta_key, Meta_value, NewFileName, Destination string
var missingFlags string

var m map[string]string

func ExecuteFlagMethod()  {

	switch Method {
	case "":
		log.Fatal("You didn't wrote method for execution! See list of methods in -help description!")
	case "project_list":

		m=ConvertStringsIntoMapStringString("offset",Offset,"limit",Limit,"token",Token)
		ExecuteIfVarsAreNotEmpty(p.GetProjectsList,Method,m)

	case "files_list":

		m=ConvertStringsIntoMapStringString("project",ProjectID,"token",Token)
		ExecuteIfVarsAreNotEmpty(p.GetFilesList,Method, m)

	case "files_stat":

		m=ConvertStringsIntoMapStringString("file",FileID,"token",Token)
		ExecuteIfVarsAreNotEmpty(p.GetFilesStats,Method, m)

	case "files_update_metadata":

		m=ConvertStringsIntoMapStringString("file",FileID,"mk",Meta_key,"mv",Meta_value,"token",Token)
		ExecuteIfVarsAreNotEmpty(p.UpdateFileMetadata,Method, m)

	case "files_update_name":

		m=ConvertStringsIntoMapStringString("file",FileID,"nfn",NewFileName,"token",Token)
		ExecuteIfVarsAreNotEmpty(p.UpdateFileName,Method, m)

	case "files_download":

		m=ConvertStringsIntoMapStringString("file",FileID,"token",Token,"dest",Destination)
		ExecuteIfVarsAreNotEmpty(p.DownloadFile,Method, m)
	}
}

func ExecuteIfVarsAreNotEmpty(f func(...string), method string, m map[string]string){

	for k,v := range m{
		if v == ""{
			missingFlags +=fmt.Sprintf("Flag <%v> is missing!",k)
		}
	}
	if missingFlags !=""{
		missingFlags +=fmt.Sprintf("Method <%v> can't be executed!\n",method)
		log.Fatal(missingFlags)
	}
	var l []string
	for _,p := range m{
		l=append(l,p)
	}

	switch len(l) {
	case 2:
		f(l[0],l[1])
	case 3:
		f(l[0],l[1],l[2])
	case 4:
		f(l[0],l[1],l[2],l[3])
	}

}

func ConvertStringsIntoMapStringString (s ...string) map[string]string{

	p := make(map[string]string)

	for len(s)!=0{
		p[s[0]]=s[1]
		s = s[2:]
	}

	return p
}