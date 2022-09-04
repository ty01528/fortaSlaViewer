package handler

import (
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)

func GetScoreFromLeveldb() []User {
	addresses := readAddressFromLeveldb()
	var scoreList []User
	var wg sync.WaitGroup
	for _, address := range addresses {
		address := address
		wg.Add(1)
		go func() {
			temp := address
			var err error = nil
			temp.RealtimeScore, err = GetRealTimeScore(temp.Address)
			if err != nil {
				temp.RealtimeScore = "Unregister/NetworkError"
			}
			temp.WeekScore, err = GetWeekScore(temp.Address)
			if err != nil {
				temp.WeekScore = "Unregister/NetworkError"
			}
			scoreList = append(scoreList, temp)
			wg.Done()
		}()
	}
	wg.Wait()
	return scoreList
}

func readAddressFromLeveldb() []User {
	var machineList []User
	db, _ := leveldb.OpenFile("./address.db", nil)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		var temp User
		temp.Address = string(iter.Key())
		temp.Name = string(iter.Value())
		if !(temp.Address == "") {
			machineList = append(machineList, temp)
		}
	}
	iter.Release()
	db.Close()
	return machineList
}
