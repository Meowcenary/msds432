package models

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
    "time"
)

// Custom type for long/lat pairs
type Point struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

// Implement custom UnmarshalJSON to handle GeoJSON "Point" type
func (p *Point) UnmarshalJSON(data []byte) error {
    // data is actually being returned in GeoJson format. This struct is necessary
    // to avoid infinite recursion by call unmarshal on "p" over and over. Kind
    // of odd to have this here, but it works, so it's staying.
    var geoJSON struct {
        Type        string    `json:"type"`
        Coordinates [2]float64 `json:"coordinates"`
    }

    // Unmarshal the GeoJSON into the temporary struct
    if err := json.Unmarshal(data, &geoJSON); err != nil {
        return fmt.Errorf("error parsing GeoJSON: %v", err)
    }

    // Ensure that the type is "Point"
    if geoJSON.Type != "Point" {
        return fmt.Errorf("unsupported GeoJSON type: %s", geoJSON.Type)
    }

    // Set Latitude and Longitude
    p.Longitude = geoJSON.Coordinates[0]
    p.Latitude = geoJSON.Coordinates[1]

    return nil
}

// Implement the Value interface to convert Point to the PostgreSQL 'POINT' type format
func (p Point) Value() (driver.Value, error) {
	// Format the Point as '(lat, lon)' for PostgreSQL
	return fmt.Sprintf("(%f,%f)", p.Latitude, p.Longitude), nil
}

// Custom type for date fields without timezone information
type DateWithoutTimezone time.Time

// Custom layout for parsing date without timezone
const dateLayout = "2006-01-02T15:04:05.000"

// UnmarshalJSON implements the custom JSON unmarshaler for DateWithoutTimezone
func (d *DateWithoutTimezone) UnmarshalJSON(data []byte) error {
    // Trim quotes from JSON string value
    str := string(data)
    str = str[1 : len(str)-1]

    // Parse date string with the custom layout
    parsedTime, err := time.Parse(dateLayout, str)
    if err != nil {
        return err
    }
    *d = DateWithoutTimezone(parsedTime)
    return nil
}

func (d DateWithoutTimezone) Value() (driver.Value, error) {
    // Cast d to time.Time and format it as a string
    return time.Time(d).Format("2006-01-02 15:04:05"), nil
}

// Struct for TaxiTrips table
type TaxiTrip struct {
    // ID                       int        `json:"id"`
    TripID                   string     `json:"trip_id"`
    TripStart                *time.Time `json:"trip_start"`
    TripEnd                  *time.Time `json:"trip_end"`
    PickupCommunityArea      *string    `json:"pickup_community_area"`
    DropoffCommunityArea     *string    `json:"dropoff_community_area"`
    PickupCentroidLatitude   *string    `json:"pickup_centroid_latitude"`
    PickupCentroidLongitude  *string    `json:"pickup_centroid_longitude"`
    PickupCentroidLocation   *Point     `json:"pickup_centroid_location"`
    DropoffCentroidLatitude  *string    `json:"dropoff_centroid_latitude"`
    DropoffCentroidLongitude *string    `json:"dropoff_centroid_longitude"`
    DropoffCentroidLocation  *Point     `json:"dropoff_centroid_location"`
    PickupZipcode            *int       `json:"pickup_zipcode"`
    DropoffZipcode           *int       `json:"dropoff_zipcode"`
}

// Struct for TransportationNetworkProvidersTrips table
type TransportationNetworkProvidersTrip struct {
    // ID                       int       `json:"id"`
    TripID                   string     `json:"trip_id"`
    TripStart                *time.Time `json:"trip_start"`
    TripEnd                  *time.Time `json:"trip_end"`
    PickupCommunityArea      *string    `json:"pickup_community_area"`
    DropoffCommunityArea     *string    `json:"dropoff_community_area"`
    PickupCentroidLatitude   *string    `json:"pickup_centroid_latitude"`
    PickupCentroidLongitude  *string    `json:"pickup_centroid_longitude"`
    PickupCentroidLocation   *Point     `json:"pickup_centroid_location"`
    DropoffCentroidLatitude  *string    `json:"dropoff_centroid_latitude"`
    DropoffCentroidLongitude *string    `json:"dropoff_centroid_longitude"`
    DropoffCentroidLocation  *Point     `json:"dropoff_centroid_location"`
    PickupZipcode            *int       `json:"pickup_zipcode"`
    DropoffZipcode           *int       `json:"dropoff_zipcode"`
}

// Struct for PublicHealthStatistics table
type PublicHealthStatistic struct {
    // ID                int     `json:"id"`
    Zipcode           *int     `json:"zipcode"`
    BelowPovertyLevel *float64 `json:"below_poverty_level"`
    PerCapitaIncome   *int     `json:"per_capita_income"`
    Unemployment      *int     `json:"unemployment"`
}

// Struct for BuildingPermits table
type BuildingPermit struct {
    // ID             int       `json:"id"`
    PermitNumber    *string              `json:"permit_"`
    StartDate       *DateWithoutTimezone `json:"application_start_date"`
    IssueDate       *DateWithoutTimezone `json:"issue_date"`
    ProcessingTime  *string              `json:"processing_time"`
    StreetName      *string              `json:"street_name"`
    StreetDirection *string              `json:"street_direction"`
    StreetNumber    *string              `json:"street_number"`
    // This will need to be added with geocoding
    ZipCode         *string              `json:"zip_code"`
}

// Struct for Covid19Reports table
type Covid19Report struct {
    // ID                          int       `json:"id"`
    ZipCode                     int       `json:"zip_code"`
    WeekNumber                  int       `json:"week_number"`
    WeekStart                   time.Time `json:"week_start"`
    WeekEnd                     time.Time `json:"week_end"`
    CasesWeekly                 int       `json:"cases_weekly"`
    CasesCumulative             int       `json:"cases_cumulative"`
    CaseRateWeekly              int       `json:"case_rate_weekly"`
    CaseRateCumulative          int       `json:"case_rate_cumulative"`
    TestsWeekly                 int       `json:"tests_weekly"`
    TestsCumulative             int       `json:"tests_cumulative"`
    TestRateWeekly              int       `json:"test_rate_weekly"`
    TestRateCumulative          int       `json:"test_rate_cumulative"`
    PercentTestedPositiveWeekly int       `json:"percent_tested_positive_weekly"`
    PercentTestedPositiveCumulative int   `json:"percent_tested_positive_cumulative"`
    ZipCodeLocation             Point     `json:"zip_code_location"`
}

// Struct for ChicagoCovid19CommunityVulnerabilityIndex table
type ChicagoCovid19CommunityVulnerabilityIndex struct {
    GeographyType                         *string `json:"geography_type"`
    CommunityAreaOrZipCode                *string `json:"community_area_or_zip"`
    CommunityAreaName                     *string `json:"community_area_name"`
    CcviScore                             *string `json:"ccvi_score"`
    CcviCategory                          *string `json:"ccvi_category"`
    RankSocioeconomicStatus               *string `json:"rank_socioeconomic_status"`
    RankHouseholdCompositionAndDisability *string `json:"rank_household_composition"`
    RankAdultsWithNoPcp                   *string `json:"rank_adults_no_pcp"`
    RankCumulativeMobilityRatio           *string `json:"rank_cumulative_mobility_ratio"`
    RankFrontlineEssentialWorkers         *string `json:"rank_frontline_essential_workers"`
    RankAge                               *string `json:"rank_age_65_plus"`
    RankComorbidConditions                *string `json:"rank_comorbid_conditions"`
    RankCovidIncidenceRate                *string `json:"rank_covid_19_incidence_rate"`
    RankCovidHospitalAdmissionRate        *string `json:"rank_covid_19_hospital_admission_rate"`
    RankCovidCrudeMortalityRate           *string `json:"rank_covid_19_crude_mortality_rate"`
    Location                              *Point  `json:"location"`
}
