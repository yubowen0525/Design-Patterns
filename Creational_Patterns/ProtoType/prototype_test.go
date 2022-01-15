package prototype

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := KeyWords{
		"testA": &KeyWord{
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &KeyWord{
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &KeyWord{
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*KeyWord{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)

	for k, v := range got {
		fmt.Println(v)
		fmt.Printf("word:%v+, visit:%v+, updateAt:%v+\n", &words[k].Word, &words[k].Visit, &words[k].UpdatedAt)
		fmt.Printf("word:%v+, visit:%v+, updateAt:%v+\n", &got[k].Word, &got[k].Visit, &got[k].UpdatedAt)
	}

	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
