package dto

import ()

type WorkCombineIn struct {
	MediaAccountUuids, MediaAccountNames []string
}

type WorkCombineOut struct {
	MediaAccountUuidsBytes []byte
	MediaAccountNamesBytes []byte
	PlatformIDsBytes       []byte
	ConfigBytes            []byte
}

type CtxUserInfo struct {
	UserID string
	Name   string
	Phone  string
}
