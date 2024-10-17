package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{
			name:    "Test empty folder list",
			orgID:   uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
		{
			name:  "Test get from orgID",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
		},
		{
			name:  "Test get none from orgID",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil("none"),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil("none"),
				},
			},
			want: []folder.Folder{},
		},
		{
			name:  "Test mixed orgID",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo1",
					Paths: "alpha.bravo1",
					OrgId: uuid.FromStringOrNil("none"),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", get, tt.want)
			}

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name      string
		orgID     uuid.UUID
		parent    string
		folders   []folder.Folder
		want      []folder.Folder
		wantError bool
		errMsg    string
	}{
		// TODO: your tests here
		{
			name:      "Test empty folder list",
			orgID:     uuid.FromStringOrNil(folder.DefaultOrgID),
			parent:    "alpha",
			folders:   []folder.Folder{},
			want:      []folder.Folder{},
			wantError: true,
			errMsg:    "Folder does not exist",
		},
		{
			name:   "Test folder has no child folders",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "charlie",
			folders: []folder.Folder{
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:      []folder.Folder{},
			wantError: false,
		},
		{
			name:   "Test path is same as parent/has no child folders",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "alpha",
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:      []folder.Folder{},
			wantError: false,
		},
		{
			name:   "Test multiple successful get child folders",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "alpha",
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			wantError: false,
		},
		{
			name:   "Test wrong orgID",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "alpha",
			folders: []folder.Folder{
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil("none"),
				},
			},
			want:      []folder.Folder{},
			wantError: true,
			errMsg:    "Folder does not exist in the specified organization",
		},
		{
			name:   "Test folder does not exist",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "alpha",
			folders: []folder.Folder{
				{
					Name:  "bravo",
					Paths: "bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:      []folder.Folder{},
			wantError: true,
			errMsg:    "Folder does not exist",
		},
		{
			name:   "Test folder in middle of path",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "bravo",
			folders: []folder.Folder{
				{
					Name:  "bravo",
					Paths: "alpha.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			wantError: false,
		},
		{
			name:   "Test name string in path but not a folder",
			orgID:  uuid.FromStringOrNil(folder.DefaultOrgID),
			parent: "alpha",
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "alphabravo",
					Paths: "alphabravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "beta",
					Paths: "alphabravo.beta",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:      []folder.Folder{},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.parent)
			if tt.wantError {
				if err.Error() != tt.errMsg {
					t.Errorf("GetAllChildFolders() error = %v, wantError = %v", err, tt.errMsg)
					return
				}
			} else if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want = %v", get, tt.want)
			}

		})
	}
}
