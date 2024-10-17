package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	t.Parallel()
	tests := [...]struct {
		name    string
		src     string
		dst     string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		wantErr bool
		errMsg  string
	}{
		// TODO: your tests here

		{
			name:  "Test successful move folder",
			src:   "alpha",
			dst:   "bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "bravo.alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			wantErr: false,
		},
		{
			name:  "Test successful move folder where src has children",
			src:   "alpha",
			dst:   "bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "bravo.alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "bravo.alpha.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
		},
		{
			name:  "Test successful move folder with source folder having parents",
			src:   "bravo",
			dst:   "charlie",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.charlie",
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
					Paths: "alpha.charlie.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "alpha.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
		},
		{
			name:  "Test successful move folder where src has prefix path",
			src:   "bravo",
			dst:   "echo",
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
				{
					Name:  "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "echo",
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
					Paths: "echo.bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "charlie",
					Paths: "echo.bravo.charlie",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "echo",
					Paths: "echo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
		},
		{
			name:  "Test move folder to itself",
			src:   "alpha",
			dst:   "alpha",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:    []folder.Folder{},
			wantErr: true,
			errMsg:  "Cannot move folder to itself",
		},
		{
			name:  "Test move folder to child of itself",
			src:   "alpha",
			dst:   "bravo",
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
			want:    []folder.Folder{},
			wantErr: true,
			errMsg:  "Cannot move a folder to a child of itself",
		},
		{
			name:  "Test move folder to other orgainzation",
			src:   "alpha",
			dst:   "bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil("none"),
				},
			},
			want:    []folder.Folder{},
			wantErr: true,
			errMsg:  "Cannot move folder to a different organization",
		},
		{
			name:  "Test source folder does not exist",
			src:   "alpha",
			dst:   "bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{

				{
					Name:  "bravo",
					Paths: "bravo",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:    []folder.Folder{},
			wantErr: true,
			errMsg:  "Source folder does not exist",
		},
		{
			name:  "Test destination folder does not exist",
			src:   "alpha",
			dst:   "bravo",
			orgID: uuid.FromStringOrNil(folder.DefaultOrgID),
			folders: []folder.Folder{
				{
					Name:  "alpha",
					Paths: "alpha",
					OrgId: uuid.FromStringOrNil(folder.DefaultOrgID),
				},
			},
			want:    []folder.Folder{},
			wantErr: true,
			errMsg:  "Destination folder does not exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.src, tt.dst)
			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("MoveFolder() error = %v, wantErr %v", err, tt.errMsg)
				}
			} else if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("GetFoldersByOrgID() = %v, want %v", get, tt.want)
			}

		})
	}
}
