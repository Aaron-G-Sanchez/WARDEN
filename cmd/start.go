package cmd

import "cycle/warden/internal/journal"

func RunStart() {
	journal.Watch()
}
