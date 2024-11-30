// React, React Router, and Material UI
import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { AppBar, Toolbar, Button, Box } from "@mui/material";

// Custom Views
import BuildingPermitsTableView from "./BuildingPermitsTableView";
import ChicagoCcviTableView from "./ChicagoCcviTableView";
import HomeView from "./HomeView";
import TaxiTripsTableView from "./TaxiTripsTableView";

function App() {
  return (
    <Router>
      <Box sx={{ flexGrow: 1 }}>
        {/* Navigation Bar */}
        <AppBar position="static">
          <Toolbar>
            <Button color="inherit" component={Link} to="/">
              Home
            </Button>
            <Button color="inherit" component={Link} to="/taxi_trips_table">
              Taxi Trips Table
            </Button>
            <Button color="inherit" component={Link} to="/building_permits_table">
              Building Permits
            </Button>
            <Button color="inherit" component={Link} to="/chicago_ccvi_table">
              Chicago Covid Vulnerability Index
            </Button>
          </Toolbar>
        </AppBar>

        {/* Routes */}
        <Routes>
          <Route path="/" element={<HomeView />} />
          <Route path="/taxi_trips_table" element={<TaxiTripsTableView />} />
          <Route path="/building_permits_table" element={<BuildingPermitsTableView />} />
          <Route path="/chicago_ccvi_table" element={<ChicagoCcviTableView />} />
        </Routes>
      </Box>
    </Router>
  );
}

export default App;
