package models

type SearchRequest struct {
	TMDB            string
	IMDB		string
	Year            int
	Codec           string
	Text		string
	PreferredAuthor string
	Edition         string
	SkipSearch      bool
	EntryID         string
	MVAdjust        float64
	DryrunMode      bool
	MediaType       string
	Devices         []string
	Slots           []int
	Title           string
	// I know this is a right awful place to put this, but
	// I'm not a professional programmer and this is the
	// best I could come up with that would be available later
	// to manipulate as we're doing the query
	Season		int
	Episode		int
}

type BeqCatalog struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	SortTitle  string   `json:"sortTitle"`
	Year       int      `json:"year"`
	AudioTypes []string `json:"audioTypes"`
	Digest     string   `json:"digest"`
	MvAdjust   float64  `json:"mvAdjust"`
	Edition    string   `json:"edition"`
	MovieDbID  string   `json:"theMovieDB"`
	Author     string   `json:"author"`
}

type BeqDevices struct {
	Name         string     `json:"name"`
	MasterVolume float64    `json:"mastervolume"`
	Mute         bool       `json:"mute"`
	Slots        []BeqSlots `json:"slots"`
}

type BeqSlots struct {
	ID     string  `json:"id"`
	Last   string  `json:"last"`
	Active bool    `json:"active"`
	Gain1  float64 `json:"gain1"`
	Gain2  float64 `json:"gain2"`
	Mute1  bool    `json:"mute1"`
	Mute2  bool    `json:"mute2"`
}

type BeqPatchV2 struct {
	Mute         bool      `json:"mute"`
	MasterVolume float64   `json:"masterVolume"`
	Slots        []SlotsV2 `json:"slots"`
}
type SlotsV2 struct {
	ID     string    `json:"id"`
	Active bool      `json:"active"`
	Gains  []float64 `json:"gains"`
	Mutes  []bool    `json:"mutes"`
	Entry  string    `json:"entry"`
}

type BeqPatchV1 struct {
	Mute         bool      `json:"mute"`
	MasterVolume float64   `json:"masterVolume"`
	Slots        []SlotsV1 `json:"slots"`
}
type SlotsV1 struct {
	ID     string    `json:"id"`
	Active bool      `json:"active"`
	Gains  []float64 `json:"gains"`
	Mutes  []bool    `json:"mutes"`
	Entry  string    `json:"entry"`
}
