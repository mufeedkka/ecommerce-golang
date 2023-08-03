"use client";


import React, { useState } from "react";

const ProductForm = () => {
  const [longName, setlongName] = useState("");
  const [shortName, setshortName] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();

    // Create an object with the form data
    const formData = {
        shortName: shortName,
        longName: longName,
    };

    // Perform any client-side logic here before submitting the form data to the API
    console.log("Form Data:", formData);

    // Send the form data to the API using fetch
    fetch("http://localhost:8080/createproduct", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    })
      .then((response) => response.json())
      .then((data) => {
        // Handle the API response (if needed)
        console.log("API Response:", data);

        // Example: Reset form fields after successful submission
        setlongName("");
        setshortName("");
      })
      .catch((error) => {
        // Handle any error that occurred during the API request
        console.error("API Error:", error);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>
          LongName:
          <input
            type="text"
            value={longName}
            onChange={(e) => setlongName(e.target.value)}
          />
        </label>
      </div>
      <div>
        <label>
        shortName:
          <input
            type="text"
            value={shortName}
            onChange={(e) => setshortName(e.target.value)}
          />
        </label>
      </div>
      <div>
        <button type="submit">Create</button>
      </div>
    </form>
  );
};

export default ProductForm;
