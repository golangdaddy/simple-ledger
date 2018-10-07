package models

import (
	"fmt"
	"sync"
)

type GetInfo struct {
	BlockHeight int `json:"blockHeight"`
	BlockHash string `json:"blockHash"`
	TotalTransactions int `json:"totalTransactions"`
	TotalNativeCurrency int `json:"totalNativeCurrency"`
	txIndex map[string]*TX
	addressPermissions map[string]map[string]bool
	sync.RWMutex
}

func (self *GetInfo) Permissions(address string) map[string]bool {
	self.RLock()
	defer self.RUnlock()
	return self.addressPermissions[address]
}

func (self *GetInfo) GrantPermissions(address string, actions []string) {
	self.Lock()
	defer self.Unlock()
	if self.addressPermissions == nil {
		self.addressPermissions = map[string]map[string]bool{}
	}
	if self.addressPermissions[address] == nil {
		self.addressPermissions[address] = map[string]bool{}
	}
	for _, action := range actions {
		self.addressPermissions[address][action] = true
	}
}

func (self *GetInfo) RevokePermissions(address string, actions []string) {
	self.Lock()
	defer self.Unlock()
	if self.addressPermissions == nil {
		self.addressPermissions = map[string]map[string]bool{}
	}
	if self.addressPermissions[address] == nil {
		self.addressPermissions[address] = map[string]bool{}
	}
	for _, action := range actions {
		delete(self.addressPermissions[address], action)
	}
}

func (self *GetInfo) Index(k string, tx *TX) {
	self.Lock()
	defer self.Unlock()
	if self.txIndex == nil {
		self.txIndex = map[string]*TX{}
	}
	self.txIndex[k] = tx
}

func (self *GetInfo) Indexed(k string) (*TX, error) {
	self.RLock()
	defer self.RUnlock()
	if self.txIndex == nil {
		return nil, fmt.Errorf("TX NOT FOUND FOR KEY '%s'", k)
	}
	tx, ok := self.txIndex[k]
	if !ok {
		return nil, fmt.Errorf("NO TX FOUND FOR KEY '%s'", k)
	}
	return tx, nil
}
