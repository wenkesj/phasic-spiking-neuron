package sn;

type Connection struct {
  weight float64;
  target *SpikingNeuron;
  writeable bool;
};

func NewConnection(target *SpikingNeuron, weight float64, writeable bool) *Connection {
  return &Connection{
    weight: weight,
    target: target,
    writeable: writeable,
  };
};

func (this *Connection) GetTarget() *SpikingNeuron {
  return this.target;
};

func (this *Connection) GetWeight() float64 {
  return this.weight;
};

func (this *Connection) IsWriteable() bool {
  return this.writeable;
};

func (this *Connection) SetWriteabel(writeable bool) {
  this.writeable = writeable;
};

func (this *Connection) GetWriteabel() bool {
  return this.writeable;
};