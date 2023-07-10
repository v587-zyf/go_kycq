package user

// 是否是本服玩家
func (this *UserManager) IsLocalServerUser(userId int) bool {
	return this.usersBasicInfoMap[userId] != nil
}

func (this *UserManager) IsThisServerUser(userId, serverId int) bool {
	if this.usersBasicInfoMap[userId] == nil {
		return false
	}

	if this.usersBasicInfoMap[userId].ServerId == serverId {
		return true
	}
	return false
}

func (this *UserManager) CheckNicknameUnique(nickname string) bool {
	this.usersMu.Lock()
	defer this.usersMu.Unlock()
	for _, v := range this.usersBasicInfoMap {
		if v.NickName == nickname {
			return true
		}
	}
	return false
}
