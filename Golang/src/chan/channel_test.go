package chan


type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Produce(msg string)  {

}
type Consumer struct {
	ch chan string
}