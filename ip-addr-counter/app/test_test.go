package app

func Result(b []uint32) int {
	return (*Processor).Result(&Processor{data: b})
}

func ProcessIps(p *Processor, ips []byte) []uint32 {
	p.processIps(ips)
	return p.data
}
