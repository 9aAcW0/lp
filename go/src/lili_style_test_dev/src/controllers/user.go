package controllers

import (
	"errors"
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetUser(mail string) (models.User,error) {
	userdata, columnNum, client := utils.SearchUserDataByMail(mail)
	if userdata == nil {
		err := errors.New("メールアドレスが一致するユーザーが存在しません。")
		return models.User{}, err
	}
	answers := GetAnswers(userdata, columnNum, client)
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

func GetUserByCorporateID(corporateID, password string) ([]models.UserSearchByCorporate, error) {
	var users []models.UserSearchByCorporate
	isLogin := loginCorporate(corporateID, password)
	if isLogin {
		userdata, client := utils.SearchUserDataByCorporateID(corporateID)
		client = nil
		for _, u := range userdata {
			answers := GetAnswers(u, 0, client)
			time := u[0]
			time = time[0:10]
			user := models.UserSearchByCorporate{
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
				AltruismRate: answers.Altruism.Rate,
				BusinessStanceRate: answers.BusinessStance.Rate,
				CarrierUpRate: answers.CarrierUp.Rate,
				CommitRate: answers.Commit.Rate,
				MentalityRate: answers.Mentality.Rate,
				Answers:         answers,
				Corporate: models.Corporate {
					CorporateID: corporateID,
					Passcode: password,
				},
			}
			users = append(users, user)
		}
		return users, nil
	}
	err := errors.New("企業IDが一致しません。")
	return nil, err
}
