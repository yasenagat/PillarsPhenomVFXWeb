package mysqlStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoProject(project *utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO project
		(project_code, project_name, picture, project_describe, project_leader, status)
		VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.ProjectCode, project.ProjectName, project.Picture,
		project.ProjectDescribe, project.ProjectLeader, project.Status)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func DeleteProjectByProjectCode(email *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE project SET status = 1
		WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func QueryProjectList() (*[]utility.Project, error) {
	result, err := mysqlUtility.DBConn.Query(`SELECT project_code, project_name,
		picture, project_describe, project_leader, status, insert_datetime, update_datetime
		FROM project WHERE status = 0`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var projectList []utility.Project
	var i int = 0
	for result.Next() {
		var project utility.Project
		err = result.Scan(&(project.ProjectCode), &(project.DisplayName),
			&(project.Picture), &(project.ProjectDescribe), &(project.ProjectLeader),
			&(project.Status), &(project.InsertDatetime), &(project.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		project.Status = i%2 + 1
		projectList = append(projectList, project)
		i++
	}
	return &projectList, err

}

func DeleteProjectByProjectCode(projectCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE project SET status = 1
		WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(projectCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryProjectByProjectCode(projectCode *string) (*utility.Project, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT project_code, project_name,
		picture, project_describe, project_leader, status, insert_datetime, update_datetime
		FROM project WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var project utility.Project
	if result.Next() {
		err = result.Scan(&(project.ProjectCode), &(project.DisplayName),
			&(project.Picture), &(project.ProjectDescribe), &(project.ProjectLeader), &(project.Status),
			&(project.InsertDatetime), &(project.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &project, err

}

func UpdateProjectByProjectCode(project *utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE project SET project_name = ?,
		project_describe = ?, project_leader = ?, update_datetime = now() WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.DisplayName, project.ProjectDescribe, project.ProjectLeader, project.ProjectCode)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
