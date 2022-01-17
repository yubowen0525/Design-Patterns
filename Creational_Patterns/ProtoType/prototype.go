package prototype

import (
	"encoding/json"
	"time"
)

type KeyWord struct {
	Word      string     `json:word`
	Visit     int        `json:visit`
	UpdatedAt *time.Time `json:updateAt`
}

func (k *KeyWord) Clone() *KeyWord {
	var newKeyWord KeyWord
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyWord)
	return &newKeyWord
}

type KeyWords map[string]*KeyWord

func (words KeyWords) Clone(updateKeyWords []*KeyWord) KeyWords {
	newKeyWords := KeyWords{}
	for key, value := range words {
		// 浅拷贝
		newKeyWords[key] = value
	}

	// 更新需要更新的字段，深拷贝
	for _, word := range updateKeyWords {
		newKeyWords[word.Word] = word.Clone()
	}

	return newKeyWords
}
