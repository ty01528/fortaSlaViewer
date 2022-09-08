package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"strconv"
	"strings"
)

// User 这个结构体没啥用
type User struct {
	Name           string
	Address        string
	DepositAddress string
	WeekScore      string
	RealtimeScore  string
}

func PhraseAndSaveAddress(context *gin.Context) {
	addresses := context.PostForm("addresses")
	addressList := phraseAddresses(addresses)
	err := saveAddress(addressList)
	if err != nil {
		log.Println("写入地址时出现错误！")
	}
}

func phraseAddresses(addresses string) []string {
	addresses = strings.Replace(addresses, " ", "", -1)
	addressesList := strings.Split(addresses, "\r")
	return addressesList
}

func saveAddress(addresses []string) error {
	db, err := leveldb.OpenFile("./address.db", nil)
	if err != nil {
		log.Println("error opening databases! ")
		log.Panicln(err)
		return err
	}
	user := &User{}
	for index, address := range addresses {
		user.Name = "User" + strconv.Itoa(index)
		user.Address = strings.Replace(address, "\n", "", -1)
		if len(user.Address) == len("0x95Bc81164d901130c87185E2Dfb55D634f629773") {
			err := db.Put([]byte(user.Address), []byte(user.Name), nil)
			if err != nil {
				return err
			}
		}
	}
	defer db.Close()
	defer println("address save to db complete!")
	return nil
}

func FromAddressSaveName(address string, name string) error {
	db, err := leveldb.OpenFile("./address.db", nil)
	defer db.Close()
	if err != nil {
		log.Println("error opening databases! ")
		log.Panicln(err)
		return err
	}
	err = db.Put([]byte(address), []byte(name), nil)
	if err != nil {
		return err
	}
	log.Println(address + " 修改用户名为：" + name + " 成功！")
	return nil
}
