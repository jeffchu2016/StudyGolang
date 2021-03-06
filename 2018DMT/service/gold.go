package service

import (
	"net/http"
	"../models"
	"../dao"
	"../control/Permission"
	"net/url"
)

func GetUserGold(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id, err := GetGetInt("Id", queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := Permission.GetUserGold(id, w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
	}
	_=uid
	g, err := dao.GetUserGold(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", g, w)
}

func SetUserGold(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	f, err := Permission.SetUserGold(id, w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	g, err := GetPostInt("Gold", w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err = dao.SetUserGold(id, g)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", "", w)
}
