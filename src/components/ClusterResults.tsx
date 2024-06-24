// components/ClusterResults.tsx
import React, { useState } from 'react';

const ClusterResults = () => {
  const [result, setResult] = useState<any>(null);

  const handleCluster = async () => {
    try {
      const response = await fetch('/api/cluster', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
      });
      const data = await response.json();
      setResult(data);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  return (
    <div>
      <h1>K-Means Clustering Results</h1>
      <button onClick={handleCluster}>Run Clustering</button>
      {result && (
        <div>
          <h2>Centroids</h2>
          <ul>
            {result.centroids.map((centroid: any, index: number) => (
              <li key={index}>
                Latitude: {centroid.latitud}, Longitude: {centroid.longitud}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default ClusterResults;
