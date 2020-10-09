// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package im

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"github.com/thinkcmf/catgo/im/models"
	//"fmt"
)

//type Subscription struct {
//	Archive []models.Event      // All the events from the archive.
//	New     <-chan models.Event // New events coming in.
//}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg}
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 1)
	// Channel for exit users.
	unsubscribe = make(chan string, 1)
	// Send events here to publish them.
	publish = make(chan models.Event, 1)
	// Long polling waiting list.
	waitingList = list.New()
	//subscribers = list.New()
	subscribers = make(map[string]Subscriber)
)

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(sub.Name) {
				subscribers[sub.Name] = sub
				// Publish a JOIN event.
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:

			ws := subscribers[unsub].Conn
			delete(subscribers, unsub)
			if ws != nil {
				ws.Close()
				beego.Error("WebSocket closed:", unsub)
			}
			publish <- newEvent(models.EVENT_LEAVE, unsub, "") // Publish a LEAVE event.

		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(user string) bool {
	if _, ok := subscribers[user]; ok {
		return true
	}

	return false
}
