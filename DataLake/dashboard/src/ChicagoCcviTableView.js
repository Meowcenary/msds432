import React, { useState, useEffect } from "react";
import axios from "axios";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Typography, Box } from "@mui/material";

function ChicagoCcviTableView() {
  const [ccviRecords, setCcviRecords] = useState([]);

  useEffect(() => {
    console.log("Getting building permits");
    axios.get("http://localhost:8080/chicago_ccvi")
      .then((response) => setCcviRecords(response.data))
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Chicago Covid Vulnerability Index
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Geography Type</TableCell>
              <TableCell>Community Area or Zip</TableCell>
              <TableCell>Commmunity Area Name</TableCell>
              <TableCell>CCVI Score</TableCell>
              <TableCell>CCVI Category</TableCell>
              <TableCell>Socioeconomic Status</TableCell>
              <TableCell>Household Composition</TableCell>
              <TableCell>Adults No PCP</TableCell>
              <TableCell>Cumulative Mobility Ratio</TableCell>
              <TableCell>Frontline Essential Workers</TableCell>
              <TableCell>Age 65+</TableCell>
              <TableCell>Comorbid Conditions</TableCell>
              <TableCell>Covid 19 Incidence Rate</TableCell>
              <TableCell>Hospital Admission Rate</TableCell>
              <TableCell>Crude Mortality Rate</TableCell>
              {/* <TableCell>Location</TableCell> */}
            </TableRow>
          </TableHead>
          <TableBody>
            {ccviRecords.map((ccviRecord) => (
              <TableRow key={ccviRecord.id}>
                <TableCell>{ccviRecord.id}</TableCell>
                <TableCell>{ccviRecord.geography_type}</TableCell>
                <TableCell>{ccviRecord.community_area_or_zip}</TableCell>
                <TableCell>{ccviRecord.community_area_name}</TableCell>
                <TableCell>{ccviRecord.ccvi_score}</TableCell>
                <TableCell>{ccviRecord.ccvi_category}</TableCell>
                <TableCell>{ccviRecord.rank_socioeconomic_status}</TableCell>
                <TableCell>{ccviRecord.rank_household_composition}</TableCell>
                <TableCell>{ccviRecord.rank_adults_no_pcp}</TableCell>
                <TableCell>{ccviRecord.rank_cumulative_mobility_ratio}</TableCell>
                <TableCell>{ccviRecord.rank_frontline_essential_workers}</TableCell>
                <TableCell>{ccviRecord.rank_age_65_plus}</TableCell>
                <TableCell>{ccviRecord.rank_comorbid_conditions}</TableCell>
                <TableCell>{ccviRecord.rank_covid_19_incidence_rate}</TableCell>
                <TableCell>{ccviRecord.rank_covid_19_hospital_admission_rate}</TableCell>
                <TableCell>{ccviRecord.rank_covid_19_crude_mortality_rate}</TableCell>
                {/* <TableCell>{ccviRecord.location}</TableCell> */}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}

export default ChicagoCcviTableView;
