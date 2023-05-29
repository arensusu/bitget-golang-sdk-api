package plan

type CancelPlanBySymbolReq struct {
	/**
	 * Currency pair
	 */
	Symbol string `json:"symbol"`
	/**
	 * Plan type
	 */
	PlanType string `json:"planType"`
}
