package ArrsPlanningService

/*
   ©️ 2022 B1 Digital
   User    : ODI
   Name    : Özlem DEMİRCİ
   Date    : 15.02.2022  12:47
   Notes   :
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

func (s *ArrsPlanningService) PrintPrettyStruct(value interface{}) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", string(valueJSON))
}

func (s *ArrsPlanningService) SPrintPrettyStruct(value interface{}) (res string) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	res = fmt.Sprintf("%s\n", string(valueJSON))
	return
}

func (s *ArrsPlanningService) SavePrettyStruct(fileName string, value interface{}) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.WriteFile(fileName, []byte(valueJSON), 0644)
	if err != nil {
		return
	}
}
