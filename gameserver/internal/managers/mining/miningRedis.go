package mining

//func (this *MiningManager) UpdateMiningWork(user *objs.User) {
//	if len(this.MiningSlice) <= 0 {
//		return
//	}
//	lastUserId := this.MiningSlice[len(this.MiningSlice)-2]
//	if user.Id == lastUserId {
//		lastExpire := this.MiningSlice[len(this.MiningSlice)-1]
//		if int64(lastExpire) <= time.Now().Unix() {
//			this.DelData(lastUserId)
//		}
//	}
//}

//func (this *MiningManager) updateService() {
//	for {
//		select {
//		case msg := <-this.updateChan:
//			rmodel.Mining.ZaddMiningWork(msg.Member, base.Conf.ServerId, msg.Score)
//			this.MiningSlice = append(this.MiningSlice, msg.Member, msg.Score)
//			logger.Debug("updateService MiningSlice:%v", this.MiningSlice)
//		}
//	}
//}
//
//func (this *MiningManager) Append(member, score int) {
//	select {
//	case this.updateChan <- MiningChange{Member: member, Score: score}:
//	default:
//		logger.Warn("MiningManager: Append is full, please check .")
//	}
//}
//
//func (this *MiningManager) DelData(userId int) {
//	rmodel.Mining.ZremMiningWork(userId, base.Conf.ServerId)
//	this.MiningSlice = this.MiningSlice[:len(this.MiningSlice)-2]
//	logger.Debug("DelData MiningSlice:%v", this.MiningSlice)
//}
//
//func (this *MiningManager) GetScore(userId int) int {
//	return rmodel.Mining.ZscoreMiningWork(userId, base.Conf.ServerId)
//}
