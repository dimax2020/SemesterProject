package mystructs

import (
	"crypto"
)

type FileStruct struct {
	ID        string      `json:"file_id"`
	Hash      crypto.Hash `json:"hash"`
	Extension string      `json:"extension"`
	Name      string      `json:"name"`
	Content   []byte      `json:"content"`
}

type FolderStruct struct {
	Files     []FileStruct `json:"files"`
}
