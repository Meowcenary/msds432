package models

import (
    "time"
)

type Point struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

// Struct for TaxiTrips table
type TaxiTrip struct {
    ID                       int        `json:"id"`
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
    ID                       int       `json:"id"`
    TripID                   string    `json:"trip_id"`
    TripStart                *time.Time `json:"trip_start"`
    TripEnd                  *time.Time `json:"trip_end"`
    PickupCommunityArea      *int       `json:"pickup_community_area"`
    DropoffCommunityArea     *int       `json:"dropoff_community_area"`
    PickupCentroidLatitude   *float64   `json:"pickup_centroid_latitude"`
    PickupCentroidLongitude  *float64   `json:"pickup_centroid_longitude"`
    PickupCentroidLocation   *Point     `json:"pickup_centroid_location"`
    DropoffCentroidLatitude  *float64   `json:"dropoff_centroid_latitude"`
    DropoffCentroidLongitude *float64  `json:"dropoff_centroid_longitude"`
    DropoffCentroidLocation  *Point     `json:"dropoff_centroid_location"`
    PickupZipcode            *int       `json:"pickup_zipcode"`
    DropoffZipcode           *int       `json:"dropoff_zipcode"`
}

// Struct for PublicHealthStatistics table
type PublicHealthStatistic struct {
    ID                int     `json:"id"`
    Zipcode           *int     `json:"zipcode"`
    BelowPovertyLevel *float64 `json:"below_poverty_level"`
    PerCapitaIncome   *int     `json:"per_capita_income"`
    Unemployment      *int     `json:"unemployment"`
}

// Struct for BuildingPermits table
type BuildingPermit struct {
    ID             int       `json:"id"`
    PermitNumber   string    `json:"permit_number"`
    StartDate      time.Time `json:"start_date"`
    IssueDate      time.Time `json:"issue_date"`
    ProcessingTime int       `json:"processing_time"`
    StreetName     string    `json:"street_name"`
    StreetDirection string   `json:"street_direction"`
    StreetNumber   int       `json:"street_number"`
    ZipCode        int       `json:"zip_code"`
}

// Struct for Covid19Reports table
type Covid19Report struct {
    ID                          int       `json:"id"`
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
    GeographyType                       string `json:"geography_type"`
    CommunityAreaOrZipCode              int    `json:"community_area_or_zip_code"`
    CommunityAreaName                   string `json:"community_area_name"`
    CcviScore                           int    `json:"ccvi_score"`
    CcviCategory                        string `json:"ccvi_category"`
    RankSocioeconomicStatus             int    `json:"rank_socioeconomic_status"`
    RankHouseholdCompositionAndDisability int  `json:"rank_household_composition_and_disability"`
    RankAdultsWithNoPcp                 int    `json:"rank_adults_with_no_pcp"`
    RankCumulativeMobilityRatio         int    `json:"rank_cumulative_mobility_ratio"`
    RankFrontlineEssentialWorkers       int    `json:"rank_frontline_essential_workers"`
    RankAge                             int    `json:"rank_age"`
    RankComorbidConditions              int    `json:"rank_comorbid_conditions"`
    RankCovidIncidenceRate              int    `json:"rank_covid_incidence_rate"`
    RankCovidHospitalAdmissionRate      int    `json:"rank_covid_hospital_admission_rate"`
    RankCovidCrudeMortalityRate         int    `json:"rank_covid_crude_mortality_rate"`
    Location                            Point  `json:"location"`
}
