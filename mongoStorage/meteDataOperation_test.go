package mongoStorage

import (
	"PillarsPhenomVFXWeb/utility"
	"testing"
	"time"
)

func Test_StoreMetaData(t *testing.T) {
	str := "testing"
	//metaDataCode := utility.GenerateCode(&str)
	materialCode := utility.GenerateCode(&str)
	metaDataName := utility.GenerateCode(&str)
	metaDataValue := utility.GenerateCode(&str)

	metaData := utility.MaterialMetadata{
		MetaDataCode:   "da03aa47171075adc32c6590585aeba1",
		MetaDataIndex:  1,
		MaterialCode:   *materialCode,
		MetaDataName:   *metaDataName,
		MetaDataValue:  *metaDataValue,
		Status:         0,
		InsertDatetime: time.Now(),
		UpdateDatetime: time.Now(),
	}
	_, err := StoreMetaData(&metaData)
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_QueryMetaDataByMetaDataCode(t *testing.T) {
	code := "da03aa47171075adc32c6590585aeba1"
	_, err := QueryMetaDataByMetaDataCode(&code)
	if err != nil {
		t.Error(err.Error())
	}
}
