package main

import (
	t "./terminal_flags"
	i "./interfaces_impl"
)


func main() {
	i.Token, i.Method, i.Offset, i.Limit, i.ProjectID, i.FileID, i.Meta_key, i.Meta_value, i.NewFileName, i.Destination = t.GenerateFlags()
	i.ExecuteFlagMethod()
}



