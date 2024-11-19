package models

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
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

// Scan implements the sql.Scanner interface for Point
func (p *Point) Scan(value interface{}) error {
    // Check if the value is a []uint8 (returned by PostgreSQL driver for textual data)
    bytes, ok := value.([]uint8)
    if !ok {
        return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type Point", value)
    }

    // Convert the byte slice to a string
    str := string(bytes)

    // Parse PostgreSQL POINT format: "(longitude, latitude)"
    if !strings.HasPrefix(str, "(") || !strings.HasSuffix(str, ")") {
        return fmt.Errorf("invalid POINT format: %s", str)
    }

    // Trim parentheses
    coords := strings.TrimPrefix(str, "(")
    coords = strings.TrimSuffix(coords, ")")

    // Split the coordinates on the comma
    parts := strings.Split(coords, ",")
    if len(parts) != 2 {
        return fmt.Errorf("invalid POINT data, expected two coordinates: %s", str)
    }

    // Parse Longitude and Latitude
    lon, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
    if err != nil {
        return fmt.Errorf("invalid longitude: %v", err)
    }

    lat, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
    if err != nil {
        return fmt.Errorf("invalid latitude: %v", err)
    }

    // Set the parsed Longitude and Latitude
    p.Longitude = lon
    p.Latitude = lat

    return nil
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

// Struct for ChicagoCovid19CommunityVulnerabilityIndex table
// A lot of these fields are blank, but that's just how the data is
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

// Struct for PublicHealthStatistics table
type PublicHealthStatistic struct {
    // ID                int     `json:"id"`
    CommunityAreaName *string `json:"community_area"`
    BelowPovertyLevel *string `json:"below_poverty_level"`
    PerCapitaIncome   *string `json:"per_capita_income"`
    Unemployment      *string `json:"unemployment"`
}

// Struct for Covid19Reports table
type Covid19Report struct {
    // ID                              int       `json:"id"`
    ZipCode                         *string   `json:"zip_code"`
    WeekNumber                      *string   `json:"week_number"`
    WeekStart                       *DateWithoutTimezone `json:"week_start"`
    WeekEnd                         *DateWithoutTimezone `json:"week_end"`
    CasesWeekly                     *string   `json:"cases_weekly"`
    CasesCumulative                 *string   `json:"cases_cumulative"`
    CaseRateWeekly                  *string   `json:"case_rate_weekly"`
    CaseRateCumulative              *string   `json:"case_rate_cumulative"`
    TestsWeekly                     *string   `json:"tests_weekly"`
    TestsCumulative                 *string   `json:"tests_cumulative"`
    TestRateWeekly                  *string   `json:"test_rate_weekly"`
    TestRateCumulative              *string   `json:"test_rate_cumulative"`
    PercentTestedPositiveWeekly     *string   `json:"percent_tested_positive_weekly"`
    PercentTestedPositiveCumulative *string   `json:"percent_tested_positive_cumulative"`
    ZipCodeLocation                 *Point    `json:"zip_code_location"`
}
