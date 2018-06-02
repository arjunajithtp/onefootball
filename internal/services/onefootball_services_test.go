package services

import (
	"github.com/arjunajithtp/onefootball/internal/utils"
	"github.com/arjunajithtp/onefootball/public/dtos/api/ofbapi"
	"github.com/pborman/uuid"
	"testing"
)

type MockDownloader struct {
	DownloadResponse *ofbapi.OneFootBallTeamDetails
	DownloadError    error
}

func (d *MockDownloader) GetTeamDetails(_ int) (*ofbapi.OneFootBallTeamDetails, error) {
	return d.DownloadResponse, d.DownloadError
}

func TestDownloader_GetAllPlayers(t *testing.T) {
	testPlayer := ofbapi.Player{
		FirstName: uuid.New(),
		LastName:  uuid.New(),
		Age:       uuid.New(),
	}

	mock := MockDownloader{
		DownloadResponse: &ofbapi.OneFootBallTeamDetails{
			Data: ofbapi.Data{
				Team: ofbapi.Team{
					Name: uuid.New(),
					Players: []ofbapi.Player{
						testPlayer,
					},
				},
			},
		},
		DownloadError: nil,
	}
	d := Downloader{ofbDownloader: &mock}
	mockRequiredList := map[string]bool{mock.DownloadResponse.Data.Team.Name: true}
	testAllPlayers, testPlayerNames, err := d.GetAllPlayers(mockRequiredList)
	utils.ShouldBeEqual(t, nil, err)
	utils.ShouldBeEqual(t, 1, len(testAllPlayers))
	utils.ShouldBeEqual(t, 1, len(testPlayerNames))
}
