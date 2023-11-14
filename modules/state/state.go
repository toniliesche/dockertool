package state

import (
	"github.com/toniliesche/dockertool/modules/build"
	"github.com/toniliesche/dockertool/modules/console"
	"os"
	"os/signal"
	"syscall"
)

var AppState *State

type State struct {
	AuthorName      string
	AuthorMail      string
	Copyright       string
	CopyrightYear   string
	Sudo            bool
	SudoUser        string
	BuildCommit     string
	BuildDate       string
	BuildVersion    string
	ShutdownChannel chan os.Signal
	ReleaseChannel  string
}

func (s *State) ResetHandler() {
	signal.Reset(os.Interrupt, syscall.SIGTERM)
	signal.Notify(s.ShutdownChannel, os.Interrupt, syscall.SIGTERM)
}

func CreateState() {
	AppState = &State{
		AuthorName:     build.AuthorName,
		AuthorMail:     build.AuthorMail,
		Copyright:      build.Copyright,
		CopyrightYear:  build.CopyrightYear,
		Sudo:           console.YesNoToBool(build.Sudo),
		SudoUser:       build.SudoUser,
		BuildCommit:    build.Commit,
		BuildDate:      build.Date,
		BuildVersion:   build.Version,
		ReleaseChannel: build.Channel,
	}
}
