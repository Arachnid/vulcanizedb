// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vat_move

import (
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type MockVatMoveRepository struct {
	createError               error
	missingHeaders            []core.Header
	missingHeadersError       error
	PassedStartingBlockNumber int64
	PassedEndingBlockNumber   int64
	PassedHeaderID            int64
	PassedModels              []interface{}
	CheckedHeaderIDs          []int64
	CheckedHeaderError        error
	SetDbCalled               bool
}

func (repository *MockVatMoveRepository) Create(headerID int64, models []interface{}) error {
	repository.PassedHeaderID = headerID
	repository.PassedModels = models
	return repository.createError
}

func (repository *MockVatMoveRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	repository.PassedStartingBlockNumber = startingBlockNumber
	repository.PassedEndingBlockNumber = endingBlockNumber
	return repository.missingHeaders, repository.missingHeadersError
}

func (repository *MockVatMoveRepository) SetMissingHeadersError(e error) {
	repository.missingHeadersError = e
}

func (repository *MockVatMoveRepository) SetMissingHeaders(headers []core.Header) {
	repository.missingHeaders = headers
}

func (repository *MockVatMoveRepository) SetCreateError(e error) {
	repository.createError = e
}

func (repository *MockVatMoveRepository) MarkHeaderChecked(headerId int64) error {
	repository.CheckedHeaderIDs = append(repository.CheckedHeaderIDs, headerId)
	return repository.CheckedHeaderError
}

func (repository *MockVatMoveRepository) SetCheckedHeaderError(e error) {
	repository.CheckedHeaderError = e
}

func (repository *MockVatMoveRepository) SetDB(db *postgres.DB) {
	repository.SetDbCalled = true
}
