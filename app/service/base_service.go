package service

type BaseService[Req any, Resp any, Filter any] interface {
	Create(Req) (response Resp)
	List(*Filter) (responses []Resp)
}
