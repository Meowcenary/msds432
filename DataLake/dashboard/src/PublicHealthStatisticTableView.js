import React, { useState, useEffect } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from '@mui/x-data-grid';

function PublicHealthStatisticTableView() {
  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [selectedIds, setSelectedIds] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const columns: GridColDef<(typeof rows)[number]>[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'community_area', headerName: 'Community Area', width: 110 },
    { field: 'below_poverty_level', headerName: 'Below Poverty Level', width: 110 },
    { field: 'per_capita_income', headerName: 'Per Capita Income', width: 110 },
    { field: 'unemployment', headerName: 'Unemployment', width: 110 },
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
    console.log("Getting public health stats useEffect");
    axios.get("http://localhost:8080/public_health_stats")
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
        <h1>Loading public health statistics...</h1>
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
      <h2>Chicago Public Health Statistics</h2>
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

export default PublicHealthStatisticTableView
