//Package sync provides a sync logic between two folder trees
package sync

import (
	"errors"
	"path"
	"path/filepath"
)

//ErrFilesHaveSameAge is returned by functions which are comparing the last change
//time. If both have the same time that error have to be returned.
var ErrFilesHaveSameAge = errors.New("Both files have the same last change time.")

//The Filesyncer interface abstracts the sync process for syncing a
//file.
//A file is represented by a string. The string should be a the path
//to the file.
type Filesyncer interface {
	//Copy is used to copy the source to the destination
	//Error is thrown, when it is not possible to copy
	//the file.
	Copy(string, string) error
	//GetNewerFile returns the filepath of the newer file as string
	GetNewerFile(string, string) (string, error)
}

//Sync syncs a file. The newer file is copied.
//A implementation of the filesyncer interface is the basis for
//all the file functions.
func Sync(fs Filesyncer, filePath1, filePath2 string) error {
	var dst string
	src, err := fs.GetNewerFile(filePath1, filePath2)
	if err != nil {
		return err
	}
	if filePath2 == src {
		dst = filePath1
	} else {
		dst = filePath2
	}
	err = fs.Copy(src, dst)

	return err
}

//MakeDistPath takes the relative path to a root folder and adds
//that to a dist folder (the target of the sync).
func MakeDistPath(fpath, rootFolder, distFolder string) (string, error) {
	relPath, err := filepath.Rel(rootFolder, fpath)
	return path.Join(distFolder, relPath), err
}
