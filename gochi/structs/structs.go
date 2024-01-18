package structs

type Client struct {
	Uid      string
	Email    string
	Age      int
	Password string
}

type HairSaloon struct {
	Uid         string
	Name        string
	Address     string
	Email       string
	Phone       string
	Openingtime string
	Closingtime string
}

type Haidresser struct {
	Uid        string
	SaloonID   *HairSaloon
	FirstName  string
	Speciality string
}

type Schedules struct {
	Uid           string
	HairdresserID *Haidresser
	StartHour     string
	EndHour       string
	Availability  bool
}

type Reservations struct {
	Uid           string
	SaloonID      *HairSaloon
	ClientID      *Client
	HairdresserID *Haidresser
	StartHour     string
	EndHour       string
	Status        string
}

type Admin struct {
	Uid      string
	Name     string
	Email    string
	Password string
}
