package playbook

type Command interface {
	CreateUser(user string) bool
	AddUserToSudo(user string)
	AddSSHKey(user, publicKey string)
	UserExists(user string) bool
	CreateDir(path string) bool
	DirExists(path string) bool
	CreateFile(path string) bool
	FileExist(path string) bool
	AppendToFile(path, data string)
	ChangeOwner(owner, path string)
	ChangePerm(perm, path string)
	ChangePassword(user, password string)
}
