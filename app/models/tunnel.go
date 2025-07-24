package models

type Tunnel struct {
	ID        int    `json:"id"`
	Source      Server `json:"source"`
	Destination Server `json:"destination"`
	Type      string `json:"type"` // "direct" or "reverse"
	Protocol  string `json:"protocol"` // "tcp" or "udp"
	ListenIP  string `json:"listen_ip"`
	ListenPort int    `json:"listen_port"`
	TargetIP   string `json:"target_ip"`
	TargetPort int    `json:"target_port"`
}
