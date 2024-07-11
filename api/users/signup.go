package users

import "my-docs/reqmodel"

/*
version 1.0
不考虑验证
*/
func (u *UserController) SignUp() {
	var req reqmodel.SignUpReq
	err := u.ShouldBindJSON(&req)
	if err != nil {
		u.CJSON(400, err)
	}

}
