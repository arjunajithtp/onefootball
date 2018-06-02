package services

import (
	"github.com/arjunajithtp/onefootball/internal/utils"
	"github.com/arjunajithtp/onefootball/public/dtos/api/ofbapi"
	"testing"
)

type MockDownloader struct {
	DownloadResponse *ofbapi.OneFootBallTeamDetails
	DownloadError    error
}

func TestDownloader_GetTeamDetails(t *testing.T) {
	testPlayer := ofbapi.Player{
		FirstName: "Test",
		LastName:  "Name",
		Age:       27,
	}

	mock := MockDownloader{
		DownloadResponse: &ofbapi.OneFootBallTeamDetails{
			Code: 0,
			Data: ofbapi.Data{
				Team: ofbapi.Team{
					Name: "Test Name",
					Players: []ofbapi.Player{
						testPlayer,
					},
				},
			},
		},
		DownloadError: nil,
	}
	d := Downloader{ofbDownloader: &mock}
	_, err := d.GetTeamDetails()

	utils.ShouldBeEqual(t, nil, err)
}
