package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-fsnotify/fsevents"
)

var (
	hostPath     []string
	guestPath    []string
	ignoreExt    []string
	eventHistory []string
)

func main() {
	// Handle args

	es := &fsevents.EventStream{
		// This should be the global paths var
		Paths:   []string{"./"},
		Latency: 100 * time.Millisecond,
		// Ignore WatchRoot
		Flags: fsevents.FileEvents}

	es.Start()
	ec := es.Events

	go func() {
		for msg := range ec {
			for _, event := range msg {
				handleEvent(event)
			}
		}
	}()

	in := bufio.NewReader(os.Stdin)

	if false {
		log.Print("Started, press enter to GC")
		in.ReadString('\n')
		runtime.GC()
		log.Print("GC'd, press enter to quit")
		in.ReadString('\n')
	} else {
		log.Print("Started, press enter to stop")
		in.ReadString('\n')
		es.Stop()

		log.Print("Stopped, press enter to restart")
		in.ReadString('\n')
		es.Resume = true
		es.Start()

		log.Print("Restarted, press enter to quit")
		in.ReadString('\n')
		es.Stop()
	}

	// in := bufio.NewReader(os.Stdin)

	// if false {
	// 	log.Print("Started, press enter to GC")
	// 	in.ReadString('\n')
	// 	runtime.GC()
	// 	log.Print("GC'd, press enter to quit")
	// 	in.ReadString('\n')
	// } else {
	// 	log.Print("Started, press enter to stop")
	// 	in.ReadString('\n')
	// 	es.Stop()

	// 	log.Print("Stopped, press enter to restart")
	// 	in.ReadString('\n')
	// 	es.Resume = true
	// 	es.Start()

	// 	log.Print("Restarted, press enter to quit")
	// 	in.ReadString('\n')
	// 	es.Stop()
	// }
}

// Map of events we car about
var noteDescription = map[fsevents.EventFlags]string{
	fsevents.ItemCreated:       "Created",
	fsevents.ItemRemoved:       "Removed",
	fsevents.ItemInodeMetaMod:  "InodeMetaMod",
	fsevents.ItemRenamed:       "Renamed",
	fsevents.ItemModified:      "Modified",
	fsevents.ItemFinderInfoMod: "FinderInfoMod",
	fsevents.ItemChangeOwner:   "ChangeOwner",
	fsevents.ItemXattrMod:      "XattrMod",
	fsevents.ItemIsFile:        "IsFile",
	fsevents.ItemIsDir:         "IsDir",
	fsevents.ItemIsSymlink:     "IsSymlink",
}

func handleEvent(event fsevents.Event) {
	var eventFlags = []string{}

	for bit, description := range noteDescription {
		if event.Flags&bit == bit {
			eventFlags = append(eventFlags, description)
		}
	}

	if len(eventFlags) == 0 {
		return
	}

	//

	log.Printf("EventID: %d Path: %s Flags: %s", event.ID, event.Path, eventFlags)
}
