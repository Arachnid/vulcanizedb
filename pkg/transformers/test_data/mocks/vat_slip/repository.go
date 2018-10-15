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

package vat_slip

import (
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/vat_slip"
)

type MockVatSlipRepository struct {
	createError                     error
	markHeaderCheckedError          error
	markHeaderCheckedPassedHeaderID int64
	missingHeadersError             error
	missingHeaders                  []core.Header
	PassedStartingBlockNumber       int64
	PassedEndingBlockNumber         int64
	PassedHeaderID                  int64
	PassedModels                    []vat_slip.VatSlipModel
}

func (repository *MockVatSlipRepository) Create(headerID int64, models []vat_slip.VatSlipModel) error {
	repository.PassedHeaderID = headerID
	repository.PassedModels = models
	return repository.createError
}

func (repository *MockVatSlipRepository) MarkHeaderChecked(headerID int64) error {
	repository.markHeaderCheckedPassedHeaderID = headerID
	return repository.markHeaderCheckedError
}

func (repository *MockVatSlipRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	repository.PassedStartingBlockNumber = startingBlockNumber
	repository.PassedEndingBlockNumber = endingBlockNumber
	return repository.missingHeaders, repository.missingHeadersError
}

func (repository *MockVatSlipRepository) SetCreateError(e error) {
	repository.createError = e
}

func (repository *MockVatSlipRepository) SetMarkHeaderCheckedErr(e error) {
	repository.markHeaderCheckedError = e
}

func (repository *MockVatSlipRepository) SetMissingHeadersErr(e error) {
	repository.missingHeadersError = e
}

func (repository *MockVatSlipRepository) SetMissingHeaders(headers []core.Header) {
	repository.missingHeaders = headers
}

func (repository *MockVatSlipRepository) AssertMarkHeaderCheckedCalledWith(i int64) {
	Expect(repository.markHeaderCheckedPassedHeaderID).To(Equal(i))
}