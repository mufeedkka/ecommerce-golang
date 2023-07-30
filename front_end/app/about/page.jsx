"use client";
import React, { useState } from 'react';


const About = () => {
    const [count, setCount] = useState(0);

    const handleIncrement = () => {
      setCount(count + 1);
    };
  return (
    <div>
    <h2>About Mufeed Kka</h2>
    <button onClick={handleIncrement} className='bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded'>Click ME</button>
    <p>Count: {count}</p>
    </div>
  )
}

export default About