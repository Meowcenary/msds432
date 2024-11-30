import React, { useState, useEffect } from "react";
import axios from "axios";
import { MapContainer, TileLayer, useMap } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import L from "leaflet";
import "leaflet.heat";

const HeatmapLayer = ({ points }) => {
  const map = useMap();

  React.useEffect(() => {
    // Create the heatmap layer
    const heatLayer = L.heatLayer(points, {
      radius: 18,  // Radius of each "point" in the heatmap
      blur: 18,    // Blurring radius for smoothing
      maxZoom: 17, // Maximum zoom for clustering
    });

    // Add the layer to the map
    heatLayer.addTo(map);

    return () => {
      // Remove the heat layer on unmount
      map.removeLayer(heatLayer);
    };
  }, [points, map]);

  return null;
};

const TaxiTripsHeatmap = () => {
  const [heatmapData, setHeatmapData] = useState([]);

  React.useEffect(() => {
    console.log("Getting taxi trips data");
    axios.get("http://localhost:8080/taxi_trips")
      .then((response) => {
        // build heatmap data [latitude, longitude, intensity]
        var points = response.data.filter(
          (row) => row.pickup_centroid_longitude != null && row.pickup_centroid_latitude != null
        ).map((row) => [
          row.pickup_centroid_location.longitude,
          row.pickup_centroid_location.latitude,
          1
        ]);

        // console.log(points)
        setHeatmapData(points);
      })
      .catch((error) => console.error("Error fetching data:", error));
  }, []);

  return (
    <div style={{ height: "100vh" }}>
      <MapContainer
        center={[41.8781, -87.6298]} // Center map on Chicago
        zoom={12}                     // Set initial zoom level
        style={{ height: "100%", width: "100%" }}
      >
        {/* Add a TileLayer for map visuals */}
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        {/* Add the HeatmapLayer */}
        <HeatmapLayer points={heatmapData} />
      </MapContainer>
    </div>
  );
};

export default TaxiTripsHeatmap;
