package staking_service

import "fmt"

func (stakingService *StakingService) DailyRewards(uid string, startTime int64, endTime int64) (result *RewardList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	if uid != "" {
		req.SetQueryParam("uid", uid)
	}
	req.SetQueryParam("startTime", fmt.Sprintf("%d", startTime))
	req.SetQueryParam("endTime", fmt.Sprintf("%d", endTime))
	res, err := req.Send("GET", "/openapi/rewards/daily")
	if err != nil {
		return
	}

	return processResponse[RewardList](res)
}

func (stakingService *StakingService) MonthlyRewards(uid string, startTime int64, endTime int64) (result *RewardList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	if uid != "" {
		req.SetQueryParam("uid", uid)
	}
	req.SetQueryParam("startTime", fmt.Sprintf("%d", startTime))
	req.SetQueryParam("endTime", fmt.Sprintf("%d", endTime))
	res, err := req.Send("GET", "/openapi/rewards/monthly")
	if err != nil {
		return
	}

	return processResponse[RewardList](res)
}

func (stakingService *StakingService) QuarterlyRewards(uid string, startTime int64, endTime int64) (result *RewardList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	if uid != "" {
		req.SetQueryParam("uid", uid)
	}
	req.SetQueryParam("startTime", fmt.Sprintf("%d", startTime))
	req.SetQueryParam("endTime", fmt.Sprintf("%d", endTime))
	res, err := req.Send("GET", "/openapi/rewards/quaterly")
	if err != nil {
		return
	}

	return processResponse[RewardList](res)
}

func (stakingService *StakingService) AnnuallyRewards(uid string, startTime int64, endTime int64) (result *RewardList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	if uid != "" {
		req.SetQueryParam("uid", uid)
	}
	req.SetQueryParam("startTime", fmt.Sprintf("%d", startTime))
	req.SetQueryParam("endTime", fmt.Sprintf("%d", endTime))
	res, err := req.Send("GET", "/openapi/rewards/annually")
	if err != nil {
		return
	}

	return processResponse[RewardList](res)
}
