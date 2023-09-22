package service

func getUserInfoKey(username string) string {
	return "userInfo:" + username
}

func getAccountLocked(ti, username string) string {
	return "accountLocked:" + ti + ":" + username
}

func getLoginAttempts(username string) string {
	return "loginAttempts:" + username
}
