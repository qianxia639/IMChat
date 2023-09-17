package service

func getUserInfoKey(username string) string {
	return "userInfo:" + username
}

func getLocked(ti, username string) string {
	return "locked:" + ti + ":" + username
}

func getLoginAttempts(username string) string {
	return "loginAttempts:" + username
}
