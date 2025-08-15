package playbook

type Command interface {
	CreateUser(user string) bool
	UserExists(user string) bool
	CreateDir(path string) bool
	DirExists(path string) bool
	CreateFile(path string) bool
	FileExist(path string) bool
	AppendToFile(path string, data []byte)
	ChangeOwner(owner, path string)
	ChangePerm(perm, path string)
	ChangePassword(user, password string)
}
