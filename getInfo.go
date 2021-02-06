package core

func GetInfo(userID int) []byte {
	response, user := GetEnsenaData(userID)
	if user.MoodleUDP {
		GetMoodle(user.Email)
	}
	return response
}
