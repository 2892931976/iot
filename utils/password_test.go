package utils

import "testing"

func TestHashPassword(t *testing.T) {
	password := "123456"
	hashedPassword, err := HashPassword(password)
	if err == nil {
		t.Log(hashedPassword)
		t.Log(len(hashedPassword))
	} else {
		t.Log(err)
	}

	ok := CheckPasswordHash(password, hashedPassword)

	if ok {
		t.Log("密码验证正确")
	} else {
		t.Log("验证函数错误")
	}
}

func TestCheckM5Password(t *testing.T) {
	ok := CheckM5Password("123456", "e10adc3949ba59abbe56e057f20f883e")
	if ok {
		t.Log("md5 密码验证正确")
	} else {
		t.Error("md5 验证失败")
	}
}
