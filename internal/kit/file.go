package kit

import "io/ioutil"

func ReadAsString(filePath string) string {
	str, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(str)
}
