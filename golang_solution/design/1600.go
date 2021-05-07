package design

type ThroneInheritance struct {
	dict map[string][]string
	dead map[string]struct{}
	king string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		dict: make(map[string][]string),
		dead: make(map[string]struct{}),
		king: kingName,
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	if _, ok := t.dict[parentName]; !ok {
		t.dict[parentName] = []string{}
	}
	t.dict[parentName] = append(t.dict[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = struct{}{}
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	res := []string{}
	var dfs func(name string)
	dfs = func(name string) {
		if _, over := t.dead[name]; !over {
			res = append(res, name)
		}
		for _, child := range t.dict[name] {
			dfs(child)
		}
	}
	dfs(t.king)
	return res
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
