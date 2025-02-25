import React, { useState } from 'react';

function App() {
  // State to store the random number
  const [randomNumber, setRandomNumber] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // Function to fetch random number from the backend
  const fetchRandomNumber = async () => {
    setLoading(true);
    setError(null); // Reset error on new request

    try {
      const response = await fetch('http://localhost:7000/random-number');
      
      // Handle errors if the backend doesn't respond correctly
      if (!response.ok) {
        throw new Error('Failed to fetch random number');
      }

      const data = await response.json();
      setRandomNumber(data.random_number);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <h1>Random Number Generator</h1>
      <button onClick={fetchRandomNumber}>Generate Random Number</button>
      
      {loading && <p>Loading...</p>}
      
      {error && <p style={{ color: 'red' }}>Error: {error}</p>}
      
      {randomNumber !== null && (
        <div>
          <h2>Random Number: {randomNumber}</h2>
        </div>
      )}
    </div>
  );
}

export default App;
