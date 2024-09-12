package ws

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Unregister:

			delete(h.Clients, cl.Username)
		case cl := <-h.Register:

			h.Clients[cl.Username] = cl

		case m := <-h.Broadcast:
			cl, ok := h.Clients[m.Username]
			if ok {
				cl.Message <- m
			}
		}
	}
}
