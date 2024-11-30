// App.js
import React from "react";
import { Box, Typography } from "@mui/material";

function App() {
  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh",
        backgroundColor: "#f5f5f5",
      }}
    >
      <Typography variant="h4" component="h1" color="primary">
        Hello, World!
      </Typography>
    </Box>
  );
}

export default App;
