package record

type Answer struct {
	Name  uint16
	Type  uint16
	Class uint16
	TTL   uint32
	Len   uint16
	Data  []byte
}

// -- RECORD TYPES -- //

type SOA struct {
	Name    string `yaml:"name"`
	Admin   string `yaml:"admin"`
	Serial  uint32 `yaml:"serial"`
	Refresh uint32 `yaml:"refresh"`
	Retry   uint32 `yaml:"retry"`
	Expire  uint32 `yaml:"expire"`
	Minimum uint32 `yaml:"minimum"`
	TTL     uint32 `yaml:"ttl"`
}

type A struct {
	Records []*struct {
		TTL  uint32 `yaml:"ttl"`
		IPv4 string `yaml:"ipv4"`
	} `yaml:"records"`
}

type AAAA struct {
	Records []*struct {
		TTL  uint32 `yaml:"ttl"`
		IPv6 string `yaml:"ipv6"`
	} `yaml:"records"`
}

type TXT struct {
	Records []*struct {
		TTL   uint32 `yaml:"ttl"`
		Value string `yaml:"value"`
	} `yaml:"records"`
}

type CNAME struct {
	TTL    uint32 `yaml:"ttl"`
	Target string `yaml:"target"`
}

type MX struct {
	Records []*struct {
		TTL      uint32 `yaml:"ttl"`
		Priority uint16 `yaml:"priority"`
		Server   string `yaml:"server"`
	} `yaml:"records"`
}

type NS struct {
	Records []*struct {
		TTL    uint32 `yaml:"ttl"`
		Server string `yaml:"server"`
	} `yaml:"records"`
}

type PTR struct {
	Records []*struct {
		TTL    uint32 `yaml:"ttl"`
		Domain string `yaml:"domain"`
	} `yaml:"records"`
}

type SRV struct {
	Records []*struct {
		TTL      uint32 `yaml:"ttl"`
		Priority uint16 `yaml:"priority"`
		Weight   uint16 `yaml:"weight"`
		Port     uint16 `yaml:"port"`
		Target   string `yaml:"target"`
	} `yaml:"records"`
}

type CAA struct {
	Records []*struct {
		TTL   uint32 `yaml:"ttl"`
		Flag  uint8  `yaml:"flag"`
		Tag   string `yaml:"tag"`
		Value string `yaml:"value"`
	} `yaml:"records"`
}

type CERT struct {
	Records []*struct {
		TTL    uint32 `yaml:"ttl"`
		Type   uint16 `yaml:"type"`
		KeyTag uint16 `yaml:"keytag"`
		Algo   uint8  `yaml:"algo"`
		Cert   string `yaml:"cert"`
	} `yaml:"records"`
}

type DNSKEY struct {
	Records []*struct {
		TTL       uint32 `yaml:"ttl"`
		Flags     uint16 `yaml:"flags"`
		Proto     uint8  `yaml:"proto"`
		Algo      uint8  `yaml:"algo"`
		PublicKey string `yaml:"publickey"`
	} `yaml:"records"`
}

type DS struct {
	Records []*struct {
		TTL        uint32 `yaml:"ttl"`
		KeyTag     uint16 `yaml:"keytag"`
		Algo       uint8  `yaml:"algo"`
		DigestType uint8  `yaml:"digesttype"`
		Digest     string `yaml:"digest"`
	} `yaml:"records"`
}

type HTTPS struct {
	Records []*struct {
		TTL           uint32            `yaml:"ttl"`
		Priority      uint16            `yaml:"priority"`
		Target        string            `yaml:"target"`
		ALPN          []string          `yaml:"alpn"`
		IPv4Hint      []string          `yaml:"ipv4hint"`
		IPv6Hint      []string          `yaml:"ipv6hint"`
		NoDefaultALPN bool              `yaml:"no-default-alpn"`
		Mandatory     []string          `yaml:"mandatory"`
		Port          uint16            `yaml:"port"`
		DOHPath       string            `yaml:"dohpath"`
		Other         map[string]string `yaml:"other"`
	} `yaml:"records"`
}

type Coords struct {
	Deg int     `yaml:"deg"`
	Min int     `yaml:"min"`
	Sec float64 `yaml:"sec"`
	Hem string  `yaml:"hem"`
}

type LOC struct {
	Records []*struct {
		TTL uint32  `yaml:"ttl"`
		Lat *Coords `yaml:"lat"`
		Lon *Coords `yaml:"lon"`

		Prec *struct {
			Alt  float64 `yaml:"alt"`
			Size float64 `yaml:"size"`
			Horz float64 `yaml:"horz"`
			Vert float64 `yaml:"vert"`
		} `yaml:"prec"`
	} `yaml:"records"`
}

type NAPTR struct {
	Records []*struct {
		TTL     uint32 `yaml:"ttl"`
		Order   uint16 `yaml:"order"`
		Pref    uint16 `yaml:"pref"`
		Flags   string `yaml:"flags"`
		Service string `yaml:"service"`
		RegEx   string `yaml:"regex"`
		Replace string `yaml:"replace"`
	} `yaml:"records"`
}

type SMIMEA struct {
	Records []*struct {
		TTL      uint32 `yaml:"ttl"`
		Usage    uint8  `yaml:"usage"`
		Selector uint8  `yaml:"selector"`
		Match    uint8  `yaml:"matchtype"`
		Cert     string `yaml:"cert"`
	} `yaml:"records"`
}

type SSHFP struct {
	Records []*struct {
		TTL         uint32 `yaml:"ttl"`
		Algo        uint8  `yaml:"algo"`
		Type        uint8  `yaml:"type"`
		Fingerprint string `yaml:"fingerprint"`
	} `yaml:"records"`
}

type SVCB struct {
	Records []*struct {
		TTL           uint32            `yaml:"ttl"`
		Priority      uint16            `yaml:"priority"`
		Target        string            `yaml:"target"`
		ALPN          []string          `yaml:"alpn"`
		IPv4Hint      []string          `yaml:"ipv4hint"`
		IPv6Hint      []string          `yaml:"ipv6hint"`
		NoDefaultALPN bool              `yaml:"no-default-alpn"`
		Mandatory     []string          `yaml:"mandatory"`
		Port          uint16            `yaml:"port"`
		DOHPath       string            `yaml:"dohpath"`
		Other         map[string]string `yaml:"other"`
	} `yaml:"records"`
}

type TLSA struct {
	Records []*struct {
		TTL      uint32 `yaml:"ttl"`
		Usage    uint8  `yaml:"usage"`
		Selector uint8  `yaml:"selector"`
		Match    uint8  `yaml:"matchtype"`
		Cert     string `yaml:"cert"`
	} `yaml:"records"`
}

type URI struct {
	Records []*struct {
		TTL      uint32 `yaml:"ttl"`
		Priority uint16 `yaml:"priority"`
		Weight   uint16 `yaml:"weight"`
		Target   string `yaml:"target"`
	} `yaml:"records"`
}
