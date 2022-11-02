package day07

import "aoc2016/utils"

type IPv7Address string

const hypernetStart = '['
const hypernetEnd = ']'

type IPv7AddressAnalyze struct {
	abba string
	abas []string
	babs []string
}

func (analyze IPv7AddressAnalyze) ContainsBAB(bab string) bool {
	for _, candidate := range analyze.babs {
		if candidate == bab {
			return true
		}
	}
	return false
}

func (address IPv7Address) Analyze() IPv7AddressAnalyze {
	result := IPv7AddressAnalyze{}
	hypernetSequence := false
	abbaBlocked := false
	for i := 0; i < len(address) - 2; i++ {
		if address[i] == hypernetStart {
			hypernetSequence = true
		} else if address[i] == hypernetEnd {
			hypernetSequence = false
		} else {
			if i < len(address) - 3 {
				if address[i] == address[i + 3] && address[i + 1] == address[i + 2] && address[i] != address[i + 1] {
					if hypernetSequence {
						// the IP must not have an ABBA within any hypernet sequences
						abbaBlocked = true
						result.abba = ""
					} else if !abbaBlocked {
						abba := make([]byte, 4)
						abba[0] = address[i]
						abba[1] = address[i+1]
						abba[2] = address[i+2]
						abba[3] = address[i+3]
						result.abba = string(abba)
					}
				}
			}
			if address[i] == address[i + 2] && address[i] != address[i + 1] {
				aba := make([]byte, 3)
				aba[0] = address[i]
				aba[1] = address[i+1]
				aba[2] = address[i+2]
				if hypernetSequence {
					result.babs = append(result.babs, string(aba))
				} else {
					result.abas = append(result.abas, string(aba))
				}
			}
		}
	}
	return result
}

func (address IPv7Address) SupportsTLS() bool {
	analyze := address.Analyze()
	return analyze.abba != ""
}

func (address IPv7Address) SupportsSSL() bool {
	analyze := address.Analyze()
	for _, aba := range analyze.abas {
		bab := make([]byte, 3)
		bab[0] = aba[1]
		bab[1] = aba[0]
		bab[2] = aba[1]
		if analyze.ContainsBAB(string(bab)) {
			return true
		}
	}
	return false
}

func Solution1() int {
	count := 0
	for _, line := range utils.ReadLines("assets/input07.txt") {
		address := IPv7Address(line)
		if address.SupportsTLS() {
			count++
		}
	}
	return count
}

func Solution2() int {
	count := 0
	for _, line := range utils.ReadLines("assets/input07.txt") {
		address := IPv7Address(line)
		if address.SupportsSSL() {
			count++
		}
	}
	return count
}