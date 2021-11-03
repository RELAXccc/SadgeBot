package server

import (
  "github.com/jaredfolkins/badactor"
  "log"
)

func (cb *CozyBot) WhenJailed(a *badactor.Actor, r *badactor.Rule) error {
  // todo: mute user


  switch r.Name {
	case "spamBlock":
	  //idk
	case "spamKick":
		// todo: change code to acomedate Guild (discord channel identity)
	default:
		log.Println("Error: undefined jailor rule: " + r.Name)
	}
	return nil
}

func (cb *CozyBot) WhenTimeServed(a *badactor.Actor, r *badactor.Rule) error {
  // Do something here. Log, email, etc...
  //todo: unmite user
	return nil
}