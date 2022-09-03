package handler

import (
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)

func GetScoreFromAddresses() []User {
	addresses := readAddressFromLeveldb()
	var scoreList []User
	var wg sync.WaitGroup
	for _, address := range addresses {
		address := address
		wg.Add(1)
		temp := address
		var err error = nil
		temp.realtimeScore, err = GetRealTimeScore(temp.address)
		if err != nil {
			temp.realtimeScore = address.realtimeScore
		}
		temp.weekScore, err = GetWeekScore(temp.address)
		if err != nil {
			temp.weekScore = address.weekScore
		}
		scoreList = append(scoreList, temp)
		wg.Done()
		wg.Wait()
	}
	return scoreList
}

func readAddressFromLeveldb() []User {
	var machineList []User
	db, _ := leveldb.OpenFile("./address.db", nil)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		var temp User
		temp.address = string(iter.Key())
		temp.name = string(iter.Value())
		machineList = append(machineList, temp)
	}
	iter.Release()
	return machineList
}
