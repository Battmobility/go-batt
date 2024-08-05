package batt

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/umahmood/haversine"
)

func TestSearchNonAvailability(t *testing.T) {

}

func TestBattClient_SearchBookings(t *testing.T) {
	startP := time.Now()
	endP := startP.Add(3 * time.Hour)
	type fields struct {
		Base  string
		Token string
	}
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *SearchVehicleResponse
		wantErr bool
	}{
		{
			name: "TestGetBookings",
			fields: fields{
				Base:  "https://webapp.battmobility.com/api/web-bff-service",
				Token: os.Getenv("token"),
			},
			args: args{
				start: startP,
				end:   endP,
			},
			wantRes: &SearchVehicleResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BattClient{
				SofBattHost: tt.fields.Base,
				Token:       &tt.fields.Token,
			}
			gotRes, err := c.SearchBookings(&tt.args.start, &tt.args.end)
			if err != nil {
				t.Error(err)
			}
			vids := []string{}
			for _, b := range gotRes.Bookings {
				if b.Period.ParsedStart.After(startP) {
					vids = append(vids, b.Vehicle.ID)
				}
			}
			s := startP.Add(-1 * time.Hour)
			e := startP
			n, err := c.SearchAvailabilities(&s, &e, vids)
			if err != nil {
				t.Error(err)
			}
			for _, b := range gotRes.Bookings {
				if b.Period.ParsedStart.After(startP) {
					var activeNav *NonAvailability
					for _, nav := range n.Events[b.Vehicle.ID].NonAvailabilities {
						if nav.Period.ParsedEnd.After(time.Now()) && nav.Period.ParsedStart.Before(time.Now()) {
							activeNav = &nav
						}
					}
					if activeNav != nil {
						fmt.Println("Active non availability")
						fmt.Printf("end: %s\n", activeNav.Period.End)
						fmt.Printf("type: %s\n", activeNav.Type2)
						if activeNav.Type2 == "BOOKING" {
							fmt.Println(activeNav.Booking.UserDisplayName)
						}
					} else {
						fmt.Println("no active non availability")
					}
					gps, err := c.RefreshLocation(b.Vehicle.ID)
					if err != nil {
						fmt.Println(err)
						fmt.Println("============")
						continue
					}
					loc := haversine.Coord{
						Lat: gps.GpsCoordinateDto.Latitude,
						Lon: gps.GpsCoordinateDto.Longitude,
					}
					home := haversine.Coord{
						Lat: b.Vehicle.HomePosition.Latitude,
						Lon: b.Vehicle.HomePosition.Longitude,
					}
					_, km := haversine.Distance(loc, home)
					battery, err := c.GetBatteryStatus(b.Vehicle.ID)
					if err != nil {
						fmt.Println(err)
						fmt.Println("============")
						continue
					}
					fmt.Println("Next booking")
					fmt.Println(b.Vehicle.Name)
					fmt.Println(b.Period.Start)
					fmt.Println(b.UserName)
					fmt.Println(b.User.RemoteID)
					fmt.Println("charging " + battery.Charging)
					fmt.Printf("battery percentage: %d \n", battery.BatteryPercentage)
					fmt.Printf("current address %s\n", gps.Address)
					fmt.Printf("home address %s\n", b.Vehicle.Address)
					fmt.Printf("distance from home %.2f km\n", km)
					fmt.Println("============")
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("BattClient.SearchBookings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
