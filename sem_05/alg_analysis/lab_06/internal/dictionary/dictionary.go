package dictionary

import (
	"lab_06_01/internal/database"
	"lab_06_01/internal/models"
)

type Word struct {
	Key   int
	Value models.Cat
}

type Dictionary []Word

func InitDictionary(db *database.DB) (Dictionary, error) {
	elems, err := db.GetDBData()
	if err != nil {
		return nil, err
	}

	dict := make(Dictionary, 0)
	for _, v := range elems {
		word := Word{
			Key:   v.Fluffiness,
			Value: v,
		}
		dict = append(dict, word)
	}

	return dict, nil
}

func (d *Dictionary) BinarySearch(key int, isStart bool) int {
	last := len(*d) - 1
	if key > (*d)[last].Key {
		return last
	}
	if key < (*d)[0].Key {
		return 0
	}

	lo := 0
	hi := last
	for lo <= hi {
		m := (lo + hi) / 2
		if key > (*d)[m].Key {
			lo = m + 1
		} else if key < (*d)[m].Key {
			hi = m - 1
		} else {
			return m
		}
	}

	crossed := (*d)[lo].Key > (*d)[hi].Key
	if crossed == isStart {
		return lo
	}

	return hi
}

func (d *Dictionary) Search(start, end int) []models.Cat {
	lim1 := d.BinarySearch(start, true)
	lim2 := d.BinarySearch(end, false)

	res := make([]models.Cat, 0)
	for i := lim1; i < lim2+1; i++ {
		res = append(res, (*d)[i].Value)
	}

	return res
}
