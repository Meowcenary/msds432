import React, { useState, useEffect } from "react";
import axios from "axios";
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Typography, Box } from "@mui/material";

function TaxiTripsTableView() {
  const [taxiTrips, setTaxiTrips] = useState([]);

  useEffect(() => {
    console.log("Getting taxi trips data");
    axios.get("http://localhost:8080/taxi_trips")
      .then((response) => setTaxiTrips(response.data))
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        Chicago Taxi Trips 2013-2023
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Trip ID</TableCell>
              <TableCell>Trip Start</TableCell>
              <TableCell>Trip End</TableCell>
              <TableCell>Pickup Community Area</TableCell>
              <TableCell>Pickup Zip</TableCell>
              <TableCell>Dropoff Community Area</TableCell>
              <TableCell>Dropoff Zip</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {taxiTrips.map((taxiTrip) => (
              <TableRow key={taxiTrip.id}>
                <TableCell>{taxiTrip.id}</TableCell>
                <TableCell>{taxiTrip.trip_id}</TableCell>
                <TableCell>{taxiTrip.trip_start}</TableCell>
                <TableCell>{taxiTrip.trip_end}</TableCell>
                <TableCell>{taxiTrip.pickup_community_area}</TableCell>
                <TableCell>{taxiTrip.pickup_zipcode}</TableCell>
                <TableCell>{taxiTrip.dropoff_community_area}</TableCell>
                <TableCell>{taxiTrip.dropoff_zipcode}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}

export default TaxiTripsTableView;
