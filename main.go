package main

import (
	"fmt"
	"os"
)

type RepInfo struct {
	NumCommits    int
	firstCommitId string
	RepName       string
	RepUrl        string
	RepBranchName string
}

func concatLastCommitUrl(repInfo RepInfo) {
	commitsHistoryUrl := repInfo.RepUrl + "/commits/" + repInfo.RepBranchName + "?after=" + repInfo.firstCommitId + "+" + fmt.Sprint(repInfo.NumCommits-2)
	Ulog(commitsHistoryUrl)
}

/*
Steps As

	From "https://github.com/dotnet/runtime" get Commit Number
	From "https://github.com/dotnet/runtime/commits/main/" get First Commit Id (by html:script-payload)
	join string with "id" and "commit nums" for last Commit url like
	 https://github.com/{userName}/{repName}/commits/{branchName}?after={firstCommitId}+{commit nums - 2}
*/
func main() {
	args := os.Args
	if len(args) != 2 {
		Ulog("cmd params nums not match")
	}
	repName := args[1]
	url := "https://github.com/" + repName
	repInfo := RepInfo{
		RepName: repName,
		RepUrl:  url,
	}
	doc := ReqRepHomePage(url)
	repInfo.NumCommits = ParseCommitsNum(doc)
	repInfo.firstCommitId = ParseFirstCommitId(doc)
	repInfo.RepBranchName = ParseBranchName(doc)
	Ulog("repInfo:", repInfo)
	concatLastCommitUrl(repInfo)
}
