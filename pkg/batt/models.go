package batt

import (
	"strings"
	"time"
)

const (
	BATTERYSTATUS_NOT_CHARGING       = "NOT_CHARGING"
	BATTERYSTATUS_NORMAL_CHARGE      = "NORMAL_CHARGE"
	BATTERYSTATUS_QUICK_CHARGE       = "QUICK_CHARGE"
	BATTERYSTATUS_UNKNOWN            = "UNKNOWN"
	ISSUETYPE_POST_BOOKING           = "POST_BOOKING"
	ISSUETYPE_PRE_BOOKING            = "PRE_BOOKING"
	ISSUETYPE_FLEET                  = "FLEET"
	ISSUEREASON_BATTERY              = "BATTERY"
	ISSUEREASON_LOCATION             = "LOCATION"
	ISSUEREASON_CLOSE_NAV            = "CLOSE_NAV"
	ISSUESTATUS_CREATED              = "CREATED"
	ISSUESTATUS_RESOLVED             = "RESOLVED"
	TELEMATICSPROVIDER_FLESPI        = "FLESPI"
	TELEMATICSPROVIDER_FLESPI_TWILIO = "FLESPI_TWILIO"
	TELEMATICSPROVIDER_INVERS        = "INVERS"
)

type Period struct {
	Start *time.Time `json:"start,omitempty"`
	End   *time.Time `json:"end,omitempty"`
}

type SearchVehicleRequest struct {
	Period         *Period        `json:"period,omitempty"`
	FilterCriteria FilterCriteria `json:"filterCriteria"`
}
type FilterCriteria struct{}
type SearchBookingRequest struct {
	VehicleId       string   `json:"vehicleId"`
	Period          Period   `json:"period"`
	EndPeriod       Period   `json:"endPeriod"`
	Statuses        []string `json:"statuses"`
	NeedsCorrection *bool    `json:"needsCorrection,omitempty"`
	DoNotInvoice    *bool    `json:"doNotInvoice,omitempty"`
}
type SearchAvailabilityRequest struct {
	Period     Period   `json:"period"`
	VehicleIds []string `json:"vehicleIds"`
}

type CreateBookingRequest struct {
	Period    Period `json:"period"`
	VehicleId string `json:"vehicleId"`
	Comments  string `json:"comments"`
}

type CreateBookingResponse struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	User   struct {
		RemoteID     string      `json:"remoteId"`
		UserName     string      `json:"userName"`
		DisplayName  string      `json:"displayName"`
		ImageURL     string      `json:"imageUrl"`
		Biography    interface{} `json:"biography"`
		Role         string      `json:"role"`
		JoinedSince  interface{} `json:"joinedSince"`
		HomeLocation struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
		} `json:"homeLocation"`
		WorkLocation struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
		} `json:"workLocation"`
		FavoriteLocation struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
		} `json:"favoriteLocation"`
		PushNotificationsEnabled bool          `json:"pushNotificationsEnabled"`
		TripRegistrationEnabled  bool          `json:"tripRegistrationEnabled"`
		Organizations            []interface{} `json:"organizations"`
		Memberships              []interface{} `json:"memberships"`
		DefaultOrganization      interface{}   `json:"defaultOrganization"`
		DefaultMembership        interface{}   `json:"defaultMembership"`
		TripTypeNames            interface{}   `json:"tripTypeNames"`
	} `json:"user"`
	Vehicle struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		LicensePlate string `json:"licensePlate"`
		Address      string `json:"address"`
		HomePosition struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		} `json:"homePosition"`
		LastPosition struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			AddressLine1 string `json:"addressLine1"`
			AddressLine2 string `json:"addressLine2"`
		} `json:"lastPosition"`
		FarFromHomePosition      bool        `json:"farFromHomePosition"`
		DistanceFromHomePosition float64     `json:"distanceFromHomePosition"`
		Distance                 interface{} `json:"distance"`
		Rating                   float64     `json:"rating"`
		Owner                    string      `json:"owner"`
		Favorite                 bool        `json:"favorite"`
		TimeZone                 string      `json:"timeZone"`
		ImageURL                 string      `json:"imageUrl"`
		ElectricRange            int         `json:"electricRange"`
		Price                    float64     `json:"price"`
		PriceType                interface{} `json:"priceType"`
		KilometerPrice           float64     `json:"kilometerPrice"`
		ApprovalType             string      `json:"approvalType"`
		InstantBookingPossible   bool        `json:"instantBookingPossible"`
		OrganizationReferenceDto struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"organizationReferenceDto"`
		TripTypeNames           interface{} `json:"tripTypeNames"`
		AuthCommentsMandatory   bool        `json:"authCommentsMandatory"`
		AuthCommentsDescription interface{} `json:"authCommentsDescription"`
		ManualLink              interface{} `json:"manualLink"`
		VehicleInfo             interface{} `json:"vehicleInfo"`
		VehicleInfoPreBooking   interface{} `json:"vehicleInfoPreBooking"`
	} `json:"vehicle"`
	Period struct {
		TimeZone string `json:"timeZone"`
	} `json:"period"`
	SplittedPeriod          []interface{} `json:"splittedPeriod"`
	Comments                string        `json:"comments"`
	AuthorizationComments   interface{}   `json:"authorizationComments"`
	CancelDate              interface{}   `json:"cancelDate"`
	CancelReason            interface{}   `json:"cancelReason"`
	Business                bool          `json:"business"`
	TripType                string        `json:"tripType"`
	TripTypeNames           interface{}   `json:"tripTypeNames"`
	Status                  string        `json:"status"`
	Status2                 string        `json:"status2"`
	ApprovalReason          interface{}   `json:"approvalReason"`
	ApprovalDate            interface{}   `json:"approvalDate"`
	InBookingPeriod         bool          `json:"inBookingPeriod"`
	CanBeActivated          bool          `json:"canBeActivated"`
	Overdue                 bool          `json:"overdue"`
	ActivatedDate           interface{}   `json:"activatedDate"`
	OriginalEndDate         interface{}   `json:"originalEndDate"`
	VehicleUsageDto         interface{}   `json:"vehicleUsageDto"`
	NextPossibleStatuses    []string      `json:"nextPossibleStatuses"`
	PreviousNonAvailability struct {
		UserID   string `json:"userId"`
		UserName string `json:"userName"`
		Period   struct {
			TimeZone string `json:"timeZone"`
		} `json:"period"`
		Type   string `json:"type"`
		Type2  string `json:"type2"`
		Active bool   `json:"active"`
	} `json:"previousNonAvailability"`
	NextNonAvailability interface{} `json:"nextNonAvailability"`
	OrganizationID      string      `json:"organizationId"`
	Client              interface{} `json:"client"`
}
type Vehicle struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LicensePlate string `json:"licensePlate"`
	Address      string `json:"address"`
	HomePosition struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"homePosition"`
	LastPosition             interface{} `json:"lastPosition"`
	FarFromHomePosition      bool        `json:"farFromHomePosition"`
	DistanceFromHomePosition interface{} `json:"distanceFromHomePosition"`
	Distance                 int         `json:"distance"`
	Rating                   float64     `json:"rating"`
	Owner                    string      `json:"owner"`
	Favorite                 bool        `json:"favorite"`
	TimeZone                 string      `json:"timeZone"`
	ImageURL                 string      `json:"imageUrl"`
	ElectricRange            int         `json:"electricRange"`
	Price                    float64     `json:"price"`
	PriceType                string      `json:"priceType"`
	KilometerPrice           float64     `json:"kilometerPrice"`
	ApprovalType             string      `json:"approvalType"`
	InstantBookingPossible   bool        `json:"instantBookingPossible"`
	OrganizationReferenceDto struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"organizationReferenceDto"`
	TripTypeNames           interface{} `json:"tripTypeNames"`
	AuthCommentsMandatory   bool        `json:"authCommentsMandatory"`
	AuthCommentsDescription interface{} `json:"authCommentsDescription"`
	ManualLink              interface{} `json:"manualLink"`
	VehicleInfo             interface{} `json:"vehicleInfo"`
	VehicleInfoPreBooking   interface{} `json:"vehicleInfoPreBooking"`
}
type SearchVehicleResponse struct {
	Vehicles []Vehicle `json:"vehicles"`
}
type SearchBookingResponse struct {
	Bookings []Booking `json:"bookings"`
}

type Booking struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	UserImageURL string `json:"userImageUrl"`
	User         struct {
		RemoteID    string      `json:"remoteId"`
		UserName    interface{} `json:"userName"`
		ImageURL    string      `json:"imageUrl"`
		DisplayName string      `json:"displayName"`
	} `json:"user"`
	Vehicle Vehicle `json:"vehicle"`
	Period  struct {
		Start       string    `json:"start"`
		End         string    `json:"end"`
		ParsedStart time.Time `json:"parsedStart"`
		ParsedEnd   time.Time `json:"parsedEnd"`
		TimeZone    string    `json:"timeZone"`
	} `json:"period"`
	PlannedPeriod struct {
		Start    string `json:"start"`
		End      string `json:"end"`
		TimeZone string `json:"timeZone"`
	} `json:"plannedPeriod"`
	UsagePeriod struct {
		Start    string `json:"start"`
		End      string `json:"end"`
		TimeZone string `json:"timeZone"`
	} `json:"usagePeriod"`
	ExpressBooking          bool        `json:"expressBooking"`
	Comments                string      `json:"comments"`
	AdminComments           interface{} `json:"adminComments"`
	CancelDate              interface{} `json:"cancelDate"`
	CancelReason            interface{} `json:"cancelReason"`
	TripType                string      `json:"tripType"`
	Status                  string      `json:"status"`
	NextPossibleStatuses    []string    `json:"nextPossibleStatuses"`
	InActivePeriod          bool        `json:"inActivePeriod"`
	DoNotInvoice            bool        `json:"doNotInvoice"`
	Invoiced                *bool       `json:"invoiced"`
	ActiveBookingForVehicle interface{} `json:"activeBookingForVehicle"`
	Active                  bool        `json:"active"`
	Overdue                 bool        `json:"overdue"`
	Ended                   bool        `json:"ended"`
	ActivatedDate           string      `json:"activatedDate"`
	OriginalEndDate         interface{} `json:"originalEndDate"`
	VehicleUsageDto         struct {
		ID                string  `json:"id"`
		MileageStartValue float64 `json:"mileageStartValue"`
		MileageEndValue   float64 `json:"mileageEndValue"`
		MileageDelta      float64 `json:"mileageDelta"`
		Period            struct {
			Start    string `json:"start"`
			End      string `json:"end"`
			TimeZone string `json:"timeZone"`
		} `json:"period"`
		Status        string      `json:"status"`
		StatusMessage interface{} `json:"statusMessage"`
		StartDate     string      `json:"startDate"`
		EndDate       string      `json:"endDate"`
	} `json:"vehicleUsageDto"`
	OrganizationID string `json:"organizationId"`
	Organization   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"organization"`
	Client struct {
		ID                      string      `json:"id"`
		Name                    string      `json:"name"`
		AuthCommentsMandatory   bool        `json:"authCommentsMandatory"`
		AuthCommentsDescription interface{} `json:"authCommentsDescription"`
	} `json:"client"`
	DateCreated          string               `json:"dateCreated"`
	LastUpdated          string               `json:"lastUpdated"`
	MileageStartValue    float64              `json:"mileageStartValue"`
	MileageEndValue      float64              `json:"mileageEndValue"`
	MileageDelta         float64              `json:"mileageDelta"`
	PriceType            string               `json:"priceType"`
	Price                float64              `json:"price"`
	KmPrice              float64              `json:"kmPrice"`
	FinishedBookingPrice FinishedBookingPrice `json:"finishedBookingPrice"`
	RateCard             struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"rateCard"`
}

type BatteryStatus struct {
	BatteryPercentage int    `json:"batteryPercentage"`
	CruisingRange     int    `json:"cruisingRange"`
	Charging          string `json:"charging"`
}

type GpsLocation struct {
	GpsCoordinateDto struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"gpsCoordinateDto"`
	Address string `json:"address"`
}

type AvailabilityEventResponse struct {
	Events map[string]AvailabilityEvent `json:"availabilityEventsPerVehicleDtos"`
}

type AvailabilityEvent struct {
	Availabilities []struct {
		ID        string `json:"id"`
		VehicleID string `json:"vehicleId"`
		Vehicle   struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			LicensePlate string `json:"licensePlate"`
			ImageURL     string `json:"imageUrl"`
		} `json:"vehicle"`
		Period struct {
			Start    string `json:"start"`
			End      string `json:"end"`
			TimeZone string `json:"timeZone"`
		} `json:"period"`
		RecurringAvailabilityID interface{} `json:"recurringAvailabilityId"`
		Type                    string      `json:"type"`
	} `json:"availabilities"`
	NonAvailabilities   []NonAvailability `json:"nonAvailabilities"`
	CurrentUsageStart   interface{}       `json:"currentUsageStart"`
	CurrentUsageOverdue bool              `json:"currentUsageOverdue"`
}

type NonAvailability struct {
	VehicleID string `json:"vehicleId"`
	Type      string `json:"type"`
	Type2     string `json:"type2"`
	Period    struct {
		Start       string    `json:"start"`
		End         string    `json:"end"`
		ParsedStart time.Time `json:"-"`
		ParsedEnd   time.Time `json:"-"`
		TimeZone    string    `json:"timeZone"`
	} `json:"period"`
	Booking struct {
		ID   string `json:"id"`
		User struct {
			RemoteID    string      `json:"remoteId"`
			UserName    interface{} `json:"userName"`
			ImageURL    string      `json:"imageUrl"`
			DisplayName string      `json:"displayName"`
		} `json:"user"`
		Comments string `json:"comments"`
		Status   string `json:"status"`
		Overdue  bool   `json:"overdue"`
		Vehicle  struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			LicensePlate string `json:"licensePlate"`
			ImageURL     string `json:"imageUrl"`
		} `json:"vehicle"`
		PlannedPeriod struct {
			Start    string `json:"start"`
			End      string `json:"end"`
			TimeZone string `json:"timeZone"`
		} `json:"plannedPeriod"`
		UsagePeriod struct {
			Start    string      `json:"start"`
			End      interface{} `json:"end"`
			TimeZone string      `json:"timeZone"`
		} `json:"usagePeriod"`
		UserID          string      `json:"userId"`
		ActivatedDate   string      `json:"activatedDate"`
		OriginalEndDate string      `json:"originalEndDate"`
		UserName        interface{} `json:"userName"`
		UserImageURL    string      `json:"userImageUrl"`
		UserDisplayName string      `json:"userDisplayName"`
	} `json:"booking"`
	PlannedPeriod struct {
		Start    string `json:"start"`
		End      string `json:"end"`
		TimeZone string `json:"timeZone"`
	} `json:"plannedPeriod"`
	UsagePeriod struct {
		Start    string      `json:"start"`
		End      interface{} `json:"end"`
		TimeZone string      `json:"timeZone"`
	} `json:"usagePeriod"`
	Availability interface{} `json:"availability"`
	Overdue      bool        `json:"overdue"`
}

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type BackOfficeUserResponse struct {
	Users []BackOfficeUser `json:"results"`
}

type SearchBackOfficeUsersRequest struct {
	SofBattRemoteId string `json:"sofBattRemoteId"`
}
type BackOfficeUser struct {
	CellPhone          string `json:"cellPhone"`
	Created            string `json:"created"`
	Email              string `json:"email"`
	FirstName          string `json:"firstName"`
	ID                 int    `json:"id"`
	LastName           string `json:"lastName"`
	SofBattRemoteID    string `json:"sofBattRemoteId"`
	SofBattDisplayName string `json:"-"`
}

type CreateIssueRequest struct {
	VehicleId          string `json:"vehicleId"`
	AssignedToRemoteId string `json:"assignedToRemoteId"`
	Title              string `json:"title"`
}
type Issue struct {
	Booking           *Booking         `json:"booking"`
	Vehicle           *Vehicle         `json:"vehicle"`
	Nav               *NonAvailability `json:"nav"`
	PrevUser          *BackOfficeUser  `json:"prev_user"`
	NextUser          *BackOfficeUser  `json:"next_user"`
	BatteryStatus     *BatteryStatus   `json:"battery_status"`
	Location          *GpsLocation     `json:"location"`
	BatteryError      string           `json:"battery_error"`
	LocationError     string           `json:"location_error"`
	LastUpdated       string           `json:"lastUpdated"`
	LastUpdatedParsed time.Time        `json:"-"`
	Distance          float64          `json:"distance"`
	Reason            string           `json:"reason"`
	Status            string           `json:"status"`
	IssueType         string           `json:"issue_type"`
}

// parse last updated from issue
func (i *Issue) ParseLastUpdated() {
	parsed, err := time.Parse(time.RFC3339, strings.TrimSuffix(i.LastUpdated, "[Etc/UTC]"))
	if err != nil {
		return
	}
	i.LastUpdatedParsed = parsed
}

type IssueResponse struct {
	Issues []Issue `json:"issues"`
}

type SearchIssueRequest struct {
	VehicleId          string   `json:"vehicleId,omitempty"`
	BookingId          string   `json:"bookingId,omitempty"`
	UserRemoteId       string   `json:"userRemoteId,omitempty"`
	AssignedToRemoteId string   `json:"assignedToRemoteId,omitempty"`
	Title              string   `json:"title,omitempty"`
	Statuses           []string `json:"statuses,omitempty"`
}

type VehicleGroupsPage struct {
	VehicleGroups []VehicleGroup `json:"vehicleGroups"`
}
type VehicleGroup struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Vehicles []Vehicle `json:"vehicles"`
}

type VehicleTelematics struct {
	VehicleID        string `json:"vehicleId"`
	ProviderDeviceId string `json:"providerDeviceId"`
	ProviderId       string `json:"providerId"`
	PhoneNumber      string `json:"phoneNumber"`
	MaxRange         int    `json:"maxRange"`
}

type UpdateBookingRequest struct {
	CorrectedKm     *int  `json:"correctedKm,omitempty"`
	NeedsCorrection *bool `json:"needsCorrection,omitempty"`
}

/*
	"finishedBookingPrice": {
	  "startDate": "2025-01-19T11:00:00Z[Etc/UTC]",
	  "endDate": "2025-01-19T20:00:00Z[Etc/UTC]",
	  "kmPrice": 0.18,
	  "km": 0,
	  "amount": 1,
	  "timeAmount": 1,
	  "unit": "DAYS",
	  "unitPrice": 27.45,
	  "timeTotal": 27.45,
	  "hourlyTotalNotCharged": null,
	  "changeFromHourlyToDaily": false,
	  "kmTotal": 0,
	  "total": 27.45,
	  "kmPriceExclVat": 0.1488,
	  "unitPriceExclVat": 22.686,
	  "hourlyTotalNotChargedExclVat": null,
	  "timeTotalExclVat": 22.686,
	  "kmTotalExclVat": 0,
	  "totalExclVat": 22.686,
	  "totalVat": 4.764,
	  "vatRate": 1.21
	}
*/
type FinishedBookingPrice struct {
	StartDate                    string   `json:"startDate"`
	EndDate                      string   `json:"endDate"`
	KmPrice                      float64  `json:"kmPrice"`
	Km                           int      `json:"km"`
	Amount                       float64  `json:"amount"`
	TimeAmount                   float64  `json:"timeAmount"`
	Unit                         string   `json:"unit"`
	UnitPrice                    float64  `json:"unitPrice"`
	TimeTotal                    float64  `json:"timeTotal"`
	HourlyTotalNotCharged        *float64 `json:"hourlyTotalNotCharged,omitempty"`
	ChangeFromHourlyToDaily      bool     `json:"changeFromHourlyToDaily"`
	KmTotal                      float64  `json:"kmTotal"`
	Total                        float64  `json:"total"`
	KmPriceExclVat               float64  `json:"kmPriceExclVat"`
	UnitPriceExclVat             float64  `json:"unitPriceExclVat"`
	HourlyTotalNotChargedExclVat *float64 `json:"hourlyTotalNotChargedExclVat,omitempty"`
	TimeTotalExclVat             float64  `json:"timeTotalExclVat"`
	KmTotalExclVat               float64  `json:"kmTotalExclVat"`
	TotalExclVat                 float64  `json:"totalExclVat"`
	TotalVat                     float64  `json:"totalVat"`
	VatRate                      float64  `json:"vatRate"`
}
type BackOfficeVehicleResponse struct {
	Vehicles map[string]BackOfficeVehicle `json:"vehicles"`
}

type BackOfficeVehicle struct {
	ID                       string  `json:"id"`
	LicensePlate             string  `json:"licensePlate"`
	LeasingMonthlyPriceExVat float64 `json:"leasingMonthlyPriceExVat"`
	SofBattId                string  `json:"sofbattId"`
}
