package types

func (vault Vault) StrategyDenoms() []string {
	denoms := []string{}
	denomsUsed := make(map[string]bool)
	for _, strategy := range vault.StrategyWeights {
		if denomsUsed[strategy.Denom] {
			continue
		}
		denomsUsed[strategy.Denom] = true
		denoms = append(denoms, strategy.Denom)
	}
	return denoms
}
