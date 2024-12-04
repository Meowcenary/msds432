import React, { useState, useEffect } from "react";
import { Box, Typography, Card, CardContent, Grid } from '@mui/material';

function TaxiTripsSummaryView() {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);


  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Taxi Trip Table Summary Statistics
      </Typography>
      <Grid container spacing={3}>
          {[{label: "Total Trips", value: "11,400"}, {label: "Average Trip Time", value: "20 minutes"}, {label: "Most Frequently Visited Zip Code", value: "60606"}].map((stat, index) => (
          <Grid item xs={12} sm={6} md={3} key={index}>
            <Card sx={{ textAlign: 'center', bgcolor: '#f5f5f5' }}>
              <CardContent>
                <Typography variant="h6" color="text.secondary">
                  {stat.label}
                </Typography>
                <Typography variant="h4" color="primary">
                  {stat.value}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default TaxiTripsSummaryView;
