package crawler

type GaijinData struct {
	Nick         string `json:"nick" mapstructure:"nick"`
	Clan         string `json:"clan" mapstructure:"clan"`
	ClanUrl      string `json:"clan_url" mapstructure:"clanUrl"`
	Banned       bool   `json:"banned" mapstructure:"banned"`
	RegisterDate string `json:"register_date" mapstructure:"register_date"`
	Title        string `json:"title" mapstructure:"title"`
	Level        string `json:"level" mapstructure:"level"`
	UserStat     struct {
		Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
		Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
		Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
	} `json:"user_stat" mapstructure:"user_stat"`
	UserRate struct {
		Aviation struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"aviation" mapstructure:"aviation"`
		GroundVehicles struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"ground_vehicles" mapstructure:"ground_vehicles"`
		Fleet struct {
			Ab map[string]string `json:"ab,omitempty" mapstructure:"ab"`
			Rb map[string]string `json:"rb,omitempty" mapstructure:"rb"`
			Sb map[string]string `json:"sb,omitempty" mapstructure:"sb"`
		} `json:"fleet" mapstructure:"fleet"`
	} `json:"user_rate" mapstructure:"user_rate"`
}
