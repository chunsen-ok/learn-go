package main

type Params struct{}

type Adder interface {
	Add(p *Params)
}

type Deleter interface {
	Del(p *Params)
}

type Getter interface {
	Get(p *Params)
}

type AllGetter interface {
	GetAll(p *Params)
}

type Updater interface {
	Update(p *Params)
}
