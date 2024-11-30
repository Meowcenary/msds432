import React, { useState } from 'react';
import { Link } from "react-router-dom";
import { AppBar, Toolbar, Button, Menu, MenuItem } from "@mui/material";
// import BuildIcon from '@mui/icons-material/Build';
import HomeIcon from '@mui/icons-material/Home';
import ListAltIcon from '@mui/icons-material/ListAlt';
import MapIcon from '@mui/icons-material/Map';
import TableChartIcon from '@mui/icons-material/TableChart';

function Dashboard() {
  const [anchorEl, setAnchorEl] = useState(null);
  const [currentMenu, setCurrentMenu] = useState(null);

  const handleMenuClick = (event, menuType) => {
    setAnchorEl(event.currentTarget);
    setCurrentMenu(menuType);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
    setCurrentMenu(null);
  };

  return (
    <AppBar position="static">
      <Toolbar>
        <Button color="inherit" component={Link} to="/" startIcon={<HomeIcon />}>
          Home
        </Button>
        {/* Dropdown menu for Tables */}
        <Button color="inherit" onClick={(event) => handleMenuClick(event, 'tables')} startIcon={<TableChartIcon />}>
          Tables
        </Button>
        <Menu anchorEl={anchorEl} open={Boolean(anchorEl) && currentMenu === 'tables'} onClose={handleMenuClose} >
          <MenuItem component={Link} to="/taxi_trips_table" onClick={handleMenuClose} >
            Taxi Trips Table
          </MenuItem>
          <MenuItem component={Link} to="/building_permits_table" onClick={handleMenuClose} >
            Building Permits
          </MenuItem>
          <MenuItem component={Link} to="/chicago_ccvi_table" onClick={handleMenuClose} >
            Chicago Covid Vulnerability Index
          </MenuItem>
          <MenuItem component={Link} to="/public_health_statistics" onClick={handleMenuClose} >
            Public Health Statistics
          </MenuItem>
        </Menu>
        <Button color="inherit" onClick={(event) => handleMenuClick(event, 'summarystats')} startIcon={<ListAltIcon />} >
          Summary Statistics
        </Button>
        <Menu anchorEl={anchorEl} open={Boolean(anchorEl) && currentMenu === 'summarystats'} onClose={handleMenuClose} >
          <MenuItem component={Link} to="/taxi_trips_summary" onClick={handleMenuClose} >
            Taxi Trips Summary
          </MenuItem>
        </Menu>
        <Button color="inherit" onClick={(event) => handleMenuClick(event, 'heatmaps')} startIcon={<MapIcon />} >
            Heat Maps
        </Button>
        <Menu anchorEl={anchorEl} open={Boolean(anchorEl) && currentMenu === 'heatmaps'} onClose={handleMenuClose} >
          <MenuItem component={Link} to="/taxi_trips_heatmap" onClick={handleMenuClose} >
            Taxi Trips Heatmap
          </MenuItem>
        </Menu>
        {/* Add other dropdown menus here */}
      </Toolbar>
    </AppBar>
  )
}

export default Dashboard;
