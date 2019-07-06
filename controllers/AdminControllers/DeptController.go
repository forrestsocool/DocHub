package AdminControllers

import "github.com/TruthHun/DocHub/models"

type DeptController struct {
	BaseController
}

//用户权限管理
func (this *DeptController) Get() {
	//var log models.CoinLog
	//log.Uid, _ = this.GetInt("uid")
	//log.Coin, _ = this.GetInt("score")
	//log.Log = this.GetString("log")
	Uid, _ := this.GetInt("uid")
	UserCid, _ := this.GetInt("cid")
	_, err := models.UpdateByIds(models.GetTableUserInfo(), "Cid", UserCid,  Uid)
	//if err == nil {
	//	err = models.NewCoinLog().LogRecord(log)
	//}
	if err != nil {
		this.ResponseJson(false, err.Error())
	}
	this.ResponseJson(true, "用户权限变更成功")
}
