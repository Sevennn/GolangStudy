package entities


//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(u)
	if err != nil {
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (*UserInfoAtomicService) FindAll() []UserInfo {
	allr := make([]UserInfo, 0)
	engine.Find(&allr)
	return allr
}

func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var u UserInfo
	engine.Id(id).Get(&u)
	return &u
}