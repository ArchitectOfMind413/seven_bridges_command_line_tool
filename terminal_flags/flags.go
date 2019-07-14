package terminal_flags

import (
	"flag"
)

func GenerateFlags() (string, string, string, string, string, string, string, string, string, string) {
	token := flag.String("token", "", "Token for Authentication")

	method := flag.String("method", "", "Method for execution. Go to --help to see the list of all methods.")

	offset := flag.String("offset", "0", "Offset value for project list")
	limit := flag.String("limit", "0", "Limit value for project list")

	projectId := flag.String("project", "", "Project ID")
	fileId := flag.String("file", "", "File ID")
	meta_key := flag.String("mk", "", "Key for updating metadata")
	meta_value := flag.String("mv", "", "Value for key for updating metadata")
	new_file_name := flag.String("nfn","","New name for updating file name")
	destination := flag.String("dest", "", "Download link for file")
	flag.Parse()

	return *token, *method, *offset, *limit, *projectId, *fileId, *meta_key, *meta_value, *new_file_name, *destination
}
