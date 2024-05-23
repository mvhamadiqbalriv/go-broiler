package service

import (
	"errors"
	"sync"
	"time"
)

// TokenBlacklistServiceImpl handles operations related to blacklisted tokens
type TokenBlacklistServiceImpl struct {
    blacklistedTokens map[string]time.Time
    mutex             sync.RWMutex
}

// NewTokenBlacklistServiceImpl creates a new instance of TokenBlacklistServiceImpl
func NewTokenBlacklistService() *TokenBlacklistServiceImpl {
    return &TokenBlacklistServiceImpl{
        blacklistedTokens: make(map[string]time.Time), // Initialize the map here
    }
}

// AddTokenToBlacklist adds a token to the blacklist
func (s *TokenBlacklistServiceImpl) AddTokenToBlacklist(token string) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    if _, exists := s.blacklistedTokens[token]; exists {
        return errors.New("token already blacklisted")
    }

    // Add token to blacklist with expiration time (e.g., 24 hours)
    s.blacklistedTokens[token] = time.Now().Add(24 * time.Hour)
    return nil
}

// IsTokenBlacklisted checks if a token is blacklisted
func (s *TokenBlacklistServiceImpl) IsTokenBlacklisted(token string) bool {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    // Check if the token is blacklisted
    expirationTime, exists := s.blacklistedTokens[token]
    if !exists || time.Now().After(expirationTime) {
        // Token does not exist in the blacklist or has expired
        return false
    }
    return true
}
