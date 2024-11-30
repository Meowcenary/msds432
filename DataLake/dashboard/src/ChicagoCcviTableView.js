import React, { useState, useEffect } from "react";
import axios from "axios";
import { DataGrid, GridColDef } from '@mui/x-data-grid';


function ChicagoCcviTableView() {
  const [rows, setRows] = useState([]);
  const [filteredRows, setFilteredRows] = useState([]);
  const [selectedIds, setSelectedIds] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const columns: GridColDef<(typeof rows)[number]>[] = [
    { field: 'id', headerName: 'ID', width: 90 },
    { field: 'geography_type', headerName: 'Geo Type', width: 90 },
    { field: 'community_area_or_zip', headerName: 'Community Area or Zip', width: 100 },
    { field: 'community_area_name', headerName: 'Community Area Name', width: 110 },
    { field: 'ccvi_score', headerName: 'CCVI Score', width: 90 },
    { field: 'ccvi_category', headerName: 'CCVI Score', width: 90 },
    { field: 'rank_socioeconomic_status', headerName: 'Socioeconomic Status', width: 100 },
    { field: 'rank_household_composition', headerName: 'Household Composition', width: 100 },
    { field: 'rank_adults_no_pcp', headerName: 'Adults No PCP', width: 100 },
    { field: 'rank_cumulative_mobility_ratio', headerName: 'Cumulative Mobility Ratio', width: 100 },
    { field: 'rank_frontline_essential_workers', headerName: 'Frontline Essential Workers', width: 100 },
    { field: 'rank_age_65_plus', headerName: '65+', width: 90 },
    { field: 'rank_comorbid_conditions', headerName: 'Comborbid Conditions', width: 100 },
    { field: 'rank_covid_19_incidence_rate', headerName: 'Covid 19 Incidence Rate', width: 100 },
    { field: 'rank_covid_19_hospital_admission_rate', headerName: 'Hospital Admission Rate', width: 100 },
    { field: 'rank_covid_19_crude_mortality_rate', headerName: 'Crude Mortality Rate', width: 100 },
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
    axios.get("http://localhost:8080/chicago_ccvi")
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

export default ChicagoCcviTableView;
