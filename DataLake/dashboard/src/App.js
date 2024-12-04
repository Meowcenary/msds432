// React, React Router, and Material UI
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Box } from "@mui/material";

// Custom Views...
import Dashboard from "./Dashboard";
import HomeView from "./HomeView";

// Table Views
import BuildingPermitsTableView from "./BuildingPermitsTableView";
import ChicagoCcviTableView from "./ChicagoCcviTableView";
import PublicHealthStatisticTableView from "./PublicHealthStatisticTableView";
import TaxiTripsTableView from "./TaxiTripsTableView";
import CovidReportsTableView from "./CovidReportsTableView";

// Summary
import TaxiTripsSummaryView from "./TaxiTripsSummaryView";

// Heat Maps
import TaxiTripsHeatmap from "./TaxiTripsHeatmap";

function App() {
  return (
   <Router>
      <Box sx={{ flexGrow: 1 }}>
        <Dashboard />
        <Routes>
          <Route path="/" element={<HomeView />} />
          <Route path="/taxi_trips_table" element={<TaxiTripsTableView />} />
          <Route path="/taxi_trips_summary" element={<TaxiTripsSummaryView />} />
          <Route path="/taxi_trips_heatmap" element={<TaxiTripsHeatmap />} />
          <Route path="/building_permits_table" element={<BuildingPermitsTableView />} />
          <Route path="/chicago_ccvi_table" element={<ChicagoCcviTableView />} />
          <Route path="/public_health_statistics" element={<PublicHealthStatisticTableView />} />
          <Route path="/covid_reports" element={<CovidReportsTableView />} />
        </Routes>
      </Box>
   </Router>
  );
}

export default App;
