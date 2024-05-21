package requests

type Requesting struct {
	validator      Validator
	authorizator   Permissions
	b64            B64Enc
	advRepo        Requests
	costCalculator CostCalculator
	userRepo       UserHandler
	broadCaster    BroadCaster
	files          FileHandler
}

func NewRequestHandler(val Validator, authriz Permissions,
	advRep Requests, users UserHandler, br BroadCaster, files FileHandler, b64 B64Enc) *Requesting {
	rh := new(Requesting)
	rh.validator = val
	rh.authorizator = authriz
	rh.advRepo = advRep
	rh.userRepo = users
	rh.broadCaster = br
	rh.files = files
	rh.b64 = b64
	return rh
}

func (rh *Requesting) SetCostRateCalculator(costCalc CostCalculator) {
	rh.costCalculator = costCalc
}
