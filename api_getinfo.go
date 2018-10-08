package main

import (
	"fmt"
	"sync"
	//
	"github.com/golangdaddy/tarantula/web"
	//
	"github.com/golangdaddy/simple-ledger/models"
)
type ChainInfo struct {
	BlockHeight int `json:"blockHeight"`
	BlockHash string `json:"blockHash"`
	TotalTransactions int `json:"totalTransactions"`
	TotalNativeCurrency int `json:"totalNativeCurrency"`
	txIndex map[string]*models.TX
	addressPermissions map[string]map[string]bool
	sync.RWMutex
}

func (self *ChainInfo) Permissions(address string) map[string]bool {
	self.RLock()
	defer self.RUnlock()
	return self.addressPermissions[address]
}

func (self *ChainInfo) GrantPermissions(address string, actions []string) {
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

func (self *ChainInfo) RevokePermissions(address string, actions []string) {
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

func (self *ChainInfo) Index(k string, tx *models.TX) {
	self.Lock()
	defer self.Unlock()
	if self.txIndex == nil {
		self.txIndex = map[string]*models.TX{}
	}
	self.txIndex[k] = tx
}

func (self *ChainInfo) Indexed(k string) (*models.TX, error) {
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

func (app *App) apiGetInfo(req web.RequestInterface) *web.ResponseStatus {

	app.info.RLock()
	defer app.info.RUnlock()

	info := *app.info
	return req.Respond(&info)
}
