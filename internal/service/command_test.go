package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	gmocks "github.com/ykkssyaa/Bash_Service/internal/gateway/mock"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	mockService "github.com/ykkssyaa/Bash_Service/internal/service/mock"
	lg "github.com/ykkssyaa/Bash_Service/pkg/logger"
	se "github.com/ykkssyaa/Bash_Service/pkg/serverError"
	"net/http"
	"testing"
)

func TestGetCommand(t *testing.T) {

	type funcResult struct {
		cmd  models.Command
		err  error
		skip bool // skip - flag to skip EXPECT mock result
	}

	defaultCmd := models.Command{
		Id:     1,
		Script: "echo 1",
		Status: models.StatusSuccess,
		Output: "1\n",
	}

	testTable := []struct {
		name           string
		inputId        int
		expectedResult funcResult
		repoResult     funcResult
		cacheResult    funcResult
	}{
		{
			name:           "Positive_HasNotInCache_HasInRepo",
			inputId:        1,
			expectedResult: funcResult{cmd: defaultCmd, err: nil},
			repoResult:     funcResult{cmd: defaultCmd, err: nil},
			cacheResult:    funcResult{cmd: models.Command{}, err: nil},
		},
		{
			name:           "Positive_HasInCache_HasNotInRepo",
			inputId:        1,
			expectedResult: funcResult{cmd: defaultCmd, err: nil},
			repoResult:     funcResult{cmd: models.Command{}, err: nil, skip: true},
			cacheResult:    funcResult{cmd: defaultCmd, err: nil},
		},
		{
			name:           "Positive_HasInCache_HasInRepo",
			inputId:        1,
			expectedResult: funcResult{cmd: defaultCmd, err: nil},
			repoResult:     funcResult{cmd: defaultCmd, err: nil, skip: true},
			cacheResult:    funcResult{cmd: defaultCmd, err: nil},
		},
		{
			name:           "Positive_HasInCache_HasInRepoOldVersion",
			inputId:        1,
			expectedResult: funcResult{cmd: defaultCmd, err: nil},
			repoResult: funcResult{cmd: models.Command{
				Id:     1,
				Status: models.StatusStarted,
				Output: "",
				Script: defaultCmd.Script,
			}, err: nil, skip: true},
			cacheResult: funcResult{cmd: defaultCmd, err: nil},
		},
		{
			name:           "NotFound",
			inputId:        1,
			expectedResult: funcResult{cmd: models.Command{}, err: se.ServerError{Message: "", StatusCode: http.StatusNotFound}},
			repoResult:     funcResult{cmd: models.Command{}, err: sql.ErrNoRows},
			cacheResult:    funcResult{cmd: models.Command{}, err: nil},
		},
		{
			name:           "Negative Id",
			inputId:        -1,
			expectedResult: funcResult{cmd: models.Command{}, err: se.ServerError{Message: consts.ErrorWrongId, StatusCode: http.StatusBadRequest}},
			repoResult:     funcResult{cmd: models.Command{}, err: nil, skip: true},
			cacheResult:    funcResult{cmd: models.Command{}, err: nil, skip: true},
		},
		{
			name:           "Zero Id",
			inputId:        0,
			expectedResult: funcResult{cmd: models.Command{}, err: se.ServerError{Message: consts.ErrorWrongId, StatusCode: http.StatusBadRequest}},
			repoResult:     funcResult{cmd: models.Command{}, err: nil, skip: true},
			cacheResult:    funcResult{cmd: models.Command{}, err: nil, skip: true},
		},
		{
			name:           "Some repo error",
			inputId:        1,
			expectedResult: funcResult{cmd: models.Command{}, err: se.ServerError{Message: consts.ErrorGetCommand, StatusCode: http.StatusInternalServerError}},
			repoResult:     funcResult{cmd: models.Command{}, err: errors.New("error")},
		},
	}

	logger := lg.InitLogger()

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// init mock controller
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			// init mocks
			repo := gmocks.NewMockCommand(ctl)
			cache := gmocks.NewMockCache(ctl)
			ctxStorage := gmocks.NewMockStorage(ctl)     // not in use
			executor := mockService.NewMockExecutor(ctl) // not in use

			if !test.repoResult.skip {
				// set expect return of repo
				repo.EXPECT().GetCommand(test.inputId).Return(test.repoResult.cmd, test.repoResult.err)
			}

			if !test.cacheResult.skip {
				// set expect return of cache
				cache.EXPECT().Get(test.inputId).Return(test.cacheResult.cmd, test.cacheResult.err)
			}

			// init service
			service := NewCommandService(repo, ctxStorage, cache, executor, logger)

			cmd, err := service.GetCommand(test.inputId)

			assert.Equal(t, cmd, test.expectedResult.cmd)
			assert.Equal(t, err, test.expectedResult.err)
		})
	}

}

func TestGetAllCommands(t *testing.T) {

	type funcResult struct {
		cmd  []models.Command
		err  error
		skip bool // skip - flag to skip EXPECT mock result
	}

	testData := []models.Command{
		{
			Id:     1,
			Script: "edkgkdkd",
			Output: "",
			Status: models.StatusError,
		},
		{
			Id:     2,
			Script: "echo 1",
			Output: "1\n",
			Status: models.StatusSuccess,
		},
		{
			Id:     3,
			Script: "for ((i=1; i<=4; i++))\ndo\n   echo $i\ndone",
			Output: "1\n2\n3\n4\n",
			Status: models.StatusSuccess,
		},
		{
			Id:     4,
			Script: "for ((i=1; i<=10; i++))\ndo\n   echo $i\ndone",
			Output: "",
			Status: models.StatusStarted,
		},
		{
			Id:     5,
			Script: "echo Hi",
			Output: "Hi\n",
			Status: models.StatusSuccess,
		},
		{
			Id:     6,
			Script: "for ((i=1; i<=10; i++))\ndo\n   echo $i\ndone",
			Output: "1\n",
			Status: models.StatusStopped,
		},
		{
			Id:     7,
			Script: "echo echo",
			Output: "echo\n",
			Status: models.StatusSuccess,
		},
	}

	testTable := []struct {
		name           string
		inputLimit     int
		inputOffset    int
		expectedResult funcResult
		repoResult     funcResult
	}{
		{
			name:           "Positive_BigLimit_ZeroOffset",
			inputLimit:     100,
			inputOffset:    0,
			expectedResult: funcResult{cmd: testData, err: nil},
			repoResult:     funcResult{cmd: testData, err: nil},
		},
		{
			name:           "Positive_SmallLimit_ZeroOffset",
			inputLimit:     3,
			inputOffset:    0,
			expectedResult: funcResult{cmd: testData[:3], err: nil},
			repoResult:     funcResult{cmd: testData[:3], err: nil},
		},
		{
			name:           "Positive_BigLimit_SmallOffset",
			inputLimit:     100,
			inputOffset:    1,
			expectedResult: funcResult{cmd: testData[1:], err: nil},
			repoResult:     funcResult{cmd: testData[1:], err: nil},
		},
		{
			name:           "Positive_SmallLimit_SmallOffset",
			inputLimit:     1,
			inputOffset:    1,
			expectedResult: funcResult{cmd: testData[1:2], err: nil},
			repoResult:     funcResult{cmd: testData[1:2], err: nil},
		},
		{
			name:           "Negative offset",
			inputLimit:     10,
			inputOffset:    -1,
			expectedResult: funcResult{cmd: nil, err: se.ServerError{Message: consts.ErrorWrongOffset, StatusCode: http.StatusBadRequest}},
			repoResult:     funcResult{cmd: nil, err: nil, skip: true},
		},
		{
			name:           "Negative limit",
			inputLimit:     -1,
			inputOffset:    1,
			expectedResult: funcResult{cmd: nil, err: se.ServerError{Message: consts.ErrorWrongLimit, StatusCode: http.StatusBadRequest}},
			repoResult:     funcResult{cmd: nil, err: nil, skip: true},
		},
		{
			name:           "Has not data",
			inputLimit:     1,
			inputOffset:    100,
			expectedResult: funcResult{cmd: nil, err: nil},
			repoResult:     funcResult{cmd: nil, err: nil},
		},
		{
			name:           "Some repo error",
			inputLimit:     1,
			inputOffset:    100,
			expectedResult: funcResult{cmd: nil, err: se.ServerError{Message: consts.ErrorGetCommand, StatusCode: http.StatusInternalServerError}},
			repoResult:     funcResult{cmd: nil, err: errors.New("error")},
		},
	}

	logger := lg.InitLogger()

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// init mock controller
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			// init mocks
			repo := gmocks.NewMockCommand(ctl)
			cache := gmocks.NewMockCache(ctl)            // not in use
			ctxStorage := gmocks.NewMockStorage(ctl)     // not in use
			executor := mockService.NewMockExecutor(ctl) // not in use

			if !test.repoResult.skip {
				// set expect return of repo
				repo.EXPECT().GetAllCommands(test.inputLimit, test.inputOffset).Return(test.repoResult.cmd, test.repoResult.err)
			}

			// init service
			service := NewCommandService(repo, ctxStorage, cache, executor, logger)

			cmds, err := service.GetAllCommands(test.inputLimit, test.inputOffset)

			assert.Equal(t, cmds, test.expectedResult.cmd)
			assert.Equal(t, err, test.expectedResult.err)
		})
	}

}

func TestStopCommand(t *testing.T) {

	testTable := []struct {
		name           string
		inputId        int
		expectedResult error
		storageResult  bool // storageResult = true if has result
	}{
		{
			name:           "Positive",
			inputId:        1,
			expectedResult: nil,
			storageResult:  true,
		},
		{
			name:           "NotFound",
			inputId:        1,
			expectedResult: se.ServerError{Message: "", StatusCode: http.StatusNotFound},
			storageResult:  false,
		},
	}

	logger := lg.InitLogger()

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// init mock controller
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			// init mocks
			repo := gmocks.NewMockCommand(ctl)
			cache := gmocks.NewMockCache(ctl)
			ctxStorage := gmocks.NewMockStorage(ctl)     // not in use
			executor := mockService.NewMockExecutor(ctl) // not in use

			// set expect return of ctxStorage
			if test.storageResult {
				_, cancel := context.WithCancel(context.Background())

				ctxStorage.EXPECT().Get(test.inputId).Return(cancel)
			} else {
				ctxStorage.EXPECT().Get(test.inputId).Return(nil)
			}

			// init service
			service := NewCommandService(repo, ctxStorage, cache, executor, logger)

			err := service.StopCommand(test.inputId)

			assert.Equal(t, err, test.expectedResult)

		})
	}

}

func TestCreateCommand(t *testing.T) {

	type funcResult struct {
		cmd  models.Command
		err  error
		skip bool // skip - flag to skip EXPECT mock result
	}

	defaultCmd := models.Command{
		Id:     1,
		Script: "echo 1",
		Status: models.StatusStarted,
		Output: "",
	}

	testTable := []struct {
		name           string
		inputScript    string
		expectedResult funcResult
		repoResult     funcResult
		ExecCmdResult  error
	}{
		{
			name:           "Positive",
			inputScript:    "echo 1",
			expectedResult: funcResult{cmd: defaultCmd, err: nil},
			repoResult:     funcResult{cmd: defaultCmd, err: nil},
			ExecCmdResult:  nil,
		},
	}

	logger := lg.InitLogger()

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			// init mock controller
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			// init mocks
			repo := gmocks.NewMockCommand(ctl)
			cache := gmocks.NewMockCache(ctl)
			ctxStorage := gmocks.NewMockStorage(ctl)
			executor := mockService.NewMockExecutor(ctl)

			// init service
			service := NewCommandService(repo, ctxStorage, cache, executor, logger)

			cmd, err := service.CreateCommand(test.inputScript)

			assert.Equal(t, cmd, test.expectedResult.cmd)
			assert.Equal(t, err, test.expectedResult.err)

		})
	}

}
