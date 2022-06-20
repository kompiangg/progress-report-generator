package dto

type assigneeName string
type issuesTitle string

type RepositoryData map[assigneeName][]issuesTitle

type ResponseData struct {
	Title     string
	Assignees string
}

func NewRepositoryData(responseData interface{}) *RepositoryData {
	var result RepositoryData = make(RepositoryData)

	data := responseData.(map[string]interface{})["data"]
	repository := data.(map[string]interface{})["repository"]
	repositoryNode := repository.(map[string]interface{})["issues"]
	issues := repositoryNode.(map[string]interface{})["nodes"]
	issuesNodes := issues.([]interface{})

	var responseDataArr []ResponseData

	for _, issue := range issuesNodes {
		responseEachData := ResponseData{}
		assertedIssue := issue.(map[string]interface{})
		responseEachData.Title = assertedIssue["title"].(string)

		assertedAssignees := assertedIssue["assignees"].(map[string]interface{})
		assertedNameArr := assertedAssignees["nodes"].([]interface{})

		if len(assertedNameArr) != 0 {
			responseEachData.Assignees = assertedNameArr[0].(map[string]interface{})["name"].(string)
		} else {
			responseEachData.Assignees = "Other"
		}

		responseDataArr = append(responseDataArr, responseEachData)
	}

	for _, v := range responseDataArr {
		if v.Assignees == "Other" {
			result[assigneeName("Other")] = []issuesTitle{issuesTitle(v.Title)}
			continue
		} else if _, ok := result[assigneeName(v.Assignees)]; !ok {
			result[assigneeName(v.Assignees)] = []issuesTitle{issuesTitle(v.Title)}
			continue
		}

		tempAssigneeName := result[assigneeName(v.Assignees)]
		tempAssigneeName = append(tempAssigneeName, issuesTitle(v.Title))
		result[assigneeName(v.Assignees)] = tempAssigneeName
	}

	return &result
}
