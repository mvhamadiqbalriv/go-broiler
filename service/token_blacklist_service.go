package service

type TokenBlacklistService interface {
    AddTokenToBlacklist(token string) error
    IsTokenBlacklisted(token string) bool
}