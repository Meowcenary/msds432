import React, { useState, useEffect } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from '@mui/x-data-grid';


function CovidReportsTableView() {
  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [selectedIds, setSelectedIds] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const columns: GridColDef<(typeof rows)[number]>[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'zip_code', headerName: 'Zip Code', width: 90 },
    { field: 'week_number', headerName: 'Week Number', width: 90 },
    { field: "cases_weekly", headerName: 'Cases Weekly', width: 90 },
    { field: "cases_cumulative", headerName: 'Cases Cumulative', width: 90 },
    { field: "case_rate_weekly", headerName: 'Case Rate Weekly', width: 90 },
    { field: "case_rate_cumulative", headerName: 'Case Rate Cumulative', width: 100 },
    { field: "tests_weekly", headerName: 'Tests Weekly', width: 90 },
    { field: "tests_cumulative", headerName: 'Tests Cumulative', width: 90 },
    { field: "test_rate_weekly", headerName: 'Test Rate Weekly', width: 90 },
    { field: "test_rate_cumulative", headerName: 'Test Rate Cumulative', width: 100 },
    { field: "percent_tested_positive_weekly", headerName: 'Percent Tested Positive Weekly', width: 90 },
    { field: "percent_tested_positive_cumulative", headerName: 'Percent Tested Positive Cumulative', width: 90 },
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
    axios.get("http://localhost:8080/covid_19_reports")
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
        <h1>Loading Chicago CCVI...</h1>
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
      <h2>Chicago Covid Vulnerability Index</h2>
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

export default CovidReportsTableView;
