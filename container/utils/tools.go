package utils

import "os"

func Exists(path string) bool {
	if path[len(path)-1:len(path)]!="/"{
		path=path+"/"
	}
	_, err := os.Stat(path+"sqlmapapi.py")
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}