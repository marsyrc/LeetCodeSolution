package hashmap

type AuthenticationManager struct {
	timeToLive int
	memo       map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		memo:       map[string]int{},
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.memo[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	oldTime, ok := this.memo[tokenId]
	if ok {
		if currentTime-oldTime < this.timeToLive {
			this.memo[tokenId] = currentTime
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	cnt := 0
	for _, time := range this.memo {
		if time+this.timeToLive > currentTime {
			cnt++
		}
	}
	return cnt
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
