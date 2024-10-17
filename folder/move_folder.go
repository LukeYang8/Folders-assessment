package folder

import (
	"errors"
	"strings"
)

// MoveFolder moves a folder to a new parent folder.
// @Params:
// name: The name is the name of the folder to be moved.
// dst: The dst is the name of the new parent folder.
// @Return:
// []Folder: The function returns a list of folders after the move operation.
// error: The function returns an error if:
//   - the folder does not exist
//   - the destination folder does not exist,
//   - the folder is moved to a child of itself,
//   - the folder is moved to a different organization.
//     else return nil.
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// find src folder
	var srcFolder = ""
	var srcOrgID = ""
	for _, folder := range f.folders {
		if folder.Name == name {
			srcFolder = folder.Paths
			srcOrgID = folder.OrgId.String()
			break
		}
	}
	// find dst folder
	var dstFolder = ""
	var dstOrgID = ""
	for _, folder := range f.folders {
		if folder.Name == dst {
			dstFolder = folder.Paths
			dstOrgID = folder.OrgId.String()
			break
		}
	}

	// check if src and dst folder exist
	if srcFolder == "" {
		return nil, errors.New("Source folder does not exist")
	}
	if dstFolder == "" {
		return nil, errors.New("Destination folder does not exist")
	}

	// move to itself
	if name == dst {
		return nil, errors.New("Cannot move folder to itself")
	}
	// move to child
	temp := strings.TrimSuffix(dstFolder, dst)
	if strings.HasPrefix(temp, srcFolder+".") || strings.Contains(temp, "."+srcFolder+".") {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}
	// move to different org
	if srcOrgID != dstOrgID {
		return nil, errors.New("Cannot move folder to a different organization")
	}

	// change folder paths
	for i, folder := range f.folders {

		if folder.Paths == name {
			f.folders[i].Paths = dstFolder + "." + name
		} else if strings.HasPrefix(folder.Paths, srcFolder+".") {
			newPath := dstFolder + "." + name + "." + strings.TrimPrefix(folder.Paths, srcFolder+".")
			f.folders[i].Paths = newPath
		} else if strings.HasPrefix(folder.Paths, srcFolder) {
			newPath := dstFolder + "." + name + "." + strings.TrimPrefix(folder.Paths, srcFolder)
			newPath = strings.TrimSuffix(newPath, ".")
			f.folders[i].Paths = newPath
		}
	}

	return f.folders, nil
}
