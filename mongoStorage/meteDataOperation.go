package mongoStorage

import (
	"PillarsPhenomVFXWeb/mongoUtility"
	"PillarsPhenomVFXWeb/utility"

	"gopkg.in/mgo.v2/bson"
)

func StoreMetaData(metaData *utility.MaterialMetadata) (bool, error) {
	err := mongoUtility.MetadataCollection.Insert(metaData)
	if err != nil {
		return false, err
	}
	return true, err
}

func QueryMetaDataByMetaDataCode(metaDataCode *string) (*utility.MaterialMetadata, error) {
	var metaData utility.MaterialMetadata
	err := mongoUtility.MetadataCollection.Find(bson.M{"metadatacode": *metaDataCode}).One(&metaData)
	return &metaData, err
}
