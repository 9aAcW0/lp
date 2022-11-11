package controllers

import (
	"errors"
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetUser(mail string) (models.User,error) {
	userdata := utils.SearchUserDataByMail(mail)
	if userdata == nil {
		err := errors.New("メールアドレスが一致するユーザーが存在しません。")
		return models.User{}, err
	}
	answers := GetAnswers(userdata)
	time := userdata[0][0:10]
	user := models.User{
		TimeStamp: 		 time,
		Mail: 			 userdata[1],
		CorporateID: 	 userdata[2],
		Tel: 			 userdata[3],
		Name: 			 userdata[4],
		Class: 			 userdata[5],
		University: 	 userdata[6],
		UnderGraduate:   userdata[7],
		Department: 	 userdata[8],
		GraduationYear:  userdata[9],
		CurrentEmployer: userdata[10],
		Answers:         answers,
	}
	return user, nil
}

func GetUserByCorporateID(corporateID, password string) ([]models.User, error) {
	var users []models.User
	isLogin := loginCorporate(corporateID, password)
	if isLogin {
		userdata := utils.SearchUserDataByCorporateID(corporateID)
		for _, u := range userdata {
			answers := GetAnswers(u)
			time := u[0]
			time = time[0:10]
			user := models.User{
				TimeStamp:       time,
				Mail:            u[1],
				CorporateID:     u[2],
				Tel:             u[3],
				Name:            u[4],
				Class:           u[5],
				University:      u[6],
				UnderGraduate:   u[7],
				Department:      u[8],
				GraduationYear:  u[9],
				CurrentEmployer: u[10],
				Answers:         answers,
			}
			users = append(users, user)
		}
		return users, nil
	}
	err := errors.New("企業IDが一致しません。")
	return nil, err
}

func GetUserBySlash(mail string) (models.User,error) {
	userdata := utils.SearchUserDataBySlash(mail)
	if userdata == nil {
		err := errors.New("メールアドレスが一致するユーザーが存在しません。")
		return models.User{}, err
	}
	answers := GetAnswers(userdata)
	time := userdata[0][0:10]
	user := models.User{
		TimeStamp: 		 time,
		Mail: 			 userdata[1],
		CorporateID: 	 userdata[2],
		Tel: 			 userdata[3],
		Name: 			 userdata[4],
		Class: 			 userdata[5],
		University: 	 userdata[6],
		UnderGraduate:   userdata[7],
		Department: 	 userdata[8],
		GraduationYear:  userdata[9],
		CurrentEmployer: userdata[10],
		Answers:         answers,
	}
	return user, nil
}