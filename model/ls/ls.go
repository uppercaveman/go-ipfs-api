package ls

// LsLink :  ls
type LsLink struct {
	Name, Hash string
	Size       uint64
	Type       int
}

// LsObject :
type LsObject struct {
	Hash  string
	Links []LsLink
}

// LsOutput :
type LsOutput struct {
	Objects []LsObject
}
