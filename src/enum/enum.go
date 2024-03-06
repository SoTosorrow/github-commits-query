package enum

const (
	GithubPredix        string = "https://github.com/"
	GithubCommitsSuffix string = "/commits/"

	FullAddressNotifications string = "https://github.com/notifications"
	FullAddressExplore       string = "https://github.com/explore"
	FullAddressTrending      string = "https://github.com/trending"
	FullAddressHackerNews    string = "https://news.ycombinator.com"

	/*
		** last check at 2024-3-06 **
		length of the id of commit
	*/
	ConstOidLength int = 40
	/*
		** last check at 2024-3-06 **
		Page span of commit history
		You can click *Next* in the commit history and see it (like 'after=xxx+34') at the end of url
	*/
	ConstCommitsSingle int = 35
)
