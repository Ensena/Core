package core

func GetInfo(userID int) {

	_, user := GetDB(userID)

	if user.MoodleUDP {
		GetMoodle(user.Email)
	}

}
