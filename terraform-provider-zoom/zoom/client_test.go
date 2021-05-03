package zoom

import (
	"github.com/stretchr/testify/assert"
	//"os"
	"testing"
)
/*
func TestClient_GetCreatedUserData(t *testing.T) {

	testCases := []struct {
		testName     string
		userID       string
		expectErr    bool
		expectedResp *Users
	}{
		{
			testName:  "user exists",
			userID:    "ashishdhodria1999@gmail.com",
			expectErr: false,
			expectedResp: &Users{
				ID: "_9YyowePRtacFLL5SEnK-w",
				FirstName: "ashu",
				LastName:  "malav",
				Email:     "ashishdhodria1999@gmail.com",
				Type:      1,
				Status: "active",
				Pmi:2329939862, Time_Zone:"",
				Verified:0,
				CreatedAt:"2021-04-24T11:35:36Z",
				LastLoginTime:"2021-04-24T11:35:37Z",
				PicUrl:"https://lh3.googleusercontent.com/a-/AOh14Ggg5Q8jSrFUT6C4cLClqDeyiO2A-pZ-THaDHuMh=s96-c",
				Language:"",
				RoleId:"2",
				PhoneNumber:"",
			},
		},
		{
			testName:     "user does not exist",
			userID:       "ashishdhodria@gmail.com",
			expectErr:    true,
			expectedResp: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			authToken := os.Getenv("zoom_token")
			client := NewClient(authToken)

			user, err := client.GetUserData(tc.userID)

			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}
*/
/*
func TestClient_CreateUser(t *testing.T) {

	testCases := []struct {
		testName  string
		newItem   UserCreate
		expectedResp *UserCreateInfo
		status    *Status
		userID       string
		expectErr bool
	}{
		{
			testName: "user created successfully",
			newItem: UserCreate{
				Action: "create",
				UserCreateInfo: UserCreateInfo{
					FirstName: "ashu",
					LastName:  "malav",
					Email:     "ashishdhodria1999@gmail.com",
					Type:      1,
				},
			},
			userID: "ashishdhodria1999@gmail.com",
			expectedResp: &UserCreateInfo{
				FirstName: "ashu",
				LastName: "malav",
				Email: "ashishdhodria1999@gmail.com",
				Type: 1,
			},
			status:&Status{
				Status: "pending",
			},
			expectErr: false,
		},
		{
			testName: "user already exists",
			userID: "ashishdhodria1999@gmail.com",
			expectedResp: &UserCreateInfo{
				FirstName: "ashu",
				LastName: "malav",
				Email: "ashishdhodria1999@gmail.com",
				Type: 1,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			//authToken := os.Getenv("zoom_token")
			authToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk"
			client := NewClient(authToken)
			_, err := client.CreateUser(tc.newItem, tc.newItem.UserCreateInfo.Email)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetCreatedUserData(tc.userID)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedResp, user)

		})
	}
}
*/

func TestClient_UpdateUser(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedUser UserCreateInfo
		expectedResp *UserCreateInfo
		userID       string
		expectErr   bool
	}{
		{
			testName: "user exists",
			updatedUser: UserCreateInfo{
				FirstName: "ashu",
				LastName: "malav",
				Email: "ashishdhodria1999@gmail.com",
				Type: 1,
			},
			expectedResp:& UserCreateInfo{
				FirstName: "ashu",
				LastName: "malav",
				Email: "ashishdhodria1999@gmail.com",
				Type: 1,
			} ,
			userID: "ashishdhodria1999@gmail.com",
			expectErr: false,
		},
		{
			testName: "item does not exist",
			userID: "ashishdhodria@gmail.com",
			updatedUser: UserCreateInfo{
				FirstName: "ashish",
				LastName: "dhodria",
				Email: "ashishdhodria1999@gmail.com",
				Type: 1,
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			//authToken := os.Getenv("zoom_token")
			authToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk"

			client := NewClient(authToken)
			_,err := client.UpdateUser(tc.userID,tc.updatedUser)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			user, err := client.GetCreatedUserData(tc.userID)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, user)
		})
	}
}

/*
func TestClient_DeleteUser(t *testing.T) {
	testCases := []struct {
		testName  string
		userID  string

		expectErr bool
	}{
		{
			testName: "user exists",
			userID: "ashishdhodria1999@gmail.com",
			expectErr: false,
		},
		{
			testName:  "item does not exist",
			userID: "ashishdhodria1999@gmail.com",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			//authToken := os.Getenv("zoom_token")
			authToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjAxMDExOTUsImlhdCI6MTYxOTQ5NjQwN30.P85E91mA-T_-tISlKA7GX0XRcD6bheVArg-spWLwSTk"

			client := NewClient(authToken)
			err := client.DeleteUser(tc.userID)

			if tc.expectErr {
				assert.Error(t, err)
				return
			}
		//	_, err = client.GetCreatedUserData(tc.userID)
		//	assert.Error(t, err)
		})
	}
}
*/