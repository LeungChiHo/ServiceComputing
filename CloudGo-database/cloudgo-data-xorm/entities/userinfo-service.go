package entities



//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session  := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(u)
	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
		return err
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	alluser := make([]UserInfo, 0)
	err := engine.Find(&alluser)
	checkErr(err)
	return alluser
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var u UserInfo
	_, err := engine.Id(id).Get(&u)
	checkErr(err)
	return &u
}


