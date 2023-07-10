package managersI

type IModule interface {
	GetGsServers() IGsServers
	GetUser() IUser
	GetCcsChallenge() IChallengeCcs
	GetShaBakeCcs() IShaBake
	GetActiveUser()  IActiveUser

}
