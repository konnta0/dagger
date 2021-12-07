package plancontext

import (
	"crypto/sha256"
	"fmt"
)

// Context holds the execution context for a plan.
type Context struct {
	Platform  *platformContext
	FS        *fsContext
	LocalDirs *localDirContext
	Secrets   *secretContext
	Services  *serviceContext
}

func New() *Context {
	return &Context{
		Platform: &platformContext{
			platform: defaultPlatform,
		},
		FS: &fsContext{
			store: make(map[string]*FS),
		},
		LocalDirs: &localDirContext{
			store: []string{},
		},
		Secrets: &secretContext{
			store: make(map[string]*Secret),
		},
		Services: &serviceContext{
			store: make(map[string]*Service),
		},
	}
}

func hashID(values ...string) string {
	hash := sha256.New()
	for _, v := range values {
		if _, err := hash.Write([]byte(v)); err != nil {
			panic(err)
		}
	}
	return fmt.Sprintf("%x", hash)
}
