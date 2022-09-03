package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// User 这个结构体没啥用
type User struct {
	name           string
	address        string
	depositAddress string
	weekScore      string
	realtimeScore  string
}

func PhraseAndSaveAddress(context *gin.Context) {
	addresses := context.PostForm("addresses")
	addressList := phraseAddresses(addresses)
	err := saveAddress(addressList)
	if err != nil {
		log.Println("写入地址时出现错误！")
		context.HTML(http.StatusBadRequest, "index.html", gin.H{})
	} else {
		context.HTML(http.StatusOK, "index.html", gin.H{})
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
		user.name = "User" + strconv.Itoa(index)
		user.address = strings.Replace(address, "\n", "", -1)
		err := db.Put([]byte(user.address), []byte(user.name), nil)
		if err != nil {
			return err
		}
	}
	defer db.Close()
	defer println("address save to db complete!")
	return nil
}
