package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterial(m *utility.Material) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material (library_code, material_code, material_name, material_type, material_path, video_track_count, width, height, video_audio_framerate, timecode_framerate, video_frame_count, start_absolute_timecode, end_absolute_timecode, start_edge_timecode, end_edge_timecode, meta_data, picture, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.LibraryCode, m.MaterialCode, m.MaterialName, m.MaterialType, m.MaterialPath, m.VideoTrackCount, m.Width, m.Height, m.VideoAudioFramerate, m.TimecodeFramerate, m.VideoFrameCount, m.StartAbsoluteTimecode, m.EndAbsoluteTimecode, m.StartEdgeTimecode, m.EndEdgeTimecode, m.MetaData, m.Picture, m.UserCode, m.ProjectCode, m.Status)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteMaterialByMaterialCode(materialCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material SET status = 1 WHERE material_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(materialCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}

	return true, err
}

func QueryMaterials(code string, start int64, end int64) (*[]utility.Material, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT * FROM (SELECT material_code, material_name, material_type, material_path FROM material WHERE status = 0 AND library_code = ? ORDER BY update_datetime DESC) T LIMIT ?, ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code, start, end)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var materials []utility.Material
	for result.Next() {
		var m utility.Material
		err = result.Scan(&(m.MaterialCode), &(m.MaterialName), &(m.MaterialType), &(m.MaterialPath))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		materials = append(materials, m)
	}
	return &materials, err
}

func QueryMaterialByMaterialCode(materialCode *string) (*utility.Material, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT material_code, material_path, material_name, material_type FROM material WHERE status = 0 AND material_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(materialCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var m utility.Material
	if result.Next() {
		err = result.Scan(&(m.MaterialCode), &(m.MaterialPath), &(m.MaterialName), &(m.MaterialType), &(m.Status), &(m.InsertDatetime), &(m.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}

	return &m, err
}
