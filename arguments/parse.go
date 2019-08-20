package arguments

import "flag"

//Parse - parse arguments
func Parse() string {
	var keyword *string
	keyword = flag.String("keyword", "", "keyword")
	flag.Parse()
	return *keyword
}
