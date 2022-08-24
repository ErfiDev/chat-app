package engine

func (en *Engine) Listener() {
	for {
		select {
		case m := <-en.messages:
			en.broadcast(m)

		case e := <-en.events:
			en.handleEvent(e)

		case <-en.quit:
			close(en.messages)
			close(en.events)
			close(en.quit)
			return

		default:
		}
	}
}
