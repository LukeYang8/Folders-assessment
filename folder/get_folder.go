package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

// GetAllChildFolders returns all child folders of a folder with the given name.
// Paths can contain non-alphanumeric characters, but . must be used as a separator.
// If folder does not exist return error
// @Params:
// name: The name is the name of the parent folder.
// orgID: The orgID is the organization ID of the parent folder.
// @Return:
// []Folder: The function returns a list of folders that are children of the parent folder.
// error: The function returns an error if the parent folder does not exist.
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...

	folders := f.folders
	res := []Folder{}

	folderExists := false
	folderInOrg := false

	for _, f := range folders {
		if strings.Contains(f.Paths, name) {
			folderExists = true
			if f.OrgId == orgID {
				folderInOrg = true
			}
		}
		if (strings.HasPrefix(f.Paths, name+".") || strings.Contains(f.Paths, "."+name+".")) && f.OrgId == orgID {
			res = append(res, f)
		}
	}

	if !folderExists {
		return nil, errors.New("Folder does not exist")
	} else if !folderInOrg {
		return nil, errors.New("Folder does not exist in the specified organization")
	}

	return res, nil
}
