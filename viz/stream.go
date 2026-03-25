package stream

import (
	"emergence-sim/sim"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Serve(addr string, snapshotChan chan *sim.Snapshot) error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		defer conn.Close()

		for snapshot := range snapshotChan {
			conn.WriteJSON(snapshot) // simple JSON
		}
	})
	return http.ListenAndServe(addr, nil)
}
