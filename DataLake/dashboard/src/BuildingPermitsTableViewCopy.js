import React, { useState, useEffect } from "react";
import axios from "axios";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Typography, Box } from "@mui/material";

function BuildingPermitsTableView() {
  const [buildingPermits, setBuildingPermits] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Load data for view
  useEffect(() => {
    console.log("Getting building permits useEffect");
    axios.get("http://localhost:8080/building_permits")
      .then((response) => {
        setBuildingPermits(response.data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  // Rendering
  if (loading)
    return (
      <div>
        <h1>Loading building permits...</h1>
      </div>
    );
  if (error) return <div>Error: {error}</div>;
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Chicago Building Permits
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Permit Number </TableCell>
              {/* <TableCell>Start Date</TableCell> */}
              {/* <TableCell>Issue Date</TableCell> */}
              <TableCell>Processing Time</TableCell>
              <TableCell>Street Name</TableCell>
              <TableCell>Street Direction</TableCell>
              <TableCell>Street Number</TableCell>
              <TableCell>Zip Code</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {buildingPermits.map((buildingPermit) => (
              <TableRow key={buildingPermit.id}>
                <TableCell>{buildingPermit.id}</TableCell>
                <TableCell>{buildingPermit.permit_}</TableCell>
                {/* <TableCell>{buildingPermit.application_start_date}</TableCell> */}
                {/* <TableCell>{buildingPermit.issue_date}</TableCell> */}
                <TableCell>{buildingPermit.processing_time}</TableCell>
                <TableCell>{buildingPermit.street_name}</TableCell>
                <TableCell>{buildingPermit.street_direction}</TableCell>
                <TableCell>{buildingPermit.street_number}</TableCell>
                <TableCell>{buildingPermit.zip_code}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}

export default BuildingPermitsTableView;
