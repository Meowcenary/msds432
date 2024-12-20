import React, { useState, useEffect } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from '@mui/x-data-grid';

function TaxiTripsTableView() {
  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [selectedIds, setSelectedIds] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const columns: GridColDef<(typeof rows)[number]>[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'trip_id', headerName: 'Trip ID', width: 110 },
    { field: 'trip_start', headerName: 'Trip Start', width: 110 },
    { field: 'trip_end', headerName: 'Trip End', width: 110 },
    { field: 'pickup_community_area', headerName: 'Pickup Community Area', width: 110 },
    { field: 'pickup_zip', headerName: 'Pickup Zip', width: 110 },
    { field: 'dropoff_zip', headerName: 'Dropoff Zip', width: 90 },
    { field: 'dropoff_community_area', headerName: 'Dropoff Community Area', width: 110 }
  ];

  // Handle selection change
  const handleSelectionChange = (ids) => {
    console.log("handleSelectionChange called");
    setSelectedIds(ids); // Update selected IDs
  };

  // Filter rows for the selected IDs
  const filterBySelected = () => {
    setFilteredRows(rows.filter((row) => selectedIds.includes(row.id)));
  };

  // Reset to show all rows
  const resetFilter = () => {
    setFilteredRows(rows);
  };

  useEffect(() => {
    console.log("Getting taxi trips data");
    axios.get("http://localhost:8080/taxi_trips")
      .then((response) => {
        setRows(response.data);
        setFilteredRows(response.data);
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
        <h1>Loading taxi trips...</h1>
      </div>
    );
  if (error) return <div>Error: {error}</div>;
  return (
    <div style={{ height: 500, width: '100%' }}>
      <div style={{ marginBottom: '10px' }}>
        <button onClick={filterBySelected} disabled={selectedIds.length === 0}>
          Filter by Selected
        </button>
        <button onClick={resetFilter} style={{ marginLeft: '10px' }}>
          Reset
        </button>
      </div>
      <h2>Chicago Taxi Trips 2013-2023</h2>
      <DataGrid
        rows={filteredRows} // Display filtered rows
        columns={columns}
        checkboxSelection // Allow row selection
        onRowSelectionModelChange={(selectedId) => {
            console.log("onSelectionModelChange");
            handleSelectionChange(selectedId);
          }
        }
        rowSelectionModel={selectedIds}
      />
    </div>
  );
}

export default TaxiTripsTableView;
