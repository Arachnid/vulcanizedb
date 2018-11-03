// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package event_triggered

import (
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"
)

type ERC20EventDatastore interface {
	CreateTransfer(model *TransferModel, vulcanizeLogId int64) error
	CreateApproval(model *ApprovalModel, vulcanizeLogId int64) error
}

type ERC20EventRepository struct {
	*postgres.DB
}

func (repository ERC20EventRepository) CreateTransfer(transferModel *TransferModel, vulcanizeLogId int64) error {
	_, err := repository.DB.Exec(

		`INSERT INTO token_transfers (vulcanize_log_id, token_name, token_address, to_address, from_address, tokens, block, tx)
               VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		vulcanizeLogId, transferModel.TokenName, transferModel.TokenAddress, transferModel.To, transferModel.From, transferModel.Tokens, transferModel.Block, transferModel.TxHash)

	return err
}

func (repository ERC20EventRepository) CreateApproval(approvalModel *ApprovalModel, vulcanizeLogId int64) error {
	_, err := repository.DB.Exec(

		`INSERT INTO token_approvals (vulcanize_log_id, token_name, token_address, owner, spender, tokens, block, tx)
               VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
                ON CONFLICT (vulcanize_log_id) DO NOTHING`,
		vulcanizeLogId, approvalModel.TokenName, approvalModel.TokenAddress, approvalModel.Owner, approvalModel.Spender, approvalModel.Tokens, approvalModel.Block, approvalModel.TxHash)

	return err
}