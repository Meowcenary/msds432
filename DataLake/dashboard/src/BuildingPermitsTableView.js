import React, { useState, useEffect } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from '@mui/x-data-grid';

function BuildingPermitsTableView() {
  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [selectedIds, setSelectedIds] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const columns: GridColDef<(typeof rows)[number]>[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'permit_', headerName: 'Permit Number', width: 110 },
    { field: 'processing_time', headerName: 'Processing Time', width: 110 },
    { field: 'street_name', headerName: 'Street Name', width: 110 },
    { field: 'street_direction', headerName: 'Street Direction', width: 110 },
    { field: 'street_number', headerName: 'Street Number', width: 110 },
    { field: 'zip_code', headerName: 'ZIP', width: 90 }
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

  // Load data for view
  useEffect(() => {
    console.log("Getting building permits useEffect");
    axios.get("http://localhost:8080/building_permits")
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
        <h1>Loading building permits...</h1>
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
      <h2>Chicago Building Permits</h2>
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

export default BuildingPermitsTableView;
