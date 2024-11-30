import React from "react";
import { Typography, Box } from "@mui/material";

function HomeView() {
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h4" component="h1">
        Welcome to the Chicago Open Data Lake
      </Typography>
      <Typography variant="body1">
        The dashboards available from this app visualize data ingested with Go using
        <br />
        SODA API and reports generated from that same data with Google AI tools.
      </Typography>
    </Box>
  );
}

export default HomeView;
