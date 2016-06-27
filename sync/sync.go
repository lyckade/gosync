//Package sync provides a sync logic between two folder trees
package sync

import "errors"

//ErrFilesHaveSameAge is returned by functions which are comparing the last change
//time. If both have the same time that error have to be returned.
var ErrFilesHaveSameAge = errors.New("Both files have the same last change time.")

//The Syncer interface abstracts the sync process for a simple
//testing of the sync logik.
//A file is represented by a string. The string should be a the path
//to the file.
type Syncer interface {
	//Copy is used to copy the source to the destination
	//Error is thrown, when it is not possible to copy
	//the file.
	Copy(string, string) error
	//GetOlderFile returns the filpath ot the older file as string
	GetOlderFile(string, string) (string, error)
	//GetNewerFile returns the filepath of the newer file as string
	GetNewerFile(string, string) (string, error)
	//Returns true if a file should be not synced
	SkipFile(string) bool
}
