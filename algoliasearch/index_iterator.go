package algoliasearch

import (
	"errors"
)

// IndexIterator is used to iterate over indices when using `Browse`-like
// functions.
type IndexIterator struct {
	answer interface{}
	index  *Index
	params interface{}
	pos    int
}

// NewIndexIterator instantiates an new IndexIterator on the `index` given the
// `params` parameters. It also initializes it to the first page of results.
func NewIndexIterator(index *Index, params interface{}, cursor string) (*IndexIterator, error) {
	it := &IndexIterator{
		answer: map[string]interface{}{"cursor": cursor},
		index:  index,
		params: params,
		pos:    0,
	}

	return it, it.loadNextPage()
}

// Next iterates to the next result and move the cursor.
func (it *IndexIterator) Next() (interface{}, error) {
	var err error

	for err == nil {
		hits := it.answer.(map[string]interface{})["hits"].([]interface{})

		if it.pos < len(hits) {
			it.pos++
			return hits[it.pos-1], nil
		}

		if cursor, ok := it.GetCursor(); ok && len(cursor) > 0 {
			err = it.loadNextPage()
			continue
		}

		return nil, errors.New("End of the index reached")
	}

	return nil, err
}

// GetCursor returns the current underlying cursor. The returned boolean is set
// to `false` if the end of the index has been reached.
func (it *IndexIterator) GetCursor() (string, bool) {
	if cursor, ok := it.answer.(map[string]interface{})["cursor"]; cursor != nil {
		return cursor.(string), ok
	} else {
		return "", ok
	}
}

// loadNextPage loads the next page of results and resets the cursor position.
func (it *IndexIterator) loadNextPage() error {
	it.pos = 0
	cursor, _ := it.GetCursor()
	answer, err := it.index.BrowseFrom(it.params, cursor)
	it.answer = answer
	return err
}
