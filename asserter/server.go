// Copyright 2020 Coinbase, Inc.
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

package asserter

import (
	"errors"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// SupportedNetworks returns an error if there is an invalid
// types.NetworkIdentifier or there is a duplicate.
func SupportedNetworks(supportedNetworks []*types.NetworkIdentifier) error {
	if len(supportedNetworks) == 0 {
		return errors.New("no supported networks")
	}

	parsed := make([]*types.NetworkIdentifier, len(supportedNetworks))
	for i, network := range supportedNetworks {
		if err := NetworkIdentifier(network); err != nil {
			return err
		}

		if containsNetworkIdentifier(parsed, network) {
			return fmt.Errorf("supported network duplicate %+v", network)
		}
		parsed[i] = network
	}

	return nil
}

// SupportedNetwork returns a boolean indicating if the requestNetwork
// is allowed. This should be called after the requestNetwork is asserted.
func (a *Asserter) SupportedNetwork(
	requestNetwork *types.NetworkIdentifier,
) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if !containsNetworkIdentifier(a.supportedNetworks, requestNetwork) {
		return fmt.Errorf("%+v is not supported", requestNetwork)
	}

	return nil
}

// ValidSupportedNetwork returns an error if a types.NetworkIdentifier
// is not valid or not supported.
func (a *Asserter) ValidSupportedNetwork(
	requestNetwork *types.NetworkIdentifier,
) error {
	if err := NetworkIdentifier(requestNetwork); err != nil {
		return err
	}

	if err := a.SupportedNetwork(requestNetwork); err != nil {
		return err
	}

	return nil
}

// AccountBalanceRequest ensures that a types.AccountBalanceRequest
// is well-formatted.
func (a *Asserter) AccountBalanceRequest(request *types.AccountBalanceRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("AccountBalanceRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if err := AccountIdentifier(request.AccountIdentifier); err != nil {
		return err
	}

	if request.BlockIdentifier == nil {
		return nil
	}

	if request.BlockIdentifier != nil && !a.historicalBalanceLookup {
		return errors.New("historical balance lookup is not supported")
	}

	return PartialBlockIdentifier(request.BlockIdentifier)
}

// BlockRequest ensures that a types.BlockRequest
// is well-formatted.
func (a *Asserter) BlockRequest(request *types.BlockRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("BlockRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	return PartialBlockIdentifier(request.BlockIdentifier)
}

// BlockTransactionRequest ensures that a types.BlockTransactionRequest
// is well-formatted.
func (a *Asserter) BlockTransactionRequest(request *types.BlockTransactionRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("BlockTransactionRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if err := BlockIdentifier(request.BlockIdentifier); err != nil {
		return err
	}

	return TransactionIdentifier(request.TransactionIdentifier)
}

// ConstructionMetadataRequest ensures that a types.ConstructionMetadataRequest
// is well-formatted.
func (a *Asserter) ConstructionMetadataRequest(request *types.ConstructionMetadataRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionMetadataRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if request.Options == nil {
		return errors.New("ConstructionMetadataRequest.Options is nil")
	}

	return nil
}

// ConstructionSubmitRequest ensures that a types.ConstructionSubmitRequest
// is well-formatted.
func (a *Asserter) ConstructionSubmitRequest(request *types.ConstructionSubmitRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionSubmitRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if request.SignedTransaction == "" {
		return errors.New("ConstructionSubmitRequest.SignedTransaction is empty")
	}

	return nil
}

// MempoolTransactionRequest ensures that a types.MempoolTransactionRequest
// is well-formatted.
func (a *Asserter) MempoolTransactionRequest(request *types.MempoolTransactionRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("MempoolTransactionRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	return TransactionIdentifier(request.TransactionIdentifier)
}

// MetadataRequest ensures that a types.MetadataRequest
// is well-formatted.
func (a *Asserter) MetadataRequest(request *types.MetadataRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("MetadataRequest is nil")
	}

	return nil
}

// NetworkRequest ensures that a types.NetworkRequest
// is well-formatted.
func (a *Asserter) NetworkRequest(request *types.NetworkRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("NetworkRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	return nil
}

// ConstructionDeriveRequest ensures that a types.ConstructionDeriveRequest
// is well-formatted.
func (a *Asserter) ConstructionDeriveRequest(request *types.ConstructionDeriveRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionDeriveRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if err := PublicKey(request.PublicKey); err != nil {
		return err
	}

	return nil
}

// ConstructionPreprocessRequest ensures that a types.ConstructionPreprocessRequest
// is well-formatted.
func (a *Asserter) ConstructionPreprocessRequest(
	request *types.ConstructionPreprocessRequest,
) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionPreprocessRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if err := a.Operations(request.Operations, true); err != nil {
		return err
	}

	return nil
}

// ConstructionPayloadsRequest ensures that a types.ConstructionPayloadsRequest
// is well-formatted.
func (a *Asserter) ConstructionPayloadsRequest(request *types.ConstructionPayloadsRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionPayloadsRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if err := a.Operations(request.Operations, true); err != nil {
		return err
	}

	return nil
}

// ConstructionCombineRequest ensures that a types.ConstructionCombineRequest
// is well-formatted.
func (a *Asserter) ConstructionCombineRequest(request *types.ConstructionCombineRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionCombineRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if len(request.UnsignedTransaction) == 0 {
		return errors.New("UnsignedTransaction cannot be empty")
	}

	if err := Signatures(request.Signatures); err != nil {
		return err
	}

	return nil
}

// ConstructionHashRequest ensures that a types.ConstructionHashRequest
// is well-formatted.
func (a *Asserter) ConstructionHashRequest(request *types.ConstructionHashRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionHashRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if len(request.SignedTransaction) == 0 {
		return errors.New("SignedTransaction cannot be empty")
	}

	return nil
}

// ConstructionParseRequest ensures that a types.ConstructionParseRequest
// is well-formatted.
func (a *Asserter) ConstructionParseRequest(request *types.ConstructionParseRequest) error {
	if a == nil {
		return ErrAsserterNotInitialized
	}

	if request == nil {
		return errors.New("ConstructionParseRequest is nil")
	}

	if err := a.ValidSupportedNetwork(request.NetworkIdentifier); err != nil {
		return err
	}

	if len(request.Transaction) == 0 {
		return errors.New("Transaction cannot be empty")
	}

	return nil
}
