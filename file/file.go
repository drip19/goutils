package file

import "os"

//CreateDirIfNotExist create given folder
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		return err
	}

	return nil
}
