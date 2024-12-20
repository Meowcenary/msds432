-- This script is invoked when the postgres container is created if there is no
-- database found. This means that this this script only runs if no data has
-- been persisted i.e the postgres-data volume has been removed.
-- DROP DATABASE IF EXISTS msds432;
-- CREATE DATABASE msds432;
--
DROP TABLE IF EXISTS "TaxiTrips";
DROP TABLE IF EXISTS "TransportationNetworkProvidersTrips";
DROP TABLE IF EXISTS "BuildingPermits";
DROP TABLE IF EXISTS "ChicagoCovid19CommunityVulnerabilityIndex";
DROP TABLE IF EXISTS "PublicHealthStatistics";
DROP TABLE IF EXISTS "Covid19Reports";

\c msds432;

CREATE TABLE "TaxiTrips" (
    ID SERIAL PRIMARY KEY,
    TripID VARCHAR(100),
    TripStart TIMESTAMPTZ,
    TripEnd TIMESTAMPTZ,
    PickupCommunityArea INT,
    DropoffCommunityArea INT,
    PickupCentroidLatitude DOUBLE PRECISION,
    PickupCentroidLongitude DOUBLE PRECISION,
    PickupCentroidLocation POINT,
    DropoffCentroidLatitude DOUBLE PRECISION,
    DropoffCentroidLongitude DOUBLE PRECISION,
    DropoffCentroidLocation POINT,
    PickupZipcode INT,
    DropoffZipcode INT
);

CREATE TABLE "TransportationNetworkProvidersTrips" (
    ID SERIAL PRIMARY KEY,
    TripID VARCHAR(100),
    TripStart TIMESTAMPTZ,
    TripEnd TIMESTAMPTZ,
    PickupCommunityArea INT,
    DropoffCommunityArea INT,
    PickupCentroidLatitude DOUBLE PRECISION,
    PickupCentroidLongitude DOUBLE PRECISION,
    PickupCentroidLocation POINT,
    DropoffCentroidLatitude DOUBLE PRECISION,
    DropoffCentroidLongitude DOUBLE PRECISION,
    DropoffCentroidLocation POINT,
    PickupZipcode INT,
    DropoffZipcode INT
);

-- Varchars are temp values for now, might need to be adjusted
CREATE TABLE "BuildingPermits" (
    ID SERIAL PRIMARY KEY,
    PermitNumber VARCHAR(100),
    -- PermitType - Enum
    StartDate TIMESTAMPTZ,
    IssueDate TIMESTAMPTZ,
    ProcessingTime INT,
    StreetName VARCHAR(100),
    StreetDirection VARCHAR(10),
    StreetNumber INT,
    ZipCode INT
);

CREATE TABLE "ChicagoCovid19CommunityVulnerabilityIndex" (
    ID SERIAL PRIMARY KEY,
    GeographyType VARCHAR(10),
    CommunityAreaOrZipCode INT,
    CommunityAreaName VARCHAR(100),
    CcviScore DOUBLE PRECISION,
    CcviCategory VARCHAR(10),
    RankSocioeconomicStatus INT,
    RankHouseholdCompositionAndDisability INT,
    RankAdultsWithNoPcp INT,
    RankCumulativeMobilityRatio INT,
    RankFrontlineEssentialWorkers INT,
    RankAge INT,
    RankComorbidConditions INT,
    RankCovidIncidenceRate INT,
    RankCovidHospitalAdmissionRate INT,
    RankCovidCrudeMortalityRate INT,
    Location POINT
);

CREATE TABLE "PublicHealthStatistics" (
    ID SERIAL PRIMARY KEY,
    CommunityAreaName VARCHAR(50),
    BelowPovertyLevel DOUBLE PRECISION,
    PerCapitaIncome INT,
    Unemployment DOUBLE PRECISION
);

CREATE TABLE "Covid19Reports" (
    ID SERIAL PRIMARY KEY,
    ZipCode VARCHAR(25),
    WeekNumber INT,
    WeekStart TIMESTAMPTZ,
    WeekEnd TIMESTAMPTZ,
    CasesWeekly INT,
    CasesCumulative INT,
    CaseRateWeekly DOUBLE PRECISION,
    CaseRateCumulative DOUBLE PRECISION,
    TestsWeekly INT,
    TestsCumulative INT,
    TestRateWeekly DOUBLE PRECISION,
    TestRateCumulative DOUBLE PRECISION,
    PercentTestedPositiveWeekly DOUBLE PRECISION,
    PercentTestedPositiveCumulative DOUBLE PRECISION,
    ZipCodeLocation POINT
);
